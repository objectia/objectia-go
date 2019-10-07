package objectia

import (
	"path/filepath"
	"time"
)

// Message model
type Message struct {
	Domain      string    `json:"domain" xml:"domain"`
	Date        time.Time `json:"date" xml:"date"`
	From        string    `json:"from" xml:"from"`
	FromName    string    `json:"from_name" xml:"from_name"`
	ReplyTo     string    `json:"reply_to" xml:"reply_to"`
	To          []string  `json:"to" xml:"to"`
	Cc          []string  `json:"cc" xml:"cc"`
	Bcc         []string  `json:"bcc" xml:"bcc"`
	Subject     string    `json:"subject" xml:"subject"`
	Text        string    `json:"text" xml:"text"`
	HTML        string    `json:"html" xml:"html"`
	Attachments []string  `json:"attachments" xml:"attachments"`
	Charset     string    `json:"charset,omitempty" xml:"charset,omitempty"`
	Encoding    string    `json:"encoding,omitempty" xml:"encoding,omitempty"`
}

// NewMessage ...
func NewMessage(from, subject, text string, to ...string) *Message {
	return &Message{
		From:    from, //FIXME: from, fromName
		Subject: subject,
		Text:    text,
		To:      to,
	}
}

// AddCc ...
func (m *Message) AddCc(cc ...string) {
	m.Cc = append(m.Cc, cc...)
}

// AddBcc ...
func (m *Message) AddBcc(bcc ...string) {
	m.Bcc = append(m.Bcc, bcc...)
}

// AddAttachment ...
func (m *Message) AddAttachment(fileName string) {
	m.Attachments = append(m.Attachments, fileName)
}

// SetHTML ...
func (m *Message) SetHTML(html string) {
	m.HTML = html
}

// SetReplyTo ...
func (m *Message) SetReplyTo(recipient string) {
	m.ReplyTo = recipient
}

// ToParameters ...
func (m *Message) ToParameters() *Parameters {
	params := NewParameters()
	if !m.Date.IsZero() {
		params.Add("data", m.Date)
	}
	params.Add("from", m.From)
	if len(m.FromName) > 0 {
		params.Add("from_name", m.FromName)
	}
	params.Add("to", m.To)

	if len(m.Cc) > 0 {
		params.Add("cc", m.Cc)
	}
	if len(m.Bcc) > 0 {
		params.Add("bcc", m.Bcc)
	}
	params.Add("subject", m.Subject)

	params.Add("text", m.Text)

	if len(m.HTML) > 0 {
		params.Add("html", m.HTML)
	}

	//FIXME: Add rest of fields...

	for _, fn := range m.Attachments {
		key := filepath.Base(fn)
		params.AddFile(key, fn)
	}

	return params
}
