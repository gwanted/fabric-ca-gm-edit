/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/hyperledger/fabric/common/tools/protolator"

	// Import these to register the proto types
	_ "github.com/hyperledger/fabric/protos/common"
	_ "github.com/hyperledger/fabric/protos/msp"
	_ "github.com/hyperledger/fabric/protos/orderer"
	_ "github.com/hyperledger/fabric/protos/peer"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

func getMsgType(r *http.Request) (proto.Message, error) {
	vars := mux.Vars(r)
	msgName := vars["msgName"] // Will not arrive is unset

	msgType := proto.MessageType(msgName)
	if msgType == nil {
		return nil, fmt.Errorf("message name not found")
	}
	return reflect.New(msgType.Elem()).Interface().(proto.Message), nil
}

func Decode(w http.ResponseWriter, r *http.Request) {
	msg, err := getMsgType(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	err = proto.Unmarshal(buf, msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	var buffer bytes.Buffer
	err = protolator.DeepMarshalJSON(&buffer, msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	buffer.WriteTo(w)
}

func Encode(w http.ResponseWriter, r *http.Request) {
	msg, err := getMsgType(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	err = protolator.DeepUnmarshalJSON(r.Body, msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}
