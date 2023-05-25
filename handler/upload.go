package handler

import (
	"fmt"
	"io/ioutil"
	"kiedit/media"
	"kiedit/user"
	"kiedit/utils"
	"log"
	"net/http"
)

type UploadHandler struct {
}

var currentUser = new(user.User)

func processFileSegmentation(inputFile string) error {
	var splitVideoInput = media.SplitVideoInput{
		InputFile:     inputFile,
		Segment:       "10",
		OutputDirPath: currentUser.SessionDir + "/output%03d.mp4",
	}

	return media.SplitVideo(&splitVideoInput)
}

func processFileUpload(w http.ResponseWriter, r *http.Request) (string, error) {
	uploadDir := currentUser.SessionDir + "/upload"

	workingDir := []string{currentUser.SessionDir, "upload"}
	directoryUtil := new(utils.Directory)
	directoryUtil.CreateDirectory(workingDir)

	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		return "", err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(uploadDir, "upload-*.mp4")
	if err != nil {
		fmt.Println("Error creating a temp File")
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading uploaded file content")
		return "", err
	}
	tempFile.Write(fileBytes)

	data, _ := tempFile.Stat()

	filePath := uploadDir + "/" + data.Name()

	return filePath, nil
}

func uploadFile(w http.ResponseWriter, r *http.Request) error {
	filePath, err := processFileUpload(w, r)
	if err != nil {
		return err
	}

	if err := processFileSegmentation(filePath); err != nil {
		return err
	}

	return nil
}

func (self *UploadHandler) Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	currentUser.Init()

	if err := uploadFile(w, r); err != nil {
		fmt.Println("Error uploading file")
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}
