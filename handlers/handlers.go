package handlers

import (
	drive "awesomeProject3/models"
	"context"
	"errors"
)

func ValidateUserData(name string, username string, password string) bool {
	db, _ := drive.ConnectDB()
	st := drive.New(db)
	ctx := context.Background()
	u, _ := st.ListUsers(ctx)
	if len(u) == 0 {
		return true
	}
	for _, usr := range u {
		if username == usr.UserName {
			return false
		}
	}
	if len(password) <= 6 || len(name) == 0 || len(username) < 4 {
		return false
	}
	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func FindUserByUsername(username string) (int64, error) {
	db, _ := drive.ConnectDB()
	st := drive.New(db)
	ctx := context.Background()
	u, _ := st.ListUsers(ctx)
	for _, usr := range u {
		if username == usr.UserName {
			return int64(usr.UserID), nil
		}
	}
	return 0, errors.New("error")
}

func CheckUserPass(userID int64, password string) bool {
	db, _ := drive.ConnectDB()
	st := drive.New(db)
	ctx := context.Background()
	u, _ := st.ListUsers(ctx)
	for _, usr := range u {
		if userID == usr.UserID {
			if usr.PassWordHash == password {
				return true
			}
		}
	}
	return false
}

func FindUserByUserID(userID int64) (string, error) {
	db, _ := drive.ConnectDB()
	st := drive.New(db)
	ctx := context.Background()
	u, _ := st.ListUsers(ctx)
	for _, usr := range u {
		if userID == usr.UserID {
			return usr.UserName, nil
		}
	}
	return "", errors.New("error")

}

func GetNameByUserID(userID int64) (string, error) {
	db, _ := drive.ConnectDB()
	st := drive.New(db)
	ctx := context.Background()
	u, _ := st.ListUsers(ctx)
	for _, usr := range u {
		if userID == usr.UserID {
			return usr.Name, nil
		}
	}
	return "", errors.New("error")

}
