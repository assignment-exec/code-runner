package server

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"coderunner/constants"
	"coderunner/environment"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"
)

type assignmentTestingInformation struct {
	CommandToExecute string
	CommandToCompile string
	WorkDir          string
	RootDir          string
	Output           string
	CmdlineArgs      map[string]string
}

var assignTestingInfo assignmentTestingInformation

func getSupportedLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	language := os.Getenv(environment.SupportedLanguage)

	response, err := json.Marshal(language)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

// upload parses the client request and uploads the file.
func upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	output := "File uploaded successfully"
	err := readFormData(r)
	if err != nil {
		output = err.Error()
	}

	response, err := json.Marshal(output)
	if err != nil {
		log.Println(err)
	}

	// Write the response to be sent to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Builds the assignment uploaded.
func build(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var outputString string
	var currDir string
	var err error

	// Read the command to compile.
	assignTestingInfo.CommandToCompile = r.FormValue(constants.CompileCmdKey)

	// Navigate to the assignment working directory.
	currDir, err = navigateToWorkDir()
	if err != nil {
		outputString = err.Error()
	} else {
		// Execute the compile command.
		outputString, err = runCommand(assignTestingInfo.CommandToCompile)
		if err != nil {
			log.Println("error while building the assignment", err)
			outputString += err.Error()
		}

		// Navigate back to the code-runner working directory after successful execution.
		errChdir := os.Chdir(currDir)
		if errChdir != nil {
			log.Println("error while navigating to the current directory", errChdir)
			outputString += "\nError while navigating to the current directory"
		}
	}

	if err == nil {
		outputString = fmt.Sprintf("Compiled the assignment successfully \n\n %s", outputString)
	}
	response, err := json.Marshal(outputString)
	if err != nil {
		log.Println(err)
	}
	// Write the response to be sent to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Runs the assignment uploaded.
func run(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var outputString string
	var currDir string
	var err error

	// Read the command to run.
	assignTestingInfo.CommandToExecute = r.FormValue(constants.RunCmdKey)

	// Navigate to the assignment working directory.
	currDir, err = navigateToWorkDir()
	if err != nil {
		outputString = err.Error()
	} else {
		// Append the command line arguments to run command.
		runCmd := assignTestingInfo.CommandToExecute
		for key, value := range assignTestingInfo.CmdlineArgs {
			runCmd = fmt.Sprintf("%s %s %s", runCmd, key, value)
		}
		fmt.Println(runCmd)
		// Execute the assignment run command.
		outputString, err = runCommand(runCmd)
		if err != nil {
			log.Println("error while executing the assignment", err)
			outputString += err.Error()
		}

		// Navigate back to the code-runner working directory after successful execution.
		errChDir := os.Chdir(currDir)
		if errChDir != nil {
			log.Println("error while navigating to the current directory", errChDir)
			outputString += "\nError while navigating to the current directory"
		}
	}

	response, err := json.Marshal(outputString)
	if err != nil {
		log.Println(err)
	}
	// Write the response to be sent to client.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// navigateToWorkDir navigates to the provided working directory of the assignment.
func navigateToWorkDir() (string, error) {
	workDir := filepath.Join(assignTestingInfo.RootDir, assignTestingInfo.WorkDir)
	currDir, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "error while navigating to the working directory")
	}
	err = os.Chdir(filepath.Join(currDir, constants.AssignmentsDir, workDir))
	if err != nil {
		return "", errors.Wrap(err, "error while navigating to the working directory")
	}
	return currDir, nil
}

// runCommand runs the provided command.
func runCommand(cmdStr string) (string, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	var output string
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		output = fmt.Sprintf("%v\n%v", out.String(), stderr.String())
		return output, err
	}

	output = fmt.Sprintf("%v", out.String())
	return output, nil
}

// readFormData reads the compressed assignment submission and extracts the contents.
func readFormData(r *http.Request) error {
	fileHeader := make([]byte, 512)

	// Get the first file for the given key 'file'.
	file, handler, err := r.FormFile(constants.FormFileKey)
	if err != nil {
		return errors.Wrap(err, "error retrieving the file")
	}

	// Get the command line arguments.
	assignTestingInfo.CmdlineArgs = make(map[string]string)
	fmt.Println(r.Form)
	index := 1
	keyName := fmt.Sprintf("%s%d", constants.CmdArgKeyName, index)
	for r.FormValue(keyName) != "" {
		key := r.FormValue(keyName)

		argName := fmt.Sprintf("%s%d", constants.CmdArgValueName, index)
		arg := r.FormValue(argName)

		assignTestingInfo.CmdlineArgs[key] = arg
		index = index + 1
		keyName = fmt.Sprintf("%s%d", constants.CmdArgKeyName, index)
	}

	// Read the working directory, command to compile and command to run.
	assignTestingInfo.CommandToCompile = r.FormValue(constants.CompileCmdKey)
	assignTestingInfo.CommandToExecute = r.FormValue(constants.RunCmdKey)
	assignTestingInfo.WorkDir = r.FormValue(constants.WorkDirKey)

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}()

	if _, err := file.Read(fileHeader); err != nil {
		return errors.Wrap(err, "error retrieving the file")
	}

	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", http.DetectContentType(fileHeader))

	// Decompress the file and return its response.
	return decompressFile(file, fileHeader, handler)
}

// decompressFile reads and stores all files from the uploaded compressed file.
func decompressFile(file multipart.File, fileHeader []byte, handler *multipart.FileHeader) error {

	// Read the file based on the type of file compression.
	assignTestingInfo.RootDir = strings.TrimSuffix(handler.Filename, path.Ext(handler.Filename))

	if http.DetectContentType(fileHeader) == constants.ZipMimeFileType {
		// Read zip file.
		unZipped, err := zip.NewReader(file, handler.Size)
		if err != nil {
			return errors.Wrap(err, "error in unzipping file")
		}
		return storeUnzippedFiles(unZipped)

	} else if http.DetectContentType(fileHeader) == constants.TarGzMimeFileType {
		// Read tar.gz file.
		assignTestingInfo.RootDir = strings.TrimSuffix(assignTestingInfo.RootDir,
			path.Ext(assignTestingInfo.RootDir))
		fileReader, err := handler.Open()
		gZipReader, err := gzip.NewReader(fileReader)
		if err != nil {
			return errors.Wrap(err, "error in untaring file")
		}
		unTarred := tar.NewReader(gZipReader)
		return storeUnTarredFiles(unTarred)

	} else {
		// Read tar file.
		var fileReader io.ReadCloser = file
		unTarred := tar.NewReader(fileReader)
		return storeUnTarredFiles(unTarred)
	}
}

// storeUnTarredFiles stores unTared files to 'assignments/<tarball_name>' directory.
func storeUnTarredFiles(unTarred *tar.Reader) error {

	dest := filepath.Join(constants.AssignmentsDir, assignTestingInfo.RootDir)
	for {
		header, err := unTarred.Next()
		if err == io.EOF {
			// End of tar file.
			break
		}
		if err != nil {
			return errors.Wrap(err, "error in untaring")
		}

		filename := header.Name
		switch header.Typeflag {
		case tar.TypeDir:
			err := os.MkdirAll(filepath.Join(dest, filename), os.FileMode(header.Mode))
			if err != nil {
				return errors.Wrap(err, "error in untaring")
			}

		case tar.TypeReg:
			err := os.MkdirAll(filepath.Join(dest, filepath.Dir(filename)), os.FileMode(header.Mode))
			writer, err := os.Create(filepath.Join(dest, filename))
			if err != nil {
				return errors.Wrap(err, "error in untaring")
			}

			_, err = io.Copy(writer, unTarred)
			if err != nil {
				return errors.Wrap(err, "error in untaring")
			}

			err = os.Chmod(filepath.Join(dest, filename), os.FileMode(header.Mode))

			if err != nil {
				return errors.Wrap(err, "error in untaring")
			}

			writer.Close()
		default:
			return errors.Wrap(err, "error in untaring")
		}
	}
	return nil
}

// storeUnzippedFiles stores unzipped files to 'assignments/<tarball_name>' directory.
func storeUnzippedFiles(unZipped *zip.Reader) error {
	dest := filepath.Join(constants.AssignmentsDir, assignTestingInfo.RootDir)

	for _, file := range unZipped.File {
		fPath := filepath.Join(dest, file.Name)

		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return errors.New("error in unzipping")
		}

		if file.FileInfo().IsDir() {
			err := os.MkdirAll(fPath, os.ModePerm)
			if err != nil {
				return errors.Wrap(err, "error in unzipping")
			}
			continue
		}

		err := os.MkdirAll(filepath.Dir(fPath), os.ModePerm)
		if err != nil {
			return errors.Wrap(err, "error in unzipping")
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return errors.Wrap(err, "error in unzipping")
		}

		fileReader, err := file.Open()
		if err != nil {
			return errors.Wrap(err, "error in unzipping")
		}

		_, err = io.Copy(outFile, fileReader)

		// Close the file without defer to close before next iteration of loop.
		outFile.Close()
		fileReader.Close()

		if err != nil {
			return errors.Wrap(err, "error in unzipping")
		}

	}
	return nil
}

// listenAndServe listens to requests on the given port number.
func listenAndServe(wg *sync.WaitGroup, server *http.Server) {
	defer wg.Done()

	http.Handle("/", http.FileServer(http.Dir("./client")))
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

// StartServer starts service at given port.
func StartServer(port string) {

	var wg sync.WaitGroup

	http.HandleFunc("/getSupportedLanguage", getSupportedLanguage)
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/build", build)
	http.HandleFunc("/run", run)
	wg.Add(1)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	server := &http.Server{Addr: ":" + port}

	go listenAndServe(&wg, server)
	log.Printf("** Service Started on Port " + port + " **")
	<-sig
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	wg.Wait()
}
