package pkg

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Search struct {
}

// Routes collects all handlers and returns a router (mux).
// In this specific case, there is only 1 handler.
func (s *Search) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/{searchWord}", func(r chi.Router) {
		r.Get("/", s.GetAll)
	})

	return router
}

// GetAll is a GET handler that receives a GET request with a text file location in the request body,
// and a search word in its URL.
// If the search word is shorter than 2 characters, it returns StatusBadRequest.
// In case of StatusOK, the response contains a JSON as indicated in the task description (see task.pdf).
func (s *Search) GetAll(writer http.ResponseWriter, request *http.Request) {
	var (
		wordFound           bool
		numberOfOccurrences int
		lineOfOccurrences   []int
		lineCounter         = 1 // it is more intuitive to start counting the lines in the text file from 1
	)

	searchWord := chi.URLParam(request, "searchWord")

	// Status 400 in case of too short searchWord
	if len(searchWord) < 2 {
		http.Error(writer, "GetAll request for minimum 2 letters, please!", http.StatusBadRequest)
		return
	}

	// Extract text file location information from query
	fileLocation := request.URL.Query().Get("location")

	// Opening the text file
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Scanning through the text file
	for scanner.Scan() {
		line := scanner.Text()

		if numberOfOccurrencesIn1Line := strings.Count(line, searchWord); numberOfOccurrencesIn1Line > 0 {
			wordFound = true

			numberOfOccurrences += numberOfOccurrencesIn1Line

			for k := 0; k < numberOfOccurrencesIn1Line; k++ {
				lineOfOccurrences = append(lineOfOccurrences, lineCounter)
			}
		}

		lineCounter++
	}

	if err := scanner.Err(); err != nil {
		log.Printf("GetAll has failed: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// Writing response in JSON
	response := []byte(fmt.Sprintf(`{
		"wordFound": %v,
		"numOccurrences": %v,
		"lineOccurrences": %v
		}`, wordFound, numberOfOccurrences, lineOfOccurrences))

	writer.Header().Set("Content-Type", "text/html")

	_, err = writer.Write(response)
	if err != nil {
		log.Printf("GetAll has failed: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
