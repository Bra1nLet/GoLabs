package view

import (
	drive "awesomeProject3/pkg/models"
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func DockerGO() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
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
