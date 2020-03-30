package objectia

import (
	"bytes"

	"github.com/objectia/objectia-go/internal"
)

// PDFCreateParams model
type PDFCreateParams struct {
	DocumentURL       string   `json:"document_url"`
	DocumentHTML      string   `json:"document_html"`
	DocumentName      string   `json:"document_name"`
	PageFormat        string   `json:"page_format"`
	PageWidth         string   `json:"page_width"`
	PageHeight        string   `json:"page_height"`
	MarginTop         string   `json:"margin_top"`
	MarginBottom      string   `json:"margin_bottom"`
	MarginLeft        string   `json:"margin_left"`
	MarginRight       string   `json:"margin_right"`
	PrintBackground   bool     `json:"print_background"`
	HeaderText        string   `json:"header_text"`
	HeaderAlign       string   `json:"header_align"`
	HeaderMargin      int      `json:"header_margin"`
	HeaderURL         string   `json:"header_url"`
	HeaderHTML        string   `json:"header_html"`
	FooterText        string   `json:"footer_text"`
	FooterAlign       string   `json:"footer_align"`
	FooterMargin      int      `json:"footer_margin"`
	FooterURL         string   `json:"footer_url"`
	FooterHTML        string   `json:"footer_html"`
	Orientation       string   `json:"orientation"`
	PageRanges        string   `json:"page_ranges"`
	PreferCSSPageSize bool     `json:"prefer_css_page_size"`
	Scale             float64  `json:"scale"`
	Author            string   `json:"author"`
	Title             string   `json:"title"`
	Creator           string   `json:"creator"`
	Subject           string   `json:"subject"`
	Keywords          []string `json:"keywords"`
	Language          string   `json:"language"`
	WatermarkURL      string   `json:"watermark_url"`
	WatermarkPosition string   `json:"watermark_position"`
	WatermarkOffsetX  int      `json:"watermark_offset_x"`
	WatermarkOffsetY  int      `json:"watermark_offset_y"`
	Encryption        string   `json:"encryption"`
	OwnerPassword     string   `json:"owner_password"`
	UserPassword      string   `json:"user_password"`
	Permissions       string   `json:"permissions"`
}

// ToParameters ...
func (p *PDFCreateParams) ToParameters() *internal.Parameters {
	params := internal.NewParameters()
	if len(p.DocumentURL) > 0 {
		params.Add("document_url", p.DocumentURL)
	} else if len(p.DocumentHTML) > 0 {
		params.Add("document_html", p.DocumentHTML)
	}

	params.Add("document_name", p.DocumentName)
	params.Add("page_format", p.PageFormat)
	params.Add("page_width", p.PageWidth)
	params.Add("page_height", p.PageHeight)
	params.Add("margin_top", p.MarginTop)
	params.Add("margin_bottom", p.MarginBottom)
	params.Add("margin_left", p.MarginLeft)
	params.Add("margin_right", p.MarginRight)
	params.Add("print_background", p.PrintBackground)

	if len(p.HeaderURL) > 0 {
		params.Add("header_url", p.HeaderURL)

	} else if len(p.HeaderHTML) > 0 {
		params.Add("header_html", p.HeaderHTML)

	} else if len(p.HeaderText) > 0 {
		params.Add("header_text", p.HeaderText)
		params.Add("header_align", p.HeaderAlign)
		params.Add("header_margin", p.HeaderMargin)
	}

	if len(p.FooterURL) > 0 {
		params.Add("footer_url", p.FooterURL)

	} else if len(p.FooterHTML) > 0 {
		params.Add("footer_html", p.FooterHTML)

	} else if len(p.FooterText) > 0 {
		params.Add("footer_text", p.FooterText)
		params.Add("footer_align", p.FooterAlign)
		params.Add("footer_margin", p.FooterMargin)
	}

	params.Add("orientation", p.Orientation)
	params.Add("page_ranges", p.PageRanges)
	params.Add("prefer_css_page_size", p.PreferCSSPageSize)
	params.Add("scale", p.Scale)
	params.Add("author", p.Author)
	params.Add("title", p.Title)
	params.Add("creator", p.Creator)
	params.Add("subject", p.Subject)
	params.Add("keywords", p.Keywords)
	params.Add("language", p.Language)

	if len(p.WatermarkURL) > 0 {
		params.Add("watermark_url", p.WatermarkURL)
		params.Add("watermark_position", p.WatermarkPosition)
		params.Add("watermark_offset_x", p.WatermarkOffsetX)
		params.Add("watermark_offset_y", p.WatermarkOffsetY)
	}

	if len(p.Encryption) > 0 {
		params.Add("encryption", p.Encryption)
		params.Add("owner_password", p.OwnerPassword)
		params.Add("user_password", p.UserPassword)
		params.Add("permissions", p.Permissions)
	}

	return params
}

// Create a PDF document from HTML.
func (c *PDF) Create(params *PDFCreateParams) ([]byte, error) {
	p := params.ToParameters()

	var resp bytes.Buffer
	err := c.client.post("/v1/pdf/create", p, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Bytes(), nil
}
