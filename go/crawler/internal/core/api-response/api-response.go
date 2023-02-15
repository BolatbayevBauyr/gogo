package api_response

import "time"

type APIResponse struct {
	Total    int    `json:"total"`
	Next     string `json:"next"`
	DataList []Data `json:"data_list"`
}

type Data struct {
	Uid             string      `json:"uid"`
	PageTitle       string      `json:"page_title"`
	PageDescription string      `json:"page_description"`
	Alias           string      `json:"alias"`
	IsBold          bool        `json:"is_bold"`
	PublicType      []int       `json:"public_type"`
	PublishedDate   time.Time   `json:"published_date"`
	NewsImage       interface{} `json:"news_image"`
}
