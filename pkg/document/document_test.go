package document

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	textLocation := []byte("test1.txt")
	searchWord := "bla"

	router := Routes()

	// GET request
	url := fmt.Sprintf("/%v", searchWord)

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
	}
}
