package server

import (
	"bytes"
	"coderunner/constants"
	"flag"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

var portNumber = flag.String("port", "52453", "Port number for server to listen on")

// Creates a new file upload http request with params as cmd arguments.
func newRequest(uri string, params map[string]string, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		err = file.Close()
	}()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(constants.FormFileKey, filepath.Base(path))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}

// Sends the constructed request to server.
func sendRequest(filePath string, port string) (*http.Response, error) {
	params := map[string]string{
		"key1": "Key1",
		"arg1": "100",
		"key2": "Key2",
		"arg2": "200",
	}

	request, err := newRequest("http://localhost:"+port+"/upload", params, filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	return resp, err
}

// Sends request for a zip file. For testing purpose only.
func sendZip(port string) (*http.Response, error) {
	return sendRequest("test-archive_1.zip", port)
}

// Sends request for a tar file. For testing purpose only.
func sendTar(port string) (*http.Response, error) {
	return sendRequest("test-archive_2.tar", port)
}

// Sends request for a tar.gz file. For testing purpose only.
func sendTarGz(port string) (*http.Response, error) {
	return sendRequest("test-archive_3.tar.gz", port)
}

func TestSendZip(t *testing.T) {
	port := *portNumber
	_, err := sendZip(port)
	assert.Nil(t, err)
}

func TestSendTar(t *testing.T) {
	port := *portNumber
	_, err := sendTar(port)
	assert.Nil(t, err)
}

func TestSendTarGz(t *testing.T) {
	port := *portNumber
	_, err := sendTarGz(port)
	assert.Nil(t, err)
}
