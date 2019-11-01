package objectia

// Send sends a mail message
func (c *Mail) Send(message *Message) (*SendMailReceipt, error) {
	var resp Response

	params := message.ToParameters()

	err := c.client.post("/v1/mail/send", params, &resp)
	if err != nil {
		return nil, err
	}

	result := &SendMailReceipt{}
	err = fromMap(resp.Data, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
