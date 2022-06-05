package main

import (
	"fmt"

	"github.com/jan-r-dev/go_course_full_mcleod_notes/section_27_exercises_12/dog"
)

func main() {
	dy := dog.Years(8)

	fmt.Println(dy)
}

// Tested local documentation with: godoc -http=:8080
