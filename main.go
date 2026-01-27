package lunartools

import (
	"github.com/lunartools/lunar-go-sdk/src/client"
	"github.com/lunartools/lunar-go-sdk/src/types"
)

type Client = client.Client
type Config = types.Config
type Webhook = types.Webhook
type Embed = types.Embed
type Field = types.Field
type Footer = types.Footer
type Author = types.Author
type Thumbnail = types.Thumbnail
type Image = types.Image
type WebhookResponse = types.WebhookResponse
type AddProduct = types.AddProduct
type AddOrder = types.AddOrder
type AddProfile = types.AddProfile
type Address = types.Address
type Payment = types.Payment

var NewClient = client.NewClient

func String(s string) *string {
	return &s
}

func Int(i int) *int {
	return &i
}

func Float64(f float64) *float64 {
	return &f
}

func Bool(b bool) *bool {
	return &b
}
