package server

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"coderunner/constants"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type assignmentData struct {
	CommandToExecute string
	CommandToCompile string
	WorkDir          string
	RootDir          string
	Output           string
	CmdlineArgs      map[string]string
}

var AssignmentData assignmentData

// Parses the client request and uploads the file.
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	responseString := readFormData(r)

	if len(responseString) <= 0 {
		responseString += `"Upload Status":"Successfully Uploaded File(s)"`
	}

	responseS, err := json.Marshal(responseString)
	if err != nil {
		log.Println(err)
	}

	// Write the response to be sent to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseS)
}

// Builds and runs the assignment uploaded.
func buildRun(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Build and Run Endpoint Hit")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	workDir := filepath.Join(AssignmentData.RootDir, AssignmentData.WorkDir)
	currDir, _ := os.Getwd()
	err := os.Chdir(filepath.Join(currDir, constants.AssignmentsDir, workDir))
	if err != nil {
		log.Fatalf("error while navigating to the working directory: %v", err)
	}
	var outputString string

	outputString, err = runCommand(AssignmentData.CommandToCompile)

	if err == nil {
		outputString, err = runCommand(AssignmentData.CommandToExecute)
	}

	responseS, err := json.Marshal(outputString)
	if err != nil {
		log.Println(err)
	}

	err = os.Chdir(currDir)
	if err != nil {
		log.Fatalf("error while navigating to the working directory: %v", err)
	}

	// Write the response to be sent to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseS)
}

func build() {

}

func run() {

}
// Runs the provided command.
func runCommand(cmdStr string) (string,error) {
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Sprintf("%v", stderr.String()),err
	}

	return fmt.Sprintf("%v", out.String()),nil
}

// Reads the compressed file and invokes function to decompress it.
func readFormData(r *http.Request) string {
	fileHeader := make([]byte, 512)

	// Returns the first file for the given key 'file'.
	file, handler, err := r.FormFile(constants.FormFileKey)
	if err != nil {
		responseString := `"File Error":"Error in retrieving the file"`
		log.Println("error Retrieving the file", err)
		return responseString
	}

	AssignmentData.CmdlineArgs = make(map[string]string)
	for index := 1; index <= len(r.Form); index++ {
		keyName := fmt.Sprintf("%s%d", constants.CmdArgKeyName, index)
		key := r.FormValue(keyName)

		argName := fmt.Sprintf("%s%d", constants.CmdArgValueName, index)
		arg := r.FormValue(argName)
		AssignmentData.CmdlineArgs[key] = arg
	}

	// Read the working directory and command to run.
	AssignmentData.CommandToCompile = r.FormValue(constants.CompileCmdKey)
	AssignmentData.CommandToExecute = r.FormValue(constants.RunCmdKey)
	AssignmentData.WorkDir = r.FormValue(constants.WorkDirKey)

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	if _, err := file.Read(fileHeader); err != nil {
		log.Println(err)
	}

	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", http.DetectContentType(fileHeader))

	return decompressFile(file, fileHeader, handler)
}

// Reads all files from the uploaded compressed file.
func decompressFile(file multipart.File, fileHeader []byte, handler *multipart.FileHeader) string {
	// Based on the type of file compression, read the file.


	AssignmentData.RootDir = strings.TrimSuffix(handler.Filename, path.Ext(handler.Filename))
	if http.DetectContentType(fileHeader) == "application/octet-stream" {

		// For tar or tar.gz file.
		var fileReader io.ReadCloser = file
		unTarred := tar.NewReader(fileReader)
		return storeUnTarredFiles(unTarred)
	} else {
		// For zip file.
		unZipped, err := zip.NewReader(file, handler.Size)
		if err != nil {


			responseString := `"Unzip Error":"Error in unzipping uploaded file"`
			log.Println("error in unzipping file", err)
			return responseString
		}

		return storeUnzippedFiles(unZipped)
	}
}

// Stores unTared files to a folder.
func storeUnTarredFiles(unTarred *tar.Reader) string {

	errResponse := `"UnTar Error":"Error in un-tarring uploaded file"`
	dest := filepath.Join(constants.AssignmentsDir, AssignmentData.RootDir)
	for {
		header, err := unTarred.Next()
		if err == io.EOF {
			// End of tar file.
			break
		}
		if err != nil {
			log.Println("unTar error: ", err)
			return errResponse
		}

		filename := header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			err := os.MkdirAll(filepath.Join(dest, filename), os.FileMode(header.Mode)) // or use 0755 if you prefer
			if err != nil {
				log.Println("unTar error: ", err)
				return errResponse
			}

		case tar.TypeReg:
			err := os.MkdirAll(filepath.Join(dest, filepath.Dir(filename)), os.FileMode(header.Mode))
			writer, err := os.Create(filepath.Join(dest, filename))
			if err != nil {
				log.Println("unTar error: ", err)
				return errResponse
			}

			_, err = io.Copy(writer, unTarred)
			if err != nil {
				log.Println("unTar error: ", err)
				return errResponse
			}

			err = os.Chmod(filepath.Join(dest, filename), os.FileMode(header.Mode))

			if err != nil {
				log.Println("unTar error: ", err)
				return errResponse
			}

			writer.Close()
		default:
			log.Println("unable to unTar type : ", header.Typeflag)
			return errResponse
		}
	}
	return ""
}

// Stores unzipped files to a folder.
func storeUnzippedFiles(unZipped *zip.Reader) string {
	dest := filepath.Join(constants.AssignmentsDir, AssignmentData.RootDir)

	errorResponse := `"Unzip Error":"Error in unzipping uploaded file"`

	for _, file := range unZipped.File {
		fPath := filepath.Join(dest, file.Name)

		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			log.Println("unzip error: illegal filepath")
			return errorResponse
		}

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fPath, os.ModePerm)
			if err != nil {
				log.Println("unzip error: ", err)
				return errorResponse
			}
			continue
		}

		err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
		if err != nil {
			log.Println("unzip error: ", err)
			return errorResponse
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Println("unzip error: ", err)
			return errorResponse
		}

		rc, err := file.Open()
		if err != nil {
			log.Println("unzip error: ", err)
			return errorResponse
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop.
		outFile.Close()
		rc.Close()

		if err != nil {
			log.Println("unzip error: ", err)
			return errorResponse
		}

	}
	return ""
}
func listenAndServe(wg *sync.WaitGroup, port string) {
	defer wg.Done()

	log.Printf("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
func StartServer(port string) {

	var wg sync.WaitGroup

	// Start service at given port.
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/buildRun", buildRun)
	wg.Add(1)
	go listenAndServe(&wg, port)
	wg.Wait()
}
