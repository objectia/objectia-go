package objectia

import (
	"path/filepath"
	"time"
)

// MailMessage model
type MailMessage struct {
	Date        time.Time `json:"date"`
	From        string    `json:"from"`
	FromName    string    `json:"from_name"`
	ReplyTo     string    `json:"reply_to"`
	To          []string  `json:"to"`
	Cc          []string  `json:"cc"`
	Bcc         []string  `json:"bcc"`
	Subject     string    `json:"subject"`
	Text        string    `json:"text"`
	HTML        string    `json:"html"`
	Attachments []string  `json:"attachments"`
	Tags        []string  `json:"tags"`
	Charset     string    `json:"charset,omitempty"`
	Encoding    string    `json:"encoding,omitempty"`

	// Options:
	RequireTLS             *bool `json:"require_tls,omitempty"`
	VerifyCertificate      *bool `json:"verify_cert,omitempty"`
	OpenTracking           *bool `json:"open_tracking,omitempty"`
	ClickTracking          *bool `json:"click_tracking,omitempty"`
	PlainTextClickTracking *bool `json:"text_click_tracking"`
	UnsubscribeTracking    *bool `json:"unsubscribe_tracking,omitempty"`
	TestMode               *bool `json:"test_mode,omitempty"`
}

// NewMessage ...
func NewMessage(from, subject, text string, to ...string) *MailMessage {
	return &MailMessage{
		From:    from,
		Subject: subject,
		Text:    text,
		To:      to,
	}
}

// AddCc ...
func (m *MailMessage) AddCc(cc ...string) {
	m.Cc = append(m.Cc, cc...)
}

// AddBcc ...
func (m *MailMessage) AddBcc(bcc ...string) {
	m.Bcc = append(m.Bcc, bcc...)
}

// AddAttachment ...
func (m *MailMessage) AddAttachment(fileName string) {
	if len(m.Attachments) < 10 {
		m.Attachments = append(m.Attachments, fileName)
	}
}

// AddTag ...
func (m *MailMessage) AddTag(tag string) {
	if len(m.Tags) < 3 {
		m.Tags = append(m.Tags, tag)
	}
}

// SetHTML ...
func (m *MailMessage) SetHTML(html string) {
	m.HTML = html
}

// SetReplyTo ...
func (m *MailMessage) SetReplyTo(recipient string) {
	m.ReplyTo = recipient
}

// SetTestMode ...
func (m *MailMessage) SetTestMode(flag bool) {
	m.TestMode = new(bool)
	*m.TestMode = flag
}

// SetRequireTLS ...
func (m *MailMessage) SetRequireTLS(flag bool) {
	m.RequireTLS = new(bool)
	*m.RequireTLS = flag
}

// SetVerifyCertificate ...
func (m *MailMessage) SetVerifyCertificate(flag bool) {
	m.VerifyCertificate = new(bool)
	*m.VerifyCertificate = flag
}

// SetOpenTracking ...
func (m *MailMessage) SetOpenTracking(flag bool) {
	m.OpenTracking = new(bool)
	*m.OpenTracking = flag
}

// SetClickTracking ...
func (m *MailMessage) SetClickTracking(flag bool) {
	m.ClickTracking = new(bool)
	*m.ClickTracking = flag
}

// SetPlainTextClickTracking ...
func (m *MailMessage) SetPlainTextClickTracking(flag bool) {
	m.PlainTextClickTracking = new(bool)
	*m.PlainTextClickTracking = flag
}

// SetUnsubscribeTracking ...
func (m *MailMessage) SetUnsubscribeTracking(flag bool) {
	m.UnsubscribeTracking = new(bool)
	*m.UnsubscribeTracking = flag
}

// ToParameters ...
func (m *MailMessage) ToParameters() *Parameters {
	params := NewParameters()
	if !m.Date.IsZero() {
		params.Add("date", m.Date)
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

	if len(m.ReplyTo) > 0 {
		params.Add("reply_to", m.ReplyTo)
	}

	if len(m.Tags) > 0 {
		params.Add("tags", m.Tags)
	}

	for _, fn := range m.Attachments {
		key := filepath.Base(fn)
		params.AddFile(key, fn)
	}

	params.Add("charset", m.Charset)
	params.Add("encoding", m.Encoding)

	if m.RequireTLS != nil {
		params.Add("require_tls", *m.RequireTLS)
	}
	if m.VerifyCertificate != nil {
		params.Add("verify_cert", *m.VerifyCertificate)
	}
	if m.OpenTracking != nil {
		params.Add("open_tracking", *m.OpenTracking)
	}
	if m.ClickTracking != nil {
		params.Add("click_tracking", *m.ClickTracking)
	}
	if m.PlainTextClickTracking != nil {
		params.Add("text_click_tracking", *m.PlainTextClickTracking)
	}
	if m.UnsubscribeTracking != nil {
		params.Add("unsubscribe_tracking", *m.UnsubscribeTracking)
	}
	if m.TestMode != nil {
		params.Add("test_mode", *m.TestMode)
	}

	return params
}
