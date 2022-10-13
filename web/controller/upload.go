package controller

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request) {

	start := "{"
	content := ""
	end := "}"

	file, _, err := r.FormFile("file")
	if err != nil {
		content = "\"error\":1,\"result\":{\"msg\":\"Invalid file specified\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		content = "\"error\":1,\"result\":{\"msg\":\"Unable to read file content\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	filetype := http.DetectContentType(fileBytes)
	//log.Println("filetype = " + filetype)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		content = "\"error\":1,\"result\":{\"msg\":\"File type error\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	fileName := randToken(12)                           // Specify file name
	fileEndings, err := mime.ExtensionsByType(filetype) // Get file extension
	//log.Println("fileEndings = " + fileEndings[0])
	// Specify file storage path
	newPath := filepath.Join("web", "static", "images", fileName+fileEndings[0])
	//fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

	newFile, err := os.Create(newPath)
	if err != nil {
		log.Println("Failed to create file：" + err.Error())
		content = "\"error\":1,\"result\":{\"msg\":\"Failed to create file\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		log.Println("Failed to write file：" + err.Error())
		content = "\"error\":1,\"result\":{\"msg\":\"Failed to save file content\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	path := "/static/images/" + fileName + fileEndings[0]
	content = "\"error\":0,\"result\":{\"fileType\":\"image/png\",\"path\":\"" + path + "\",\"fileName\":\"ce73ac68d0d93de80d925b5a.png\"}"
	w.Write([]byte(start + content + end))
	return
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
