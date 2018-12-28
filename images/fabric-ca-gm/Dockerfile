# Dockerfile for tjfoc fabric-ca image.
# If you need a peer node to run, please see the yeasy/tjfoc-peer image.
# Workdir is set to $GOPATH/src/github.com/tjfoc/fabric-ca-gm
# More usage infomation, please see https ://github .com/yeasy/dockerhyper-ledger-fabric-ca. 

FROM golang:1.9
LABEL maintainer ”zhangqichao <zhangqichao@gjk365.com>”

# ca-server and ca-client will check the following env in order, to get the home cfg path
ENV FABRIC_CA_HOME /etc/tjfoc/fabric-ca-server-gm
ENV FABRIC_CA_SERVER_HOME /etc/tjfoc/fabric-ca-server-gm 
ENV FABRIC_CA_CLIENT_HOME $HOME/.fabric-ca-client-gm
ENV CA_CFG_PATH /etc/tjfoc/fabric-ca-gm

# This is go simplify this Dockerfile

ENV FABRIC_CA_ROOT $GOPATH/src/github.com/tjfoc/fabric-ca-gm
ENV FABRIC_CA_OLD_ROOT $GOPATH/src/github.com/hyperledger/fabric-ca

# Usually the binary will be installed into $GOPATH/bin, but we add local build path, too

ENV PATH=$FABRIC_CA_ROOT/bin:$PATH

# fabric-ca-server will open service to '0.0.0.0:7054/api/vl /'
EXPOSE 7054

RUN mkdir -p $GOPATH/src/github.com/tjfoc \ 
$GOPATH/src/github.com/hyperledger \
$FABRIC_CA_SERVER_HOME \
$FABRIC_CA_CLIENT_HOME \
$CA_CFG_PATH \ 
/var/tjfoc/fabric-ca-server-gm

# Need libtool to provide the header file ltdl.h 
RUN apt-get update \
    && apt-get install -y libsnappy-dev zlib1g-dev libbz2-dev libltdl-dev libtool glibc-doc \ 
    && rm -rf /var/cache/apt

# clone and build ca
RUN cd $GOPATH/src/github.com/hyperledger \
    && git clone --single-branch -b release-1.1 --depth 1 https://github.com/hyperledger/fabric-ca

RUN cd $GOPATH/src/github.com/tjfoc \
    && git clone --single-branch -b master --depth 1 https://github.com/tjfoc/fabric-ca-gm \
    # This will install fabric-ca-server and fabric-ca-client into $GOPATH/bin/ 
    && cd fabric-ca-gm \
    && make  
    # Copy example ca and key files
    # && cp $FABRIC_CA_OLD_ROOT/images/fabric-ca/payload/*.pem $FABRIC_CA_HOME/ 
    
VOLUME $FABRIC_CA_SERVER_HOME
 
VOLUME $FABRIC_CA_CLIENT_HOME

WORKDIR $FABRIC_CA_ROOT
# if no config exists under $FABRIC_CA_HOME, will init fabric-ca-serverconfig.yaml and fabric-ca-server.db

CMD fabric-ca-server start -b admin:adminpw