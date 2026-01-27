# Lunar Tools Go SDK

Official Go SDK for the Lunar Tools API.

## Installation

```bash
go get github.com/lunaraio/lunar-go-sdk
```

## Getting Started

### 1. Obtain Your API Key
Visit the [Lunar Tools Developer Portal](https://www.lunartools.co/developers) to generate your API key.

### 2. Get User Access Tokens
Each API call requires a user's access token. Users can find their access token in the Lunar Tools application under Settings > Developer.

## Usage

### Initialize the Client

```go
package main

import (
    lunartools "github.com/lunaraio/lunar-go-sdk"
)

func main() {
    client := lunartools.NewClient(lunartools.Config{
        APIKey: "your-api-key-here",
    })
}
```

### Add Product to Inventory

```go
err := client.AddProduct(lunartools.AddProduct{
    Token: "user-access-token",
    Name:  "Charizard VMAX",
    SKU:   "SWSH-074",
    Qty:   5,
    Value: lunartools.Float64(150.00),
    Spent: lunartools.Float64(120.00),
    Size:  lunartools.String("Standard"),
    Store: lunartools.String("TCGPlayer"),
})
if err != nil {
    log.Fatal(err)
}
```

### Add Order

```go
err := client.AddOrder(lunartools.AddOrder{
    Token:       "user-access-token",
    Name:        "Pokemon Booster Box",
    Status:      "shipped",
    OrderNumber: "ORD-12345",
    Price:       lunartools.String("120.00"),
    OrderTotal:  lunartools.String("132.00"),
    Retailer:    lunartools.String("Amazon"),
    Tracking:    lunartools.String("1Z999AA10123456784"),
    Qty:         lunartools.String("1"),
})
if err != nil {
    log.Fatal(err)
}
```

### Add Profile Analytics

```go
err := client.AddProfile(lunartools.AddProfile{
    Token:   "user-access-token",
    Success: true,
    Billing: lunartools.Address{
        Name:     "John Doe",
        Phone:    "5551234567",
        Line1:    "123 Main St",
        Line2:    lunartools.String("Apt 4B"),
        PostCode: "10001",
        City:     "New York",
        State:    "NY",
        Country:  "United States",
    },
    Shipping: lunartools.Address{
        Name:     "John Doe",
        Phone:    "5551234567",
        Line1:    "123 Main St",
        Line2:    lunartools.String("Apt 4B"),
        PostCode: "10001",
        City:     "New York",
        State:    "NY",
        Country:  "United States",
    },
    Payment: lunartools.Payment{
        Name:     "John Doe",
        Type:     "Visa",
        LastFour: "4242",
        ExpMonth: "12",
        ExpYear:  "2025",
        CVV:      lunartools.String("123"),
    },
})
if err != nil {
    log.Fatal(err)
}
```

### Forward Webhook to Discord

```go
response, err := client.Webhook(
    "https://www.lunartools.co/api/webhooks/YOUR_TOKEN_HERE",
    lunartools.Webhook{
        Content: lunartools.String("New product in stock!"),
        Embeds: []lunartools.Embed{
            {
                Title:       lunartools.String("Product Alert"),
                Description: lunartools.String("Charizard VMAX is now available"),
                Color:       lunartools.Int(0x5865F2),
                Fields: []lunartools.Field{
                    {
                        Name:   "Price",
                        Value:  "$150.00",
                        Inline: lunartools.Bool(true),
                    },
                    {
                        Name:   "Quantity",
                        Value:  "5",
                        Inline: lunartools.Bool(true),
                    },
                },
                Timestamp: lunartools.String(time.Now().Format(time.RFC3339)),
            },
        },
    },
)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Status: %s, Queue Length: %d\n", response.Status, response.QueueLength)
```

## API Reference

### Types

#### Config

```go
type Config struct {
    APIKey  string
    BaseURL string // Optional, defaults to https://www.lunartools.co
}
```

#### AddProduct

```go
type AddProduct struct {
    Token string   // Required: User's access token
    Name  string   // Required: Product name
    SKU   string   // Required: Product SKU
    Qty   int      // Required: Quantity
    Size  *string  // Optional: Product size
    Store *string  // Optional: Store name
    Value *float64 // Optional: Product value
    Spent *float64 // Optional: Amount spent
}
```

#### AddOrder

```go
type AddOrder struct {
    Token       string  // Required: User's access token
    Name        string  // Required: Order name
    Status      string  // Required: Order status
    OrderNumber string  // Required: Order number
    Image       *string // Optional: Product image URL
    Tracking    *string // Optional: Tracking number
    Date        *string // Optional: Order date (format: MM/DD/YYYY, HH:MM:SS AM/PM)
    Qty         *string // Optional: Quantity
    Price       *string // Optional: Item price
    OrderTotal  *string // Optional: Total order amount
    Account     *string // Optional: Account name
    Retailer    *string // Optional: Retailer name
    Tags        *string // Optional: Order tags
}
```

#### AddProfile

```go
type AddProfile struct {
    Token    string  // Required: User's access token
    Success  bool    // Required: Whether the checkout was successful
    Billing  Address // Required: Billing address information
    Shipping Address // Required: Shipping address information
    Payment  Payment // Required: Payment information
}

type Address struct {
    Name     string  // Required: Full name
    Phone    string  // Required: Phone number
    Line1    string  // Required: Address line 1
    Line2    *string // Optional: Address line 2
    PostCode string  // Required: Postal/ZIP code
    City     string  // Required: City
    Country  string  // Required: Country
    State    string  // Required: State/province
}

type Payment struct {
    Name     string  // Required: Name on card
    Type     string  // Required: Card type (e.g., Visa, Mastercard)
    LastFour string  // Required: Last 4 digits of card
    ExpMonth string  // Required: Expiration month (MM)
    ExpYear  string  // Required: Expiration year (YYYY)
    CVV      *string // Optional: CVV code
}
```

#### Webhook

```go
type Webhook struct {
    Username  *string // Optional: Override webhook username
    AvatarURL *string // Optional: Override webhook avatar
    Content   *string // Optional: Message content
    Embeds    []Embed // Optional: Array of embeds (max 10)
}

type Embed struct {
    Author      *Author    // Optional: Embed author
    Title       *string    // Optional: Embed title
    URL         *string    // Optional: Embed URL
    Description *string    // Optional: Embed description
    Color       *int       // Optional: Embed color (hex)
    Fields      []Field    // Optional: Array of fields (max 25)
    Thumbnail   *Thumbnail // Optional: Thumbnail image
    Image       *Image     // Optional: Main image
    Footer      *Footer    // Optional: Footer
    Timestamp   *string    // Optional: Timestamp (ISO 8601)
}
```

### Methods

#### AddProduct

Add a new product to a user's inventory.

```go
func (c *Client) AddProduct(product AddProduct) error
```

**Example:**

```go
err := client.AddProduct(lunartools.AddProduct{
    Token: "user-access-token",
    Name:  "Product Name",
    SKU:   "SKU-123",
    Qty:   10,
    Value: lunartools.Float64(50.00),
})
if err != nil {
    log.Fatal(err)
}
```

#### AddOrder

Add a new order to a user's orders.

```go
func (c *Client) AddOrder(order AddOrder) error
```

**Example:**

```go
err := client.AddOrder(lunartools.AddOrder{
    Token:       "user-access-token",
    Name:        "Pokemon Cards",
    Status:      "delivered",
    OrderNumber: "ORD-456",
    Price:       lunartools.String("99.99"),
    Retailer:    lunartools.String("eBay"),
})
if err != nil {
    log.Fatal(err)
}
```

#### AddProfile

Add profile analytics data for tracking successful/declined checkouts.

```go
func (c *Client) AddProfile(profile AddProfile) error
```

**Example:**

```go
err := client.AddProfile(lunartools.AddProfile{
    Token:   "user-access-token",
    Success: true,
    Billing: lunartools.Address{
        Name:     "John Doe",
        Phone:    "5551234567",
        Line1:    "123 Main St",
        PostCode: "10001",
        City:     "New York",
        State:    "NY",
        Country:  "United States",
    },
    Shipping: lunartools.Address{
        Name:     "John Doe",
        Phone:    "5551234567",
        Line1:    "456 Oak Ave",
        PostCode: "10002",
        City:     "Brooklyn",
        State:    "NY",
        Country:  "United States",
    },
    Payment: lunartools.Payment{
        Name:     "John Doe",
        Type:     "Visa",
        LastFour: "4242",
        ExpMonth: "12",
        ExpYear:  "2025",
    },
})
if err != nil {
    log.Fatal(err)
}
```

#### Webhook

Forward a webhook payload to Discord via Lunar Tools.

```go
func (c *Client) Webhook(webhookURL string, payload Webhook) (*WebhookResponse, error)
```

**Example:**

```go
response, err := client.Webhook(
    "https://www.lunartools.co/api/webhooks/TOKEN",
    lunartools.Webhook{
        Content: lunartools.String("Hello!"),
        Embeds: []lunartools.Embed{
            {
                Title:       lunartools.String("Alert"),
                Description: lunartools.String("Something happened"),
                Color:       lunartools.Int(0xFF0000),
                Fields: []lunartools.Field{
                    {
                        Name:  "Field 1",
                        Value: "Value 1",
                        Inline: lunartools.Bool(true),
                    },
                },
            },
        },
    },
)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Status: %s\n", response.Status)
```

## Helper Functions

The SDK provides helper functions for creating pointers to primitive types:

```go
lunartools.String("value")   // *string
lunartools.Int(123)          // *int
lunartools.Float64(99.99)    // *float64
lunartools.Bool(true)        // *bool
```

## Authentication

The SDK uses two-level authentication:

1. **API Key**: Your developer API key (passed in constructor)
   - Obtain from: [Developer Portal](https://www.lunartools.co/developers)
   - Used in `x-api-key` header for all requests

2. **User Access Token**: Individual user's access token (passed per request)
   - Users find this in: Lunar Tools App > Settings > Developer
   - Identifies which user's data to modify

## Error Handling

All methods return errors that should be handled appropriately:

```go
err := client.AddProduct(lunartools.AddProduct{
    Token: "user-access-token",
    Name:  "",
    SKU:   "SKU-123",
    Qty:   5,
})
if err != nil {
    log.Printf("Error: %v", err) // "product name is required"
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "time"

    lunartools "github.com/lunaraio/lunar-go-sdk"
)

func main() {
    // Initialize client
    client := lunartools.NewClient(lunartools.Config{
        APIKey: "your-api-key-here",
    })

    userToken := "user-access-token"

    // Add product
    err := client.AddProduct(lunartools.AddProduct{
        Token: userToken,
        Name:  "Limited Edition Sneakers",
        SKU:   "SNKR-001",
        Qty:   10,
        Value: lunartools.Float64(200.00),
        Spent: lunartools.Float64(150.00),
        Store: lunartools.String("Footlocker"),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("✅ Product added")

    // Add order
    err = client.AddOrder(lunartools.AddOrder{
        Token:       userToken,
        Name:        "Nike Air Max",
        Status:      "confirmed",
        OrderNumber: "ORD-12345",
        Price:       lunartools.String("150.00"),
        OrderTotal:  lunartools.String("165.00"),
        Retailer:    lunartools.String("Nike"),
        Tracking:    lunartools.String("1Z999AA10123456784"),
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("✅ Order added")

    // Track profile analytics
    err = client.AddProfile(lunartools.AddProfile{
        Token:   userToken,
        Success: true,
        Billing: lunartools.Address{
            Name:     "John Doe",
            Phone:    "5551234567",
            Line1:    "123 Main St",
            PostCode: "10001",
            City:     "New York",
            State:    "NY",
            Country:  "United States",
        },
        Shipping: lunartools.Address{
            Name:     "John Doe",
            Phone:    "5551234567",
            Line1:    "456 Oak Ave",
            PostCode: "10002",
            City:     "Brooklyn",
            State:    "NY",
            Country:  "United States",
        },
        Payment: lunartools.Payment{
            Name:     "John Doe",
            Type:     "Visa",
            LastFour: "4242",
            ExpMonth: "12",
            ExpYear:  "2025",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("✅ Profile analytics added")

    // Send Discord notification
    webhookURL := "https://www.lunartools.co/api/webhooks/YOUR_TOKEN"
    response, err := client.Webhook(webhookURL, lunartools.Webhook{
        Embeds: []lunartools.Embed{
            {
                Title:       lunartools.String("✅ Checkout Success"),
                Description: lunartools.String("Successfully checked out Nike Air Max"),
                Color:       lunartools.Int(0x00FF00),
                Fields: []lunartools.Field{
                    {
                        Name:   "Order Number",
                        Value:  "ORD-12345",
                        Inline: lunartools.Bool(true),
                    },
                    {
                        Name:   "Total",
                        Value:  "$165.00",
                        Inline: lunartools.Bool(true),
                    },
                },
                Timestamp: lunartools.String(time.Now().Format(time.RFC3339)),
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("✅ Webhook sent: %s (Queue: %d)\n", response.Status, response.QueueLength)

    fmt.Println("\n🎉 All operations completed successfully!")
}
```

## Performance

This SDK uses [jsoniter](https://github.com/json-iterator/go) for faster JSON encoding/decoding compared to the standard library.

## Support

For issues, questions, or feature requests:
- Discord: [Join our server](https://discord.gg/lunartools)
- Email: support@lunartools.co
- Documentation: [docs.lunartools.co](https://docs.lunartools.co)

## License

MIT