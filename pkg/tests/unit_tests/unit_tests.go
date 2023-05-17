package unit_tests

import (
	"awesomeProject3/pkg/models"
	"testing"
)

func TestDbConnection(t *testing.T) {
	_, err := models.ConnectDB()
	if err != nil {
		t.Fatalf("Cannot connect to Data Base")
	}
}
