package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type EducationChaincode struct {
}

func (t *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}

func (t *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	// Get user intent
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addCom" {
		return t.addCom(stub, args)
	} else if fun == "queryComByCertNoAndName" {
		return t.queryComByCertNoAndName(stub, args)
	} else if fun == "queryComInfoByEntityID" {
		return t.queryComInfoByEntityID(stub, args)
	} else if fun == "updateCom" {
		return t.updateCom(stub, args)
	} else if fun == "delCom" {
		return t.delCom(stub, args)
	}

	return shim.Error("Wrong function name specified")
y
}

func main() {
	err := shim.Start(new(EducationChaincode))
	if err != nil {
		fmt.Printf("An error occurred while starting EducationChaincode: %s", err)
	}
}
