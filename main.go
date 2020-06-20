package main

import (
	"fmt"
	str "strings" // Package Alias

	"github.com/piprdu/gopacker/numbers"
	"github.com/piprdu/gopacker/strings"
	"github.com/piprdu/gopacker/strings/greeting" // Importing a nested package
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("gopacker"))

	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
