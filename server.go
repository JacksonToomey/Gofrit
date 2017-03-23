package main

import (
	"fmt"

	"github.com/jacksontoomey/gofrit/sessions"
)

func main() {
	manager := sessions.GetSessionManager()
	fmt.Printf("%v\n", manager)
}
