package view

import (
	"encoding/json"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	passwordFTP = "yourPass"
	userNameFTP = "yourName"
	addressFTP  = "localhost:21"
)

func PrintFile(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	path := r.PostFormValue("path")
	fileName := r.PostFormValue("fileName")

	err := c.ChangeDir("data/" + userID + "/" + path)
	if err != nil {
		fmt.Fprintf(w, "Path is not correct")
		return
	}
	retr, err := c.Retr(fileName)
	if err != nil {
		fmt.Fprintf(w, "File not exist")
		return
	}

	defer retr.Close()

	test, err := io.ReadAll(retr)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(test)

}

func AddFileToFolder(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	path := r.PostFormValue("path")
	fileName := r.PostFormValue("fileName")

	err := c.ChangeDir("data/" + userID + "/" + path)
	if err != nil {
		fmt.Fprintf(w, "Path is not correct")
		return
	}

	file, _ := os.OpenFile("Files/"+fileName, os.O_RDONLY, 0)
	err = c.StorFrom(fileName, file, 0)
	if err != nil {
		fmt.Println("File error")
		return
	}
	dirPath, _ := c.CurrentDir()
	files, _ := c.NameList("")
	data := UserFolder{Data: files, Path: dirPath}
	jdata, _ := json.Marshal(data)
	w.Write(jdata)
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

func CreateFolder(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	name := r.PostFormValue("folderName")
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	c.ChangeDir("data/" + userID)
	c.MakeDir(name)
	w.Write([]byte("Folder " + name + " created"))
}

type UserFolder struct {
	Data []string `json:"data"`
	Path string   `json:"path"`
}

func GetUserFolder(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID := GetValidUserID(r, c)
	path := r.PostFormValue("path")
	err := c.ChangeDir("data/" + userID + "/" + path)
	if err != nil {
		fmt.Fprintf(w, "Folder is not exist !")
		return
	}
	files, _ := c.NameList("")
	dirPath, _ := c.CurrentDir()
	data := UserFolder{Data: files, Path: dirPath}
	jdata, _ := json.Marshal(data)
	w.Write(jdata)
}

func AuthFTP(addr string, user string, pass string) *ftp.ServerConn {
	c, err := ftp.Dial(addr, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(user, pass)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

func GetValidUserID(r *http.Request, c *ftp.ServerConn) string {
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	n, _ := c.NameList("")
	if !contains(n, userID) {
		CreateUserFolder(userID)
	}
	return userID
}

func CreateUserFolder(userID string) {
	c, err := ftp.Dial("localhost:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("yourName", "yourPass")
	if err != nil {
		fmt.Println("test")
		log.Fatal(err)
	}
	c.ChangeDir("data")
	c.MakeDir(userID)
}
