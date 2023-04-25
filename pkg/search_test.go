package pkg

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	textLocation := "../test1.txt"
	searchWord := "study"

	application := NewApplication()

	// GET request
	url := fmt.Sprintf("/api/v0.1/search/%v?location=%v", searchWord, textLocation)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
	}

	// NewRecorder implements a ResponseWriter for testing
	responseWriter := httptest.NewRecorder()

	// GET response
	application.router.ServeHTTP(responseWriter, request)

	if status := responseWriter.Code; status != http.StatusOK {
		t.Error(status)
	} else {
		log.Print(responseWriter.Body.String())
	}
}
