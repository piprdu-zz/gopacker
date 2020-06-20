package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/piprdu/gopacker/greeting" // Importing a nested package
	"github.com/piprdu/gopacker/numbers"
	"github.com/piprdu/gopacker/strings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("callicoder"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
