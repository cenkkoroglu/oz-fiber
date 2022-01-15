package models

type ApiResponse struct {
	Data     interface{} `json:"data"`
	Success  bool        `json:"success"`
	Errors   []string    `json:"errors"`
	Warnings []string    `json:"warnings"`
	Messages []string    `json:"messages"`
}
