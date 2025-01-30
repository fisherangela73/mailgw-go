# mailgw-go

A lightweight and feature-complete Go wrapper for the [Mail.gw](https://mail.gw) API, providing easy access to temporary email services.

## Installation

```go
go get github.com/fisherangela73/mailgw-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/fisherangela73/mailgw-go/pkg/mailgw"
)

func main() {
    // Create a new client
    client := mailgw.NewClient()

    // Get available domains
    domains, _ := client.GetDomains()
    
    // Create a temporary email account
    account, _ := client.CreateAccount("test@"+domains[0].Domain, "password123")
    
    // Login
    client.Login(account.Address, "password123")
    
    // Check for new messages
    message, err := client.GetLastMessage()
    if err == nil {
        fmt.Printf("New message: %s\n", message.Subject)
    }
}
```

## Features

- 🚀 Complete Mail.gw API coverage
- 💡 Simple and intuitive interface
- 📦 Zero external dependencies
- ✨ Fully typed responses
- ⚙️ Configurable client

## API Reference

### Creating a Client

```go
// Default client
client := mailgw.NewClient()

// Custom configuration
client := mailgw.NewClient(
    mailgw.WithBaseURL("https://custom.mail.gw"),
    mailgw.WithToken("your-token"),
)
```

### Account Operations

```go
// Create account
account, err := client.CreateAccount("address@domain.com", "password")

// Login
err := client.Login("address@domain.com", "password")

// Get account details
me, err := client.GetMe()

// Delete account
err := client.DeleteAccount("account-id")
```

### Message Operations

```go
// Get all messages
messages, err := client.GetMessages()

// Get latest message
message, err := client.GetLastMessage()

// Get specific message
message, err := client.GetMessage("message-id")

// Mark as read
err := client.MarkMessageAsRead("message-id")

// Delete message
err := client.DeleteMessage("message-id")
```

### Domain Operations

```go
// Get available domains
domains, err := client.GetDomains()

// Get specific domain
domain, err := client.GetDomain("domain-id")
```

## Error Handling

The client returns appropriate errors for various scenarios:
- Network errors
- Authentication errors (401)
- Rate limiting (429)
- Invalid requests (400)
- Not found errors (404)

## Examples

Check the [examples](examples) directory for more detailed usage examples.

## Contributing

Contributions are welcome! Feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) for details.
