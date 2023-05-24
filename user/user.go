package user

import (
	utils "kiedit/utils"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	SessionID  string
	SessionDir string
}

func (self *User) Init() {
	self.SessionID = generateSessionID()
	self.SessionDir = createSessionDirectory(self.SessionID)
}

func generateSessionID() string {
	return uuid.NewV4().String()
}

func createSessionDirectory(sessionID string) string {
	sessionDir := []string{"./dist", sessionID}
	directoryUtil := new(utils.Directory)

	dir, err := directoryUtil.CreateDirectory(sessionDir)

	if err != nil {
		panic(err)
	}

	return dir
}
