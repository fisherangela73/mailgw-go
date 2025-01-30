package mailgw

import "time"

type Domain struct {
	ID        string    `json:"id"`
	Domain    string    `json:"domain"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Account struct {
	ID         string    `json:"id"`
	Address    string    `json:"address"`
	Quota      int       `json:"quota"`
	Used       int       `json:"used"`
	IsDisabled bool      `json:"isDisabled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type Message struct {
	ID             string         `json:"id"`
	AccountID      string         `json:"accountId"`
	MsgID          string         `json:"msgid"`
	From           EmailAddress   `json:"from"`
	To             []EmailAddress `json:"to"`
	Subject        string         `json:"subject"`
	Intro          string         `json:"intro,omitempty"`
	Text           string         `json:"text,omitempty"`
	HTML           []string       `json:"html,omitempty"`
	HasAttachments bool           `json:"hasAttachments"`
	Size           int            `json:"size"`
	DownloadURL    string         `json:"downloadUrl"`
	Seen           bool           `json:"seen"`
	IsDeleted      bool           `json:"isDeleted"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

type EmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type TokenResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
