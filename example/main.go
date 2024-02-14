package main

import (
	"fmt"
	"strings"

	"github.com/zeevallin/dotenv"
)

func main() {
	r1 := strings.NewReader(`# File number one
	HELLO=world # This delcares the HELLO variable
	# Database config below
	DATABASE=postgres
	`)

	r2 := strings.NewReader(`# File number two
	FOO="bar" # This delcares the foo variable
	PLEASE_DELETE=me
	`)

	// read the files

	f1, err := dotenv.Read(r1)
	if err != nil {
		panic(err)
	}

	f2, err := dotenv.Read(r2)
	if err != nil {
		panic(err)
	}

	// // merge the files

	if err := f1.Merge(f2); err != nil {
		panic(err)
	}

	// // set extra values

	f1.Set("MY_EXTRA_VAR", `"with a value"`, "")
	f1.Delete("PLEASE_DELETE")

	fmt.Println(f1)
}
