package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	js := `[{"First":"James","Last":"Bond","Age":32},{"First":"Jeff","Last":"Bridges","Age":12}]`
	f, err := os.Create("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(f)
	err = enc.Encode(js)
	if err != nil {
		log.Fatal(err)
	}
}
