package view

import (
	"awesomeProject3/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type tokenS struct {
	Token string `json:"token"`
}
type test struct {
	Test string `json:"tests"`
}

type tokenValidator struct {
	Valid string `json:"isValid"`
}

func Registration(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	db, _ := models.ConnectDB()
	st := models.New(db)
	var requestData models.SimpleUserData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestData.UserName
	password := requestData.Password
	name := requestData.Name

	data := tokenS{
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
	data = tokenS{
		Token: token,
	}

	result, _ := json.Marshal(data)
	w.Write(result)

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

// Generate a JWT token
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

	data := tokenValidator{
		Valid: isValid,
	}
	//fmt.Println(GetRawAuthToken(r))
	res, _ := json.Marshal(data)
	w.Write(res)
}

// Validate the JWT token
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

// Middleware function to authorize user
func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := GetRawAuthToken(r)
		fmt.Println(r.URL.Path)

		if r.URL.Path == "/auth/validate-token" {
			ValidateTokenC(w, r)
			return
		}
		if r.URL.Path == "/drive/tests" {
			TestDownload(w, r)
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
			} else if r.URL.Path == "/auth/tests" {
				AuthTest(w, r)
				return
			}

			data := tokenS{
				Token: "failed",
			}
			res, _ := json.Marshal(data)
			w.Write(res)
			return
		}
		userID, err := FindUserIDByBearer(tokenString)
		if err != nil {
			data := tokenS{
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
		data := tokenS{
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
		data := tokenS{
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
	data := tokenS{
		Token: token,
	}
	result, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(string(result))
	w.Write(result)
}

func AuthTest(w http.ResponseWriter, r *http.Request) {
	n := test{
		Test: "tests",
	}
	f, _ := json.Marshal(n)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
}
