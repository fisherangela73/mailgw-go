package mailgw

import "fmt"

type MessageResponse struct {
	Messages []Message `json:"hydra:member"`
}

func (c *Client) GetMessages() ([]Message, error) {
	var response MessageResponse
	err := c.doRequest("GET", "/messages", nil, &response)
	return response.Messages, err
}

func (c *Client) GetMessage(id string) (*Message, error) {
	var message Message
	err := c.doRequest("GET", fmt.Sprintf("/messages/%s", id), nil, &message)
	return &message, err
}

func (c *Client) GetLastMessage() (*Message, error) {
	messages, err := c.GetMessages()
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("no messages found")
	}

	return c.GetMessage(messages[0].ID)
}

func (c *Client) DeleteMessage(id string) error {
	return c.doRequest("DELETE", fmt.Sprintf("/messages/%s", id), nil, nil)
}

func (c *Client) MarkMessageAsRead(id string) error {
	return c.doRequest("PATCH", fmt.Sprintf("/messages/%s", id), nil, nil)
}
