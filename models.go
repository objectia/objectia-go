package objectia

import (
	"time"
)

// ETag model
type ETag struct {
	Tag          string
	LastModified time.Time
}

// GeoLocationOptions model
type GeoLocationOptions struct {
	DisplayFields    string
	LookupHostname   bool
	ShowSecurityInfo bool
}

// IPLocation model
type IPLocation struct {
	IPAddress     string       `json:"ip,omitempty"`
	Type          string       `json:"type,omitempty"`
	Hostname      string       `json:"hostname,omitempty"`
	Continent     string       `json:"continent_name,omitempty"`
	ContinentCode string       `json:"continent_code,omitempty"`
	Country       string       `json:"country_name,omitempty"`
	CountryNative string       `json:"country_name_native,omitempty"`
	CountryCode   string       `json:"country_code,omitempty"`
	CountryCode3  string       `json:"country_code3,omitempty"`
	Capital       string       `json:"capital,omitempty"`
	Region        string       `json:"region_name,omitempty"`
	RegionCode    string       `json:"region_code,omitempty"`
	City          string       `json:"city,omitempty"`
	Postcode      string       `json:"postcode,omitempty"`
	Latitude      float64      `json:"latitude,omitempty"`
	Longitude     float64      `json:"longitude,omitempty"`
	PhonePrefix   string       `json:"phone_prefix,omitempty"`
	Currencies    []IPCurrency `json:"currencies,omitempty"`
	Languages     []IPLanguage `json:"languages,omitempty"`
	Flag          string       `json:"flag,omitempty"`
	FlagEmoji     string       `json:"flag_emoji,omitempty"`
	IsEU          *bool        `json:"is_eu,omitempty"`
	TLD           string       `json:"internet_tld,omitempty"`
	ISP           string       `json:"isp,omitempty"`
	Timezone      *IPTimezone  `json:"timezone,omitempty"`
	Security      *IPSecurity  `json:"security,omitempty"`
}

// IPCurrency model
type IPCurrency struct {
	Code          string `json:"code,omitempty"`
	NumericCode   string `json:"num_code,omitempty"`
	Name          string `json:"name,omitempty"`
	PluralName    string `json:"name_plural,omitempty"`
	Symbol        string `json:"symbol,omitempty"`
	NativeSymbol  string `json:"symbol_native,omitempty"`
	DecimalDigits int    `json:"decimal_digits,omitempty"`
}

// IPLanguage model
type IPLanguage struct {
	Code       string `json:"code,omitempty"`
	Code2      string `json:"code2,omitempty"`
	Name       string `json:"name,omitempty"`
	NativeName string `json:"native_name,omitempty"`
	RTL        bool   `json:"rtl"`
}

// IPTimezone model
type IPTimezone struct {
	ID             string `json:"id,omitempty"`
	LocalTime      string `json:"localtime,omitempty"`
	GMTOffset      int    `json:"gmt_offset,omitempty"`
	Code           string `json:"code,omitempty"`
	DaylightSaving bool   `json:"daylight_saving"`
}

// IPSecurity model
type IPSecurity struct {
	IsProxy     bool     `json:"is_proxy"`
	ProxyType   string   `json:"proxy_type,omitempty"`
	IsCrawler   bool     `json:"is_crawler"`
	CrawlerName string   `json:"crawler_name,omitempty"`
	CrawlerType string   `json:"crawler_type,omitempty"`
	IsTOR       bool     `json:"is_tor"`
	ThreatLevel string   `json:"threat_level,omitempty"`
	ThreatTypes []string `json:"threat_types,omitempty"`
}

// APIUsage model
type APIUsage struct {
	Requests map[string]int     `json:"requests,omitempty"`
	Cost     map[string]float64 `json:"cost,omitempty"`
}

// MailReceipt model
type MailReceipt struct {
	ID                 string `json:"id"`
	AcceptedRecipients int    `json:"accepted_recipients"`
	RejectedRecipients int    `json:"rejected_recipients"`
}

// SMSReceipt model
type SMSReceipt struct {
	ID    string  `json:"id"`
	From  string  `json:"from"`
	To    string  `json:"to"`
	Text  string  `json:"text"`
	Price float64 `json:"price"`
}

// CurrencyRateOptions model
type CurrencyRateOptions struct {
	Base       string
	Currencies []string
}

// CurrencyRate model
type CurrencyRate struct {
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Timestamp int64              `json:"timestamp"`
	Rates     map[string]float64 `json:"rates"`
}

// CurrencyConvert model
type CurrencyConvert struct {
	Amount    float64 `json:"amount"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Date      string  `json:"date,omitempty"`
	Timestamp int64   `json:"timestamp"`
	Rate      float64 `json:"rate"`
	Result    float64 `json:"result"`
}

// CurrencyRateChange model
type CurrencyRateChange struct {
	StartRate     float64 `json:"start,omitempty"`
	EndRate       float64 `json:"end,omitempty"`
	Change        float64 `json:"change,omitempty"`
	ChangePercent float64 `json:"change_percent,omitempty"`
	Error         string  `json:"error,omitempty"`
}

// CurrencyChange model
type CurrencyChange struct {
	Base      string                        `json:"base"`
	StartDate string                        `json:"start_date"`
	EndDate   string                        `json:"end_date"`
	Rates     map[string]CurrencyRateChange `json:"rates"`
}

// CurrencyTimeframe model
type CurrencyTimeframe struct {
	Base      string             `json:"base"`
	Currency  string             `json:"currency"`
	StartDate string             `json:"start_date"`
	EndDate   string             `json:"end_date"`
	Timespan  int                `json:"timespan"`
	Rates     map[string]float64 `json:"rates"`
}

// CurrencyList model
type CurrencyList struct {
	Currencies map[string]string `json:"currencies"`
}
