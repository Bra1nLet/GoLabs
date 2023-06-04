package handlers

import (
	"awesomeProject3/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, _ := models.ConnectDB()
	st := models.New(db)

	u, _ := st.ListUsers(context.Background())

	for _, usr := range u {
		w.Write([]byte(fmt.Sprintf("Name : %s, ID : %d\n", usr.UserName, usr.UserID)))
	}
}

func GetRawAuthToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	userID, _ := FindUserIDByBearerAsInt(GetRawAuthToken(r))
	userName, _ := FindUserByUserID(userID)
	Name, _ := GetNameByUserID(userID)

	userData := models.SimpleUserData{
		UserID:   userID,
		UserName: userName,
		Name:     Name,
	}

	n, _ := json.Marshal(userData)
	w.WriteHeader(http.StatusOK)
	w.Write(n)
}

func FindUserIDByBearerAsInt(tokenString string) (int64, error) {
	token, _ := FindUserIDByBearer(tokenString)
	result, _ := strconv.Atoi(token)
	return int64(result), nil
}

func FindUserIDByBearer(tokenString string) (string, error) {
	token, err := validateToken(strings.TrimPrefix(tokenString, "Bearer "))
	if err != nil || !token.Valid {
		return "", err
	}
	userID, ok := token.Claims.(jwt.MapClaims)["userID"].(string)
	if !ok {
		return "", err
	}
	return userID, nil
}

func hasPermission(userID int64, url string) bool {
	// Check is user have permissions
	return true
}

// Example handler using middleware
func HandleProtectedData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Data protected \nIt's Fine !"))
}
