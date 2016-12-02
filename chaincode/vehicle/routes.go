package main

import (
	"errors"
//	"strings"

//	"github.com/hyperledger/fabric/core/crypto/primitives"
)


//type Response struct {
//	Contents interface{}
//}

func (this *VEHICLE) RunImpl(function string, args []string) ([]byte, error) {
	// Authentication
	// Replace harmfull chars
	// Load automatic generated routes

	// Handle different functions
	switch function {
	case "writeProtoVehicle":
		return nil, this.add_ProtoVehicle(args)

	case "writeEcuTestHistory":
		return nil, this.add_EcuTestHistory(args)

//	case "update_xx":
//	case "clear_data":
//	case "clearTableEcuTestHistory":
//	case "import":
//		return nil, this.imprtJson(args[0])
	default:
		return nil, errors.New("UNKNOWN_INVOCATION|Received unknown function invocation")
	}
}

//func (this *HDLS) auth(args []string) error {


// Query callback representing the query of a chaincode
func (this *VEHICLE) QueryImpl(function string, args []string) (interface{}, error) {
	// Replace harmfull chars

	switch function {
	case "debug":
		this.logger.Infof("Call: query - debug")
		this.logger.Debugf("%s",args[0])

//	case "getInvokingStatus":
//	case "queryStatus":
//	case "jsontest":
//	case "ping":
// Cert
//	case "metadata":
//		return this.db.GetCallerMetadata()
//	case "binding":
//	case "cert":
//	case "commonName":

//	case "role":
//		role, _ := this.db.ReadCertAttribute(args[0])
//		return string(role), nil
//	case "errortest":
//		return nil, errors.New("Error test")

	}

	// Authentication
	// Load automatic generated routes

	switch function {
	case "readProtoVehicle":
		return this.get_ProtoVehicle_list()

	case "readEcuTest":
		return this.get_EcuTestHistory_list(args)

	case "readEcuLatestTestResult":
		return this.get_EcuLatestTest_list(args)

//	case "get_invoke_status":
//	case "dump":
//		return this.dump()

	default:
		return nil, errors.New("Received unknown function query")
	}
}
