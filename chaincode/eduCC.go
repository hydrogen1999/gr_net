/*
*Author:ThinhNguyenCong
*2020-11-26
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric/core/peer"
)

const DOC_TYPE = "comObj"
const DOC_COM_TYPE = "comObj"

// Put commodity
func PutCom(stub shim.ChaincodeStubInterface, com Commodity) ([]byte, bool) {

	com.ObjectType = DOC_COM_TYPE

	b, err := json.Marshal(com)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(com.Primarykey, b)
	if err != nil {
		return nil, false
	}

	return b, true
}

// Get commodity information
func GetComInfo(stub shim.ChaincodeStubInterface, primarykey string) (Commodity, bool) {
	var com Commodity

	b, err := stub.GetState(primarykey)
	if err != nil {
		return com, false
	}

	if b == nil {
		return com, false
	}
	err = json.Unmarshal(b, &com)
	if err != nil {
		return com, false
	}

	return com, true
}

func getComByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

func (t *EducationChaincode) addCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("The number of given parameters does not meet the requirements!")
	}

	var com Commodity
	err := json.Unmarshal([]byte(args[0]), &com)
	if err != nil {
		return shim.Error("An error occurred while deserializing information!")
	}

	// Duplicate check: ID number must be unique
	_, exist := GetComInfo(stub, com.Primarykey)
	if exist {
		return shim.Error("The traceability ID to be added already exists")
	}

	_, bl := PutCom(stub, com)
	if !bl {
		return shim.Error("An error occurred while saving information")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Information added successfully!"))
}

// Set aside
func (t *EducationChaincode) queryComByCertNoAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("The number of given parameters does not meet the requirements")
	}
	CertNo := args[0]
	name := args[1]

	// Query string required to assemble CouchDB (a standard JSON string)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"eduObj\", \"CertNo\":\"%s\"}}", CertNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\", \"CertNo\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, CertNo, name)

	// Query data
	result, err := getComByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("An error occurred when querying information based on the certificate number and name")
	}
	if result == nil {
		return shim.Error("No relevant information can be found based on the specified certificate number and name")
	}
	return shim.Success(result)
}

func (t *EducationChaincode) queryComInfoByEntityID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("The number of given parameters does not meet the requirements")
	}

	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("Failed to query information based on traceability ID")
	}

	if b == nil {
		return shim.Error("No relevant information was found based on the traceability ID.")
	}

	// Deserialize the queried state
	var com Commodity
	err = json.Unmarshal(b, &com)
	if err != nil {
		return shim.Error("Failed to deserialize traceability information")
	}

	// Get historical change data
	iterator, err := stub.GetHistoryForKey(com.Primarykey)
	if err != nil {
		return shim.Error("Failed to query the corresponding historical change data according to the specified ID number")
	}
	defer iterator.Close()

	// Iterative processing
	var historys []HistoryItem
	var hisCom Commodity
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("Failed to obtain historical change data of edu")
		}

		var historyItem HistoryItem
		historyItem.TxId = hisData.TxId
		json.Unmarshal(hisData.Value, &hisCom)

		if hisData.Value == nil {
			var empty Commodity
			historyItem.Commodity = empty
		} else {
			historyItem.Commodity = hisCom
		}

		historys = append(historys, historyItem)

	}

	com.Historys = historys

	result, err := json.Marshal(com)
	if err != nil {
		return shim.Error("An error occurred while serializing the traceability information")
	}
	return shim.Success(result)
}

func (t *EducationChaincode) updateCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("The number of given parameters does not meet the requirements")
	}

	var info Commodity
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("Failed to deserialize com information")
	}

	result, bl := GetComInfo(stub, info.Primarykey)
	if !bl {
		return shim.Error("An error occurred when querying information based on the traceability number")
	}

	result.Type = info.Type
	result.Name = info.Name
	result.Des = info.Des
	result.Specification = info.Specification
	result.Source = info.Source
	result.Machining = info.Machining
	result.Photo = info.Photo
	result.Remarks = info.Remarks
	result.Principal = info.Principal
	result.PhoneNumber = info.PhoneNumber

	result.ShelfLife = info.ShelfLife
	result.StorageMethod = info.StorageMethod
	result.Brand = info.Brand
	result.Vendor = info.Vendor
	result.PlaceOfProduction = info.PlaceOfProduction
	result.ExecutiveStandard = info.ExecutiveStandard

	result.Time = info.Time

	_, bl = PutCom(stub, result)
	if !bl {
		return shim.Error("An error occurred while saving the information")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Information updated successfully"))
}

func (t *EducationChaincode) delCom(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("The number of given parameters does not meet the requirements")
	}

	/*var edu Education
	result, bl := GetEduInfo(stub, info.EntityID)
	err := json.Unmarshal(result, &edu)
	if err != nil {
		return shim.Error("An error occurred while deserializing information")
	}*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("An error occurred while deleting the information")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("Information deleted successfully"))
}
