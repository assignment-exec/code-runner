package server

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

// Parses the client request and uploads the file.
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// Give a max limit for file upload.
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Returns the first file for the given key 'file'.
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		log.Fatal(err)
		return
	}

	// Read the command line arguments (additional parameters passed).
	for index := 1;index <= len(r.Form); index++ {
		argName := fmt.Sprintf("%s%d","arg", index)
		arg := r.FormValue(argName)
		fmt.Println(arg)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	fileHeader := make([]byte,512)
	if _, err := file.Read(fileHeader)
		err != nil {
			log.Fatal(err)
		}

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", http.DetectContentType(fileHeader))

	// Based on the type of file compression, read the file.
	if http.DetectContentType(fileHeader) == "application/octet-stream" {	// For tar or tar.gz file.
		unTared := tar.NewReader(file)
		for {
			header, err := unTared.Next()
			if err == io.EOF {
				break // End of tar file.
			}
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(header.Name)

		}
	} else {	// For zip file.
		unZipped, _ := zip.NewReader(file,handler.Size)
		for _,f := range unZipped.File {
			fmt.Println(f.Name)
		}
	}

	// Write status to the response to be sent to client.
	w.Header().Add("Content-Type", "application/json")
	_,err = io.WriteString(w, `{"Status":"Successfully Uploaded File(s)"}`)
	if err != nil {
		log.Fatal(err)
	}
}

func listenAndServe(wg *sync.WaitGroup) {
	defer wg.Done()
	port := "8082"

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func StartServer() {

	var wg sync.WaitGroup

	// Start service at given port.
	http.HandleFunc("/upload", upload)

	wg.Add(1)
	go listenAndServe(&wg)
	wg.Wait()
}
