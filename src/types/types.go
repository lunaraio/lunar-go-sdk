package types

type Config struct {
	APIKey  string
	BaseURL string
}

type Thumbnail struct {
	URL *string `json:"url,omitempty"`
}

type Image struct {
	URL *string `json:"url,omitempty"`
}

type Footer struct {
	Text    *string `json:"text,omitempty"`
	IconURL *string `json:"icon_url,omitempty"`
}

type Author struct {
	Name    *string `json:"name,omitempty"`
	URL     *string `json:"url,omitempty"`
	IconURL *string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline *bool  `json:"inline,omitempty"`
}

type Embed struct {
	Author      *Author    `json:"author,omitempty"`
	Title       *string    `json:"title,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Description *string    `json:"description,omitempty"`
	Color       *int       `json:"color,omitempty"`
	Fields      []Field    `json:"fields,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Timestamp   *string    `json:"timestamp,omitempty"`
}

type Webhook struct {
	Username  *string `json:"username,omitempty"`
	AvatarURL *string `json:"avatar_url,omitempty"`
	Content   *string `json:"content,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

type WebhookResponse struct {
	Status      string `json:"status"`
	QueueLength int    `json:"queueLength"`
}

type AddProduct struct {
	Token string   `json:"token"`
	Name  string   `json:"name"`
	SKU   string   `json:"sku"`
	Qty   int      `json:"qty"`
	Size  *string  `json:"size,omitempty"`
	Store *string  `json:"store,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Spent *float64 `json:"spent,omitempty"`
}

type AddOrder struct {
	Token       string  `json:"token"`
	Name        string  `json:"name"`
	Status      string  `json:"status"`
	OrderNumber string  `json:"orderNumber"`
	Image       *string `json:"image,omitempty"`
	Tracking    *string `json:"tracking,omitempty"`
	Date        *string `json:"date,omitempty"`
	Qty         *string `json:"qty,omitempty"`
	Price       *string `json:"price,omitempty"`
	OrderTotal  *string `json:"orderTotal,omitempty"`
	Account     *string `json:"account,omitempty"`
	Retailer    *string `json:"retailer,omitempty"`
	Tags        *string `json:"tags,omitempty"`
}

type Address struct {
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Line1    string  `json:"line1"`
	Line2    *string `json:"line2,omitempty"`
	PostCode string  `json:"postCode"`
	City     string  `json:"city"`
	Country  string  `json:"country"`
	State    string  `json:"state"`
}

type Payment struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	LastFour string  `json:"lastFour"`
	ExpMonth string  `json:"expMonth"`
	ExpYear  string  `json:"expYear"`
	CVV      *string `json:"cvv,omitempty"`
}

type AddProfile struct {
	Token    string  `json:"token"`
	Success  bool    `json:"success"`
	Billing  Address `json:"billing"`
	Shipping Address `json:"shipping"`
	Payment  Payment `json:"payment"`
}
