package view

import (
	drive "awesomeProject3/pkg/models"
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
	if len(password) == 0 || len(name) == 0 {
		return false
	}
	return true

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
