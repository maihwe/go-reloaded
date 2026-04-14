// package main

// import (
// 	"regexp"
// )

// func fixSingleQuote(s string) string {
// 	res := regexp.MustCompile(`'\s*(.*?)\s*'`)
// 	return res.ReplaceAllString(s, "'$1'")
// }

package main

import (
	// "fmt"
	"strings"
)

func fixSingleQuote(s string) string {
	words := strings.Split(s, "'")
	for i := 1; i < len(words); i++ {
		words[i] = strings.TrimSpace(words[i])
	}
	return strings.Join(words, "'")
}

// func main()  {
// 	fmt.Printf("%q", hee("As Elton John said: ' I am the most well-known homosexual in the world '"))
// }
