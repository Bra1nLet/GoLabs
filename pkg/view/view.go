package view

import (
	"awesomeProject3/pkg/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	db, _ := models.ConnectDB()
	st := models.New(db)

	UserName := r.PostFormValue("username")
	PassWordHash := r.PostFormValue("password")
	Name := r.PostFormValue("name")

	if !ValidateUserData(Name, UserName, PassWordHash) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong Data"))
		return
	}

	_, _ = st.CreateUsers(context.Background(), models.CreateUsersParams{
		UserName:     UserName,
		PassWordHash: PassWordHash,
		Name:         Name,
	})

	w.Write([]byte(UserName))

}

type test struct {
	Test string `json:"test"`
}

func AuthTest(w http.ResponseWriter, r *http.Request) {
	n := test{
		Test: "test",
	}
	f, _ := json.Marshal(n)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(f)
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
		if tokenString == "" {
			if r.URL.Path == "/auth/new-token" {
				GenerateNewToken(w, r)
				return
			} else if r.URL.Path == "/auth/new-account" {
				Registration(w, r)
				return
			} else if r.URL.Path == "/auth/test" {
				AuthTest(w, r)
				return
			}

			http.Error(w, "Unauthorized 1", http.StatusUnauthorized)
			return
		}
		userID, err := FindUserIDByBearer(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
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

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	userID, err := FindUserByUsername(username)
	if err != nil {
		w.Write([]byte("Wrong data"))
		return
	}

	token, err := generateToken(strconv.FormatInt(userID, 10))
	if err != nil {
		w.Write([]byte("error"))
		return
	}

	fmt.Println(username, password, userID)
	data := tokenS{
		Token: token,
	}
	result, _ := json.Marshal(data)
	w.Header().Set("application/json", http.MethodPost)
	fmt.Fprintln(w, string(result))
}

type tokenS struct {
	Token string `json:"token"`
}
