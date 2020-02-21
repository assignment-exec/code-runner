package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

// Creates a new file upload http request with params as cmd arguments.
func newRequest(uri string, fileKeyName string, params map[string]string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileKeyName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

// Sends the constructed request to server.
func sendRequest(filePath string) (*http.Response, error) {
	params := map[string]string{
		"arg1": "100",
		"arg2": "200",
	}

	request, err := newRequest("http://localhost:8081/upload", "file", params, filePath)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	return resp, err
}

// Sends request for a zip file. For testing purpose only.
func sendZip() (*http.Response, error) {
	return  sendRequest("/home/bhargavi/test_archives/final-balandi1-master.zip")
}

// Sends request for a tar file. For testing purpose only.
func sendTar() (*http.Response, error) {
	return  sendRequest("/home/bhargavi/test_archives/p1-balandi1.tar")
}

// Sends request for a tar.gz file. For testing purpose only.
func sendTarGz() (*http.Response, error) {
	return  sendRequest("/home/bhargavi/test_archives/final.tar.gz")
}

func TestSendZip(t *testing.T) {
	_, err := sendZip()
	assert.Equal(t,err,nil)
}

func TestSendTar(t *testing.T) {
	_, err := sendTar()
	assert.Equal(t,err,nil)
}

func TestSendTarGz(t *testing.T) {
	_, err := sendTarGz()
	assert.Equal(t,err,nil)
}