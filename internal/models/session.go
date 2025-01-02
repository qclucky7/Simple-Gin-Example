package models

type Session struct {
	Account   string   `json:"account"`
	AccountId string   `json:"accountId"`
	Powers    []string `json:"powers"`
}
