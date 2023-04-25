package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Result struct {
	WordFound       bool  `json:"wordFound"`
	NumOccurrences  int   `json:"numOccurrences"`
	LineOccurrences []int `json:"lineOccurrences"` // it can never be nil after json.Unmarshal()
}

func TestGetAll(t *testing.T) {
	application := NewApplication()

	// Case 1: no occurrence
	{
		const (
			textLocation = "../test1.txt"
			searchWord   = "falatka"
		)

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
			var resultObject Result

			err = json.Unmarshal(responseWriter.Body.Bytes(), &resultObject)
			if err != nil {
				t.Error(err)
			}

			if resultObject.WordFound {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.WordFound, !resultObject.WordFound)
			}

			if resultObject.NumOccurrences != 0 {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.NumOccurrences, 0)
			}

			if len(resultObject.LineOccurrences) != 0 {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.LineOccurrences, 0)
			}
		}
	}

	// Case 2: one occurrence
	{
		const (
			textLocation = "../test1.txt"
			searchWord   = "which"
		)

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
			var resultObject Result

			err = json.Unmarshal(responseWriter.Body.Bytes(), &resultObject)
			if err != nil {
				t.Error(err)
			}

			if !resultObject.WordFound {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.WordFound, !resultObject.WordFound)
			}

			if resultObject.NumOccurrences != 1 {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.NumOccurrences, 1)
			}

			if len(resultObject.LineOccurrences) != 1 || resultObject.LineOccurrences[0] != 1 {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.LineOccurrences, []int{1})
			}
		}
	}

	// Case 3: multiple occurrences
	{
		textLocation := "../test1.txt"
		searchWord := "study"

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
			var resultObject Result

			err = json.Unmarshal(responseWriter.Body.Bytes(), &resultObject)
			if err != nil {
				t.Error(err)
			}

			if !resultObject.WordFound {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.WordFound, !resultObject.WordFound)
			}

			if resultObject.NumOccurrences != 3 {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.NumOccurrences, 3)
			}

			if len(resultObject.LineOccurrences) != 3 || resultObject.LineOccurrences[0] != 18 ||
				resultObject.LineOccurrences[1] != 22 || resultObject.LineOccurrences[2] != 23 {
				t.Errorf("TestGetAll has failed. Got %v instead of %v", resultObject.LineOccurrences, []int{18, 22, 23})
			}
		}
	}
}
