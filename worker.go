package main

import (
	"encoding/json"
//	"errors"
//	"fmt"
//	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type VEHICLE struct {
	stub     shim.ChaincodeStubInterface
	logger   *shim.ChaincodeLogger
}

func (this *VEHICLE) Init(function string, args []string) ([]byte, error) {
	this.createSchema()
	return nil, nil
}

// Run callback representing the invocation of a chaincode
func (this *VEHICLE) Invoke(function string, args []string) ([]byte, error) {
	ret, err := this.RunImpl(function, args)

//	// Error processing
//	argsWithComma := strings.Join(args, ",")
//	status := "OK"
//	message := ""
//	payload := fmt.Sprintf("%s(%s)", function, argsWithComma)

	if err != nil {
//		elements := strings.Split(fmt.Sprintf("%v", err), "|")
//		status = elements[0]
//		if len(elements) >= 2 {
//			message = elements[1]
//		}
//		if ABORT_AT_ERROR {
		return ret, err
//		}
	}
//
//	this.putInvokingStatus(&InvokingStatus{this.db.UUID, status, message, payload})

	return ret, nil
}

// Query callback representing the query of a chaincode
func (this *VEHICLE) Query(function string, args []string) ([]byte, error) {
	result, err := this.QueryImpl(function, args)
	if result != nil {
		j, _ := json.Marshal(result)
		return j, nil
	} else {
		return nil, err
	}
}
