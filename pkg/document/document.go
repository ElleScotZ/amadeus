package document

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
)

// Routes
func Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/{searchWord}", func(r chi.Router) {
		r.Get("/", GetAll)
	})

	return router
}

// GetAll
func GetAll(writer http.ResponseWriter, request *http.Request) {
	var (
		wordFound           bool
		numberOfOccurrences int
		lineOfOccurrences   []int
		lineCounter         = 1
	)

	searchWord := chi.URLParam(request, "searchWord")

	if len(searchWord) < 2 {
		http.Error(writer, "GetAll request for minimum 2 letters, please!", http.StatusBadRequest)
		return
	}

	// Reading request body that contains the location of the text file
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("GetAll has failed: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	fileLocation := string(requestBody)

	// Scanning through text file
	file, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

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

	// Writing response
	response := []byte(fmt.Sprintf(`{
		"wordFound": %v,
		"numOccurrences": %v,
		"lineOccurrences": %v,
		}`, wordFound, numberOfOccurrences, lineOfOccurrences))

	writer.Header().Set("Content-Type", "application/json")

	_, err = writer.Write(response)
	if err != nil {
		log.Printf("GetAll has failed: %v", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
