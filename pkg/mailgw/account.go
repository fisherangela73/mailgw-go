package mailgw

import "fmt"

type CreateAccountRequest struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

func (c *Client) CreateAccount(address, password string) (*Account, error) {
	req := CreateAccountRequest{
		Address:  address,
		Password: password,
	}
	var account Account
	err := c.doRequest("POST", "/accounts", req, &account)
	return &account, err
}

func (c *Client) Login(address, password string) error {
	req := CreateAccountRequest{
		Address:  address,
		Password: password,
	}
	var tokenResp TokenResponse
	if err := c.doRequest("POST", "/token", req, &tokenResp); err != nil {
		return err
	}
	c.Token = tokenResp.Token
	return nil
}

func (c *Client) GetMe() (*Account, error) {
	var account Account
	err := c.doRequest("GET", "/me", nil, &account)
	return &account, err
}

func (c *Client) DeleteAccount(id string) error {
	return c.doRequest("DELETE", fmt.Sprintf("/accounts/%s", id), nil, nil)
}
