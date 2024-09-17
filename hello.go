package hello

import (
	"fmt"
	"net/http"
)

func PrintHello() {
	fmt.Println("Hello from dependency!")

	// This is not something we expect from this dependency.
	_, _ = http.Get("https://example.com/not-good")
}
