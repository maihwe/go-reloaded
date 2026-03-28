// package main

// import ("fmt"
// 	"strconv"	
// )

// func hexToDec(hexStr string) (int64, error) {
// 	return strconv.ParseInt(hexStr, 16, 64)
// }

// func main() {
// 	fmt.Println(hexToDec("1E"))
// 	fmt.Println(hexToDec("FF"))
// }

// package main

// import ("fmt"
// 		"strconv"
// )

// func binToDec(binStr string) (int64, error) {
// 	return strconv.ParseInt(binStr, 2, 64)
// }

// func main() {
// 	fmt.Println(binToDec("1010"))
// 	fmt.Println(binToDec("11111111"))
// 	fmt.Println(binToDec("10"))
// }

// package main

// import ("fmt"
// 	"strings"
// )

// func firstcap(text string) string {
// 	if len(text) == 0 {
// 		return text 
// }
// 	return strings.ToUpper(text[0:1]) + strings.ToLower(text[1:])
// }
//  func main() {
// 	fmt.Println(firstcap("mATHIAS"))
// 	fmt.Println(firstcap("hELLO wORLD"))
//  }

// package main

// import ("fmt"
// 	"strings"
// )

// func UpperN(text []string, n int) string {
// 	if n > len(text) {
// 		n = len(text)
// 	}

// 	for i := len(text) - n; i < len(text); i++ {
// 		text[i] = strings.ToUpper(text[i])
// 	}
// 	return strings.Join(text, " ")
// }

// func main() {
// 	fmt.Println(UpperN([]string{"Jesus", "is", "lord"}, 2))
// }

// package main

// import ("fmt"
// 	"strings"
// )

// func punc(s string) bool {
// 	return strings.ContainsAny(s, ",.!?;:")
// }

// func main() {
// 	fmt.Println(punc(","))
// 	fmt.Println(punc("H"))
// 	fmt.Println(punc("."))
// }

// 

// package main

// import ("fmt"
// 	"strings"
// )

// func article(word string) string {

// 	vowels := "AEIOUHaeiouh"
// 	if strings.ContainsAny(string(word[0]), vowels) {
// 		return "an " + word
// 	}
// 	return "a " + word
// }

// func main() {
// 	fmt.Println(article("apple"))
// 	fmt.Println(article("banana"))
// 	fmt.Println(article("book"))
// 	fmt.Println(article("animal"))
// 	fmt.Println(article("hour"))

// }

package main

import "fmt"

func reverse(word string) string {
    runes := []rune(word)
    for i, j := 0, len(word) - 1; i <j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
func main() {
    fmt.Println(reverse("madam"))
    fmt.Println(reverse("level"))
}