package handlers

import (
	"awesomeProject3/models"
	"encoding/json"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	passwordFTP = "yourPass"
	userNameFTP = "yourName"
	addressFTP  = "ftp-server:21"
)

func PrintFile(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	path := r.PostFormValue("path")
	fileName := r.PostFormValue("fileName")
	fmt.Println(c.List(""))
	fmt.Println(c.CurrentDir())
	c.ChangeDir("")

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
	c.Quit()
}

func AddFileToFolder(file io.Reader, filename string, w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	path := r.PostFormValue("path")

	err := c.ChangeDir("data/" + userID + "/" + path)
	if err != nil {
		fmt.Fprintf(w, "Path is not correct")
		return
	}

	err = c.StorFrom(filename, file, 0)
	if err != nil {
		fmt.Println("File error")
		return
	}
	dirPath, _ := c.CurrentDir()
	folders, fileList := GetFilesFolders(c)
	data := models.UserFolder{Folders: folders, Files: fileList, Path: dirPath}
	jdata, _ := json.Marshal(data)
	w.Write(jdata)
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
	c.Quit()
}

func CreateFolder(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	name := r.PostFormValue("folderName")
	path := r.PostFormValue("path")
	userID, _ := FindUserIDByBearer(GetRawAuthToken(r))
	fmt.Println(c.CurrentDir())
	c.ChangeDir("data/" + userID + "/" + path)
	fmt.Println(c.CurrentDir())
	c.MakeDir(name)
	fmt.Println(c.CurrentDir())
	GetUserFolder(w, r)
	c.Quit()
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
	dirPath, _ := c.CurrentDir()

	folders, fileList := GetFilesFolders(c)

	data := models.UserFolder{Folders: folders, Files: fileList, Path: dirPath}
	jdata, _ := json.Marshal(data)
	w.Write(jdata)
	c.Quit()
}

func GetFilesFolders(c *ftp.ServerConn) ([]string, []string) {
	entries, err := c.List("")
	if err != nil {
		log.Fatal(err)
	}

	folders := make([]string, 0)
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFolder {
			folders = append(folders, entry.Name)
		}
	}

	fileslist := make([]string, 0)
	for _, entry := range entries {
		if entry.Type == ftp.EntryTypeFile {
			fileslist = append(fileslist, entry.Name)
		}
	}
	return folders, fileslist
}

func DeleteUserFile(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID := GetValidUserID(r, c)
	path := r.PostFormValue("path")
	filename := r.PostFormValue("filename")
	err := c.ChangeDir("data/" + userID + "/" + path)
	if err != nil {
		fmt.Fprintf(w, "Folder is not exist !")
		return
	}
	err = c.Delete(filename)
	if err != nil {
		err := c.RemoveDirRecur(filename)
		c.Quit()
		if err != nil {
			fmt.Fprintf(w, "File not found")
			return
		}
		fmt.Fprintf(w, "Folder deleted")
		return
	}

	dirPath, _ := c.CurrentDir()
	folders, fileList := GetFilesFolders(c)
	data := models.UserFolder{Files: fileList, Folders: folders, Path: dirPath}
	jdata, _ := json.Marshal(data)
	w.Write(jdata)
	c.Quit()
}

func RenameUserFile(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	userID := GetValidUserID(r, c)
	path := r.PostFormValue("path")
	filename := r.PostFormValue("filename")
	newname := r.PostFormValue("newname")
	err := c.ChangeDir("data/" + userID + "/" + path)
	if err != nil {
		fmt.Fprintf(w, "Folder is not exist !")
		return
	}
	err = c.Rename(filename, newname)
	if err != nil {
		fmt.Fprintf(w, "File is not found")
		return
	}
	dirPath, _ := c.CurrentDir()
	folders, fileList := GetFilesFolders(c)
	data := models.UserFolder{Files: fileList, Folders: folders, Path: dirPath}
	jdata, _ := json.Marshal(data)
	w.Write(jdata)
	c.Quit()
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
	c, err := ftp.Dial(addressFTP, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(userNameFTP, passwordFTP)
	if err != nil {
		log.Fatal(err)
	}
	err = c.ChangeDir("data")
	if err != nil {
		c.MakeDir("data")
		c.ChangeDir("data")
	}
	c.MakeDir(userID)
	c.ChangeDir(userID)
	c.MakeDir("init")
	c.Quit()
}

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	c := AuthFTP(addressFTP, userNameFTP, passwordFTP)
	filePath := r.PostFormValue("path")
	fileName := r.PostFormValue("filename")
	userID := GetValidUserID(r, c)

	c.ChangeDir("data/" + userID + "/" + filePath)
	n, _ := c.NameList("")
	log.Println(n)
	file, err := c.Retr(fileName)
	if err != nil {
		log.Println("Failed to retrieve file from FTP server:", err)
		http.Error(w, "Failed to retrieve file from FTP server", http.StatusInternalServerError)
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(fileBytes)))
	c.Quit()
	if _, err := w.Write(fileBytes); err != nil {
		http.Error(w, fmt.Sprintf("Failed to write response: %s", err), http.StatusInternalServerError)

		return
	}
}
