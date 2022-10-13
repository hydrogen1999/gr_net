package web

import (
	"fmt"
	"net/http"

	"github.com/hydrogen1999/grooo-network/web/controller"
)

// Start the web service and specify routing information
func WebStart(app controller.Application) {

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Specify routing information (match request)
	http.HandleFunc("/admin", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addEduInfo", app.AddEduShow)
	http.HandleFunc("/addEdu", app.AddEdu)

	http.HandleFunc("/queryPage", app.QueryPage)
	//http.HandleFunc("/query", app.FindCertByNoAndName)

	http.HandleFunc("/queryPage2", app.QueryPage2)
	http.HandleFunc("/query2", app.FindByID)

	http.HandleFunc("/modifyPage", app.ModifyShow)
	http.HandleFunc("/modify", app.Modify)

	http.HandleFunc("/upload", app.UploadFile)

	fmt.Println("Start the web service, the listening port number is: 8001")
	err := http.ListenAndServe("0.0.0.0:8001", nil)
	if err != nil {
		fmt.Printf("Web service failed to start: %v", err)
	}

}
