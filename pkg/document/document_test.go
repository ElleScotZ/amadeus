package document

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	textLocation := []byte("../../test1.txt")
	searchWord := "study"

	router := Routes()

	// GET request
	url := fmt.Sprintf("/search/%v", searchWord)

	bodyReader := bytes.NewReader(textLocation)

	request, err := http.NewRequest("GET", url, bodyReader)
	if err != nil {
		t.Error(err)
	}

	// NewRecorder implements a ResponseWriter for testing
	responseWriter := httptest.NewRecorder()

	// GET response
	router.ServeHTTP(responseWriter, request)

	if status := responseWriter.Code; status != http.StatusOK {
		t.Error(status)
	} else {
		log.Print(responseWriter.Body.String())
	}
}
