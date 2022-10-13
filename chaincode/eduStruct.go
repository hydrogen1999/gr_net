package main

/**

Commodity unique ID (traceability number):

 Event type:

 Introduction:

 Product name:

 Product specifications:

 Commodity source:

 Processing methods:

 Photo:

 Remarks information:

 Principal:

 Contact details:

 Entry Time:
*/
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
