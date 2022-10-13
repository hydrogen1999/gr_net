package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/hydrogen1999/grooo-network/sdkInit"
	"github.com/hydrogen1999/grooo-network/service"
	"github.com/hydrogen1999/grooo-network/web"
	"github.com/hydrogen1999/grooo-network/web/controller"
)

const (
	configFile  = "config.yaml"
	initialized = false
	ComCC       = "comcc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID:     "generalchannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/hydrogen1999/grooo-network/channel-artifacts/generalchannel.tx",

		OrgAdmin:       "Admin",
		OrgName:        "Grooo1",
		OrdererOrgName: "orderer.grooo.com",

		ChaincodeID:     ComCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/hydrogen1999/grooo-network/chaincode/",
		UserName:        "User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID: ComCC,
		Client:      channelClient,
	}

	coms := []service.Commodity{
		{Type: "Gạo", Primarykey: "001", Name: "Gạo lứt", Des: "Trồng trên ruộng", Specification: "1kg", Source: "Thái Bình", Machining: "Thu hoạch bằng máy gặt", Remarks: "Không", Principal: "Bảo Minh", PhoneNumber: "123456789", Photo: "/static/photo/11.png", ShelfLife: "1 năm", StorageMethod: "Nơi khô ráo, nhiệt độ bình thường", Brand: "Bảo Minh", Vendor: "Nhà máy ở Thái Bình", PlaceOfProduction: "Việt Nam", ExecutiveStandard: "TCVN 11888:2017", Time: time.Now().Format("2020-11-20 15:04:05")},
		{Type: "Gạo", Primarykey: "002", Name: "Gạo tám Thái", Des: "Trồng trên ruộng", Specification: "1kg", Source: "Nam Định", Machining: "Thu hoạch bằng máy gặt", Remarks: "Không", Principal: "Quốc Đạt", PhoneNumber: "123456789", Photo: "/static/images/abc.jpg", ShelfLife: "1 năm", StorageMethod: "Nơi khô ráo, nhiệt độ bình thường", Brand: "Quốc Đạt", Vendor: "Nhà máy ở Nam Định", PlaceOfProduction: "Việt Nam", ExecutiveStandard: "TCVN 11888:2017", Time: time.Now().Format("2020-11-20 15:04:05")},
		{Type: "Gạo", Primarykey: "003", Name: "Gạo nếp cái hoa vàng", Des: "Trồng trên ruộng", Specification: "1kg", Source: "Thái Bình", Machining: "Thu hoạch bằng máy gặt", Remarks: "Không", Principal: "Việt Hương", PhoneNumber: "123456789", Photo: "/static/photo/11.png", ShelfLife: "1 năm", StorageMethod: "Nơi khô ráo, nhiệt độ bình thường", Brand: "Việt Hương", Vendor: "Nhà máy ở Thái Bình", PlaceOfProduction: "Việt Nam", ExecutiveStandard: "TCVN 11888:2017", Time: time.Now().Format("2020-11-20 15:04:05")},
		{Type: "Gạo", Primarykey: "004", Name: "Gạo thơm nàng Sen", Des: "Trồng trên ruộng", Specification: "1kg", Source: "Đồng Nai", Machining: "Thu hoạch bằng máy gặt", Remarks: "Không", Principal: "Trà Ngọc", PhoneNumber: "123456789", Photo: "/static/photo/11.png", ShelfLife: "1 năm", StorageMethod: "Nơi khô ráo, nhiệt độ bình thường", Brand: "Trà Ngọc", Vendor: "Nhà máy ở Đồng Nai", PlaceOfProduction: "Việt Nam", ExecutiveStandard: "TCVN 11888:2017", Time: time.Now().Format("2020-11-20 15:04:05")},
	}

	for _, v := range coms {
		msg, err := serviceSetup.SaveCom(v)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("The information is successfully released, the transaction number is: " + msg)
		}
	}

	result, err := serviceSetup.FindComInfoByEntityID("001")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var com service.Commodity
		json.Unmarshal(result, &com)
		fmt.Println(com)
	}

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
