package client

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/lunaraio/lunar-go-sdk/src/types"
)

var json = jsoniter.ConfigFastest

type Client struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

func NewClient(config types.Config) *Client {
	baseURL := config.BaseURL
	if baseURL == "" {
		baseURL = "https://www.lunartools.co"
	}

	return &Client{
		apiKey:  config.APIKey,
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *Client) AddProduct(product types.AddProduct) error {
	if strings.TrimSpace(product.Token) == "" {
		return fmt.Errorf("user access token is required")
	}
	if strings.TrimSpace(product.Name) == "" {
		return fmt.Errorf("product name is required")
	}
	if strings.TrimSpace(product.SKU) == "" {
		return fmt.Errorf("product SKU is required")
	}
	if product.Qty < 0 {
		return fmt.Errorf("product quantity must be a non-negative number")
	}
	if product.Value != nil && *product.Value < 0 {
		return fmt.Errorf("product value must be a non-negative number")
	}
	if product.Spent != nil && *product.Spent < 0 {
		return fmt.Errorf("product spent must be a non-negative number")
	}

	return c.makeRequest("POST", "/api/sdk/add-product", product, nil)
}

func (c *Client) AddOrder(order types.AddOrder) error {
	if strings.TrimSpace(order.Token) == "" {
		return fmt.Errorf("user access token is required")
	}
	if strings.TrimSpace(order.Name) == "" {
		return fmt.Errorf("order name is required")
	}
	if strings.TrimSpace(order.Status) == "" {
		return fmt.Errorf("order status is required")
	}
	if strings.TrimSpace(order.OrderNumber) == "" {
		return fmt.Errorf("order number is required")
	}

	return c.makeRequest("POST", "/api/sdk/add-order", order, nil)
}

func (c *Client) AddProfile(profile types.AddProfile) error {
	if strings.TrimSpace(profile.Token) == "" {
		return fmt.Errorf("user access token is required")
	}

	if err := c.validateAddress(profile.Billing, "billing"); err != nil {
		return err
	}
	if err := c.validateAddress(profile.Shipping, "shipping"); err != nil {
		return err
	}
	if err := c.validatePayment(profile.Payment); err != nil {
		return err
	}

	return c.makeRequest("POST", "/api/sdk/add-profile", profile, nil)
}

func (c *Client) AddTask(task types.AddTask) error {
	if strings.TrimSpace(task.Token) == "" {
		return fmt.Errorf("user access token is required")
	}
	if strings.TrimSpace(task.Bot) == "" {
		return fmt.Errorf("bot name is required")
	}
	if strings.TrimSpace(task.Site) == "" {
		return fmt.Errorf("site is required")
	}
	if strings.TrimSpace(task.Mode) == "" {
		return fmt.Errorf("mode is required")
	}
	if strings.TrimSpace(task.Input) == "" {
		return fmt.Errorf("input is required")
	}

	if err := c.validateAddress(task.Profile.Billing, "billing"); err != nil {
		return err
	}
	if err := c.validateAddress(task.Profile.Shipping, "shipping"); err != nil {
		return err
	}
	if err := c.validatePayment(task.Profile.Payment); err != nil {
		return err
	}

	if strings.TrimSpace(task.Proxy) == "" {
		return fmt.Errorf("proxy is required")
	}

	return c.makeRequest("POST", "/api/sdk/add-task", task, nil)
}

func (c *Client) validateAddress(address types.Address, addressType string) error {
	if strings.TrimSpace(address.Name) == "" {
		return fmt.Errorf("%s name is required", addressType)
	}
	if strings.TrimSpace(address.Phone) == "" {
		return fmt.Errorf("%s phone is required", addressType)
	}
	if strings.TrimSpace(address.Line1) == "" {
		return fmt.Errorf("%s line1 is required", addressType)
	}
	if strings.TrimSpace(address.PostCode) == "" {
		return fmt.Errorf("%s postCode is required", addressType)
	}
	if strings.TrimSpace(address.City) == "" {
		return fmt.Errorf("%s city is required", addressType)
	}
	if strings.TrimSpace(address.Country) == "" {
		return fmt.Errorf("%s country is required", addressType)
	}
	if strings.TrimSpace(address.State) == "" {
		return fmt.Errorf("%s state is required", addressType)
	}
	return nil
}

func (c *Client) validatePayment(payment types.Payment) error {
	if strings.TrimSpace(payment.Name) == "" {
		return fmt.Errorf("payment name is required")
	}
	if strings.TrimSpace(payment.Type) == "" {
		return fmt.Errorf("payment type is required")
	}
	if strings.TrimSpace(payment.LastFour) == "" {
		return fmt.Errorf("payment lastFour is required")
	}
	if strings.TrimSpace(payment.ExpMonth) == "" {
		return fmt.Errorf("payment expMonth is required")
	}
	if strings.TrimSpace(payment.ExpYear) == "" {
		return fmt.Errorf("payment expYear is required")
	}
	return nil
}

func (c *Client) Webhook(webhookURL string, payload types.Webhook) (*types.WebhookResponse, error) {
	hasContent := payload.Content != nil && strings.TrimSpace(*payload.Content) != ""
	hasEmbeds := len(payload.Embeds) > 0

	if !hasContent && !hasEmbeds {
		return nil, fmt.Errorf("webhook payload must contain either content or at least one embed")
	}
	if len(payload.Embeds) > 10 {
		return nil, fmt.Errorf("discord webhooks support a maximum of 10 embeds")
	}

	for i, embed := range payload.Embeds {
		if len(embed.Fields) > 25 {
			return nil, fmt.Errorf("embed %d exceeds the maximum of 25 fields", i)
		}
		for j, field := range embed.Fields {
			if strings.TrimSpace(field.Name) == "" {
				return nil, fmt.Errorf("embed %d, field %d: name is required", i, j)
			}
			if strings.TrimSpace(field.Value) == "" {
				return nil, fmt.Errorf("embed %d, field %d: value is required", i, j)
			}
		}
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal webhook payload: %w", err)
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("webhook request failed with status: %d", resp.StatusCode)
	}

	var webhookResp types.WebhookResponse
	if err := json.NewDecoder(resp.Body).Decode(&webhookResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &webhookResp, nil
}

func (c *Client) makeRequest(method, path string, payload interface{}, result interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest(method, c.baseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
