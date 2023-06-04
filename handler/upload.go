package handler

import (
	"io/ioutil"
	"kiedit/media"
	"kiedit/queue"
	"kiedit/user"
	"kiedit/utils"
	"log"
	"net/http"
)

type UploadHandler struct {
}

var currentUser = new(user.User)

func (self *UploadHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	currentUser.Init()

	var filePath, err = processUploadFile(w, r)

	if err != nil {
		log.Println("Error uploading file")
		log.Fatal(err)
	}

	var splitVideoInput = media.SplitVideoInput{
		InputFile:     filePath,
		Segment:       "30",
		OutputDirPath: currentUser.SessionDir + "/output%03d.mp4",
	}

	if err := addFileToQueue(splitVideoInput); err != nil {
		log.Println("Error adding file to queue")
		log.Fatal(err)
	}

	log.Println("Successfully Uploaded File\n")
}

func processUploadFile(w http.ResponseWriter, r *http.Request) (string, error) {
	uploadDir := currentUser.SessionDir + "/upload"

	workingDir := []string{currentUser.SessionDir, "upload"}
	directoryUtil := new(utils.Directory)
	directoryUtil.CreateDirectory(workingDir)

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("Error Retrieving the File")
		return "", err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile(uploadDir, "upload-*.mp4")
	if err != nil {
		log.Println("Error creating a temp File")
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Error reading uploaded file content")
		return "", err
	}
	tempFile.Write(fileBytes)

	data, _ := tempFile.Stat()

	filePath := uploadDir + "/" + data.Name()

	return filePath, nil
}

func addFileToQueue(splitVideoInput media.SplitVideoInput) error {
	queue := new(queue.QueueStruct)
	if err := queue.Connect(); err != nil {
		return err
	}
	defer queue.Close()

	if err := queue.Publish(splitVideoInput); err != nil {
		return err
	}

	return nil
}
