package mailgw

import "fmt"

type DomainResponse struct {
	Domains []Domain `json:"hydra:member"`
}

func (c *Client) GetDomains() ([]Domain, error) {
	var response DomainResponse
	err := c.doRequest("GET", "/domains", nil, &response)
	return response.Domains, err
}

func (c *Client) GetDomain(id string) (*Domain, error) {
	var domain Domain
	err := c.doRequest("GET", fmt.Sprintf("/domains/%s", id), nil, &domain)
	return &domain, err
}
