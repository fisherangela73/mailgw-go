package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fisherangela73/mailgw-go/pkg/mailgw"
)

func main() {
	// Create a new client
	client := mailgw.NewClient()

	// Get available domains
	domains, err := client.GetDomains()
	if err != nil {
		log.Fatal(err)
	}

	if len(domains) == 0 {
		log.Fatal("No domains available")
	}

	// Create a new email address
	address := fmt.Sprintf("test%d@%s", time.Now().Unix(), domains[0].Domain)
	password := "your-secure-password"

	// Create account
	account, err := client.CreateAccount(address, password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created account: %s\n", account.Address)

	// Login to get token
	err = client.Login(address, password)
	if err != nil {
		log.Fatal(err)
	}

	// Get account details
	me, err := client.GetMe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Logged in as: %s\n", me.Address)

	// Get messages
	messages, err := client.GetMessages()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("You have %d messages\n", len(messages))

	// Get latest message
	lastMessage, err := client.GetLastMessage()
	if err != nil {
		log.Printf("No messages yet: %v\n", err)
	} else {
		fmt.Printf("Latest message: %s\n", lastMessage.Subject)
	}
}
