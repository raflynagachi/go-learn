package goweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.go.html", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("./resources/img/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")

	myTemplates.ExecuteTemplate(w, "upload.success.go.html", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources/img/"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/img/pixelart.png
var imgPixel []byte

func TestUploadFormUnit(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Rafly Rigan Nagachi")

	file, _ := writer.CreateFormFile("file", "pixelart.png")
	file.Write(imgPixel)
	writer.Close()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "localhost:8080/upload", body)
	request.Header.Add("content-type", writer.FormDataContentType())

	Upload(recorder, request)

	response := recorder.Result()
	bodyResponse, _ := io.ReadAll(response.Body)
	fmt.Println(string(bodyResponse))
}
