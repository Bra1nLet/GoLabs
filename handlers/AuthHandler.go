package handlers

import (
	"awesomeProject3/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func generateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte("secret-key"))
}

func ValidateTokenC(w http.ResponseWriter, r *http.Request) {
	token, err := validateToken(strings.TrimPrefix(GetRawAuthToken(r), "Bearer "))
	isValid := "false"
	if err != nil || token == nil {
		isValid = "false"
	} else {
		isValid = "true"
	}

	data := models.TokenValidator{
		Valid: isValid,
	}

	res, _ := json.Marshal(data)
	w.Write(res)
}

func Registration(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	db, err := models.ConnectDB()
	if err != nil {
		log.Println("Unable connect to DataBase")
		w.WriteHeader(http.StatusBadRequest)
	}
	st := models.New(db)
	var requestData models.SimpleUserData
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Println(err)
		return
	}

	username := requestData.UserName
	password := requestData.Password
	name := requestData.Name

	data := models.TokenS{
		Token: "failed",
	}
	if !ValidateUserData(name, username, password) {
		w.WriteHeader(http.StatusBadRequest)
		result, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return
	}

	_, _ = st.CreateUsers(context.Background(), models.CreateUsersParams{
		UserName:     username,
		PassWordHash: password,
		Name:         name,
	})

	UserID, _ := FindUserByUsername(username)
	token, _ := generateToken(strconv.FormatInt(UserID, 10))
	data = models.TokenS{
		Token: token,
	}

	result, _ := json.Marshal(data)
	w.Write(result)

}

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetRawAuthToken(r)
		fmt.Println(r.URL.Path)

		if r.URL.Path == "/auth/validate-token" {
			ValidateTokenC(w, r)
			return
		}
		if r.URL.Path == "/hello" {
			Hello(w, r)
			return
		}
		if r.URL.Path == "/auth/new-token" {
			GenerateNewToken(w, r)
			return
		}
		if tokenString == "" {
			if r.URL.Path == "/auth/new-account" {
				Registration(w, r)
				return
			}

			data := models.TokenS{
				Token: "failed",
			}
			res, _ := json.Marshal(data)
			w.Write(res)
			return
		}
		userID, err := FindUserIDByBearer(tokenString)
		if err != nil {
			data := models.TokenS{
				Token: "failed",
			}
			res, _ := json.Marshal(data)
			w.Write(res)
			return
		}

		intUserID, _ := strconv.ParseInt(userID, 10, 64)
		// Check user permissions
		if !hasPermission(intUserID, r.URL.Path) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		// Call next handler
		next.ServeHTTP(w, r)
	})
}

func GenerateNewToken(w http.ResponseWriter, r *http.Request) {

	db, err := models.ConnectDB()
	st := models.New(db)
	st.ListUsers(context.Background())

	var requestData models.SimpleUserData
	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestData.UserName
	password := requestData.Password

	fmt.Println(username + " " + password)

	userID, err := FindUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		data := models.TokenS{
			Token: "failed",
		}
		result, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return
	}

	fmt.Println(userID)

	if !CheckUserPass(userID, password) {
		w.WriteHeader(http.StatusBadRequest)
		data := models.TokenS{
			Token: "failed",
		}
		result, _ := json.Marshal(data)
		w.Write(result)
		return
	}

	token, err := generateToken(strconv.FormatInt(userID, 10))
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	//fmt.Println(username, password, userID)
	data := models.TokenS{
		Token: token,
	}
	result, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(result))
	w.Write(result)
}
