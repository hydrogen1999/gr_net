/*
 *Author:ThinhNguyenCong
 *2020-11-26
 */
package service

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (t *ServiceSetup) SaveCom(com Commodity) (string, error) {

	eventID := "eventAddCom"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// Serialize edu object into byte array
	b, err := json.Marshal(com)
	if err != nil {
		return "", fmt.Errorf("An error occurred while serializing the specified com object")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "addCom", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) FindComInfoByEntityID(entityID string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryComInfoByEntityID", Args: [][]byte{[]byte(entityID)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) FindComByCertNoAndName(certNo, name string) ([]byte, error) {

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "queryComByCertNoAndName", Args: [][]byte{[]byte(certNo), []byte(name)}}
	respone, err := t.Client.Query(req)
	if err != nil {
		return []byte{0x00}, err
	}

	return respone.Payload, nil
}

func (t *ServiceSetup) ModifyCom(com Commodity) (string, error) {

	eventID := "eventModifyCom"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// Serialize the edu object into a byte array
	b, err := json.Marshal(com)
	if err != nil {
		return "", fmt.Errorf("An error occurred while serializing the specified edu object")
	}

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "updateCom", Args: [][]byte{b, []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}

func (t *ServiceSetup) DelCom(entityID string) (string, error) {

	eventID := "eventDelEdu"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	req := channel.Request{ChaincodeID: t.ChaincodeID, Fcn: "delCom", Args: [][]byte{[]byte(entityID), []byte(eventID)}}
	respone, err := t.Client.Execute(req)
	if err != nil {
		return "", err
	}

	number, err := eventResult(notifier, eventID)
	if err != nil {
		return "", err
	}
	t.BlockNumber = number
	return string(respone.TransactionID), nil
}
