package main

import (
	"amadeus/pkg"
)

func main() {
	application := pkg.NewApplication()

	err := application.Start()
	if err != nil {
		panic(err)
	}

	// textLocation := []byte("test1.txt")
	// searchWord := "study"

	// // GET request
	// url := fmt.Sprintf("/api/v0.1/search/%v", searchWord)

	// bodyReader := bytes.NewReader(textLocation)

	// _, err = http.NewRequest("GET", url, bodyReader)
	// if err != nil {
	// 	log.Printf(err.Error())
	// }
}
