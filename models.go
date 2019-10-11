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
	Index *int `json:"-" xml:"index,attr,omitempty"`
	//---
	IPAddress string `json:"ip,omitempty" xml:"ip,omitempty"`
	Type      string `json:"type,omitempty" xml:"type,omitempty"`
	Hostname  string `json:"hostname,omitempty" xml:"hostname,omitempty"`

	Continent     string `json:"continent_name,omitempty" xml:"continent_name,omitempty"`
	ContinentCode string `json:"continent_code,omitempty" xml:"continent_code,omitempty"`

	Country       string `json:"country_name,omitempty" xml:"country_name,omitempty"`
	CountryNative string `json:"country_name_native,omitempty" xml:"country_name_native,omitempty"`
	CountryCode   string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	CountryCode3  string `json:"country_code3,omitempty" xml:"country_code3,omitempty"`
	Capital       string `json:"capital,omitempty" xml:"capital,omitempty"`

	Region     string `json:"region_name,omitempty" xml:"region_name,omitempty"`
	RegionCode string `json:"region_code,omitempty" xml:"region_code,omitempty"` //FIXME!
	City       string `json:"city,omitempty" xml:"city,omitempty"`
	Postcode   string `json:"postcode,omitempty" xml:"postcode,omitempty"`

	Latitude  float64 `json:"latitude,omitempty" xml:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty" xml:"longitude,omitempty"`

	PhonePrefix string       `json:"phone_prefix,omitempty" xml:"phone_prefix,omitempty"`
	Currencies  []IPCurrency `json:"currencies,omitempty" xml:"currencies,omitempty"`
	Languages   []IPLanguage `json:"languages,omitempty" xml:"languages,omitempty"`
	Flag        string       `json:"flag,omitempty" xml:"flag,omitempty"`
	FlagEmoji   string       `json:"flag_emoji,omitempty" xml:"flag_emoji,omitempty"`
	IsEU        *bool        `json:"is_eu,omitempty" xml:"is_eu,omitempty"`
	TLD         string       `json:"internet_tld,omitempty" xml:"internet_tld,omitempty"`
	ISP         string       `json:"isp,omitempty" xml:"isp,omitempty"`
	Timezone    *IPTimezone  `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Security    *IPSecurity  `json:"security,omitempty" xml:"security,omitempty"`
}

// IPCurrency model
type IPCurrency struct {
	Code          string `json:"code,omitempty" xml:"code,omitempty"`
	NumericCode   string `json:"num_code,omitempty" xml:"num_code,omitempty"`
	Name          string `json:"name,omitempty" xml:"name,omitempty"`
	PluralName    string `json:"name_plural,omitempty" xml:"plural_name,omitempty"`
	Symbol        string `json:"symbol,omitempty" xml:"symbol,omitempty"`
	NativeSymbol  string `json:"symbol_native,omitempty" xml:"native_symbol,omitempty"`
	DecimalDigits int    `json:"decimal_digits,omitempty" xml:"decimal_digits,omitempty"`
}

// IPLanguage model
type IPLanguage struct {
	Code       string `json:"code,omitempty" xml:"code,omitempty"`
	Code2      string `json:"code2,omitempty" xml:"code2,omitempty"`
	Name       string `json:"name,omitempty" xml:"name,omitempty"`
	NativeName string `json:"native_name,omitempty" xml:"native_name,omitempty"`
	RTL        bool   `json:"rtl" xml:"rtl"`
}

// IPTimezone model
type IPTimezone struct {
	ID             string `json:"id,omitempty" xml:"id,omitempty"` // Name not ID?
	LocalTime      string `json:"localtime,omitempty" xml:"localtime,omitempty"`
	GMTOffset      int    `json:"gmt_offset,omitempty" xml:"gmt_offset,omitempty"`
	Code           string `json:"code,omitempty" xml:"code,omitempty"`
	DaylightSaving bool   `json:"daylight_saving" xml:"daylight_saving"`
}

// IPSecurity model
type IPSecurity struct {
	IsProxy     bool     `json:"is_proxy" xml:"is_proxy"`
	ProxyType   string   `json:"proxy_type,omitempty" xml:"proxy_type,omitempty"`
	IsCrawler   bool     `json:"is_crawler" xml:"is_crawler"`
	CrawlerName string   `json:"crawler_name,omitempty" xml:"crawler_name,omitempty"`
	CrawlerType string   `json:"crawler_type,omitempty" xml:"crawler_type,omitempty"`
	IsTOR       bool     `json:"is_tor" xml:"is_tor"`
	ThreatLevel string   `json:"threat_level,omitempty" xml:"threat_level,omitempty"`
	ThreatTypes []string `json:"threat_types,omitempty" xml:"threat_types,omitempty"`
}

// APIUsage model
type APIUsage struct {
	GeoLocationRequests int `json:"geoip_requests,omitempty" xml:"geoip_requests,omitempty"`
	MailRequests        int `json:"mail_requests,omitempty" xml:"mail_requests,omitempty"`
}

// SendMailReceipt model
type SendMailReceipt struct {
	ID                  string `json:"id" xml:"id"`
	AcceptedReceipients int    `json:"accepted_recipients" xml:"accepted_recipients"`
	RejectedReceipients int    `json:"rejected_recipients" xml:"rejected_recipients"`
}
