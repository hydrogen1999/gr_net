package service

import (
	"fmt"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type Commodity struct {
	ObjectType    string `json:"docType"`
	Type          string `json:"type"`       //Event type
	Primarykey    string `json:"primarykey"` //Primary key
	Name          string `json:"name"`
	Des           string `json:"des"`           //Description
	Specification string `json:"specification"` //Specification
	Source        string `json:"source"`        //Source
	Machining     string `json:"machining"`     //Processing
	Remarks       string `json:"remarks"`       //Remarks information
	Principal     string `json:"principal"`     //Principal
	PhoneNumber   string `json:"phoneNumber"`
	Photo         string `json:"Photo"` // Photo

	ShelfLife         string `json:"shelfLife"`         //Shelf life
	StorageMethod     string `json:"storageMethod"`     //Storage method
	Brand             string `json:"brand"`             //Brand
	Vendor            string `json:"vendor"`            //Vendor
	PlaceOfProduction string `json:"placeOfProduction"` //Production place
	ExecutiveStandard string `json:"executiveStandard"` //Executive standard

	Historys []HistoryItem // Current history of com
	Time     string        `json:"Time"` //Time
}

type HistoryItem struct {
	TxId      string
	Commodity Commodity
}

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
	BlockNumber uint64
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("Failed to register chain code event: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) (uint64, error) {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("Chaincode event received: %v\n", ccEvent)

		return ccEvent.BlockNumber, nil
	case <-time.After(time.Second * 20):
		return 0, fmt.Errorf("Cannot receive the corresponding chaincode event according to the specified event ID(%s)", eventID)
	}
	return 0, nil
}
