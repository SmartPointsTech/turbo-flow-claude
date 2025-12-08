package main

import (
	"context"
	"fmt"

	"golang.org/x/crypto/acme"
)

func main() {
	var c *acme.Client = &acme.Client{}
	// This method usage caused the build failure
	_, err := c.FetchRenewalInfo(context.Background(), "http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("FetchRenewalInfo exists!")
	}
}
