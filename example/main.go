package main

import (
	"fmt"

	"github.com/zeevallin/dotenv/lexer2"
)

func main() {
	// r1 := `"  hello"\nworld "
	// Hello=World
	// `

	r1 := `
	REDIS_URL=redis://localhost:6379
	redis://localhost:6379=FOOBAR
	WEBSITE_TITLE=My 2nd Website
	`

	lexer := lexer2.New("r1.env", r1)
	for token := lexer.NextToken(); token.Kind != lexer2.EOF; token = lexer.NextToken() {
		if token.Kind == lexer2.NEWLINE {
			fmt.Printf("[%s %q]\n", token.Kind, token.Literal)
		} else {
			fmt.Printf("[%s %q]", token.Kind, token.Literal)
		}
	}
	fmt.Println()
}

// func main() {
// 	// r1 := strings.NewReader(`# File number one
// 	// HELLO=world # This delcares the HELLO variable
// 	// # Database config below
// 	// DATABASE=postgres
// 	// `)

// 	// r2 := strings.NewReader(`# File number two
// 	// FOO="bar" # This delcares the foo variable
// 	// PLEASE_DELETE=me
// 	// DATABASE=mysql
// 	// `)

// 	r1 := strings.NewReader(`
// 	FOO=bar
// 	FOO=baz
// 	FOO="buz"
// 	`)

// 	r2 := strings.NewReader(`
// 	FOO=qux
// 	`)

// 	// read the files

// 	f1, err := dotenv.Read(r1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	f2, err := dotenv.Read(r2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// // merge the files

// 	if err := f1.Merge(f2); err != nil {
// 		panic(err)
// 	}

// 	// // set extra values

// 	// f1.Set("MY_EXTRA_VAR", `"with a value"`, "")
// 	// f1.Delete("PLEASE_DELETE")

// 	fmt.Println(f1)
// }
