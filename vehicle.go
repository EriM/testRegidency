package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type VEHICLEChaincode struct {
}

func (this *VEHICLEChaincode) newWorker(stub shim.ChaincodeStubInterface) *VEHICLE {
	logger := shim.NewLogger("VEHCLE")
	return &VEHICLE{stub: stub, logger: logger}
}

func (this *VEHICLEChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	worker := this.newWorker(stub)
	return worker.Init(function, args)
}

func (this *VEHICLEChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	worker := this.newWorker(stub)
	return worker.Invoke(function, args)
}

func (this *VEHICLEChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	worker := this.newWorker(stub)
	return worker.Query(function, args)
}

func main() {
	err := shim.Start(new(VEHICLEChaincode))
	if err != nil {
		fmt.Printf("Error starting VEHICLE Chaincode: %s", err)
	}
}
