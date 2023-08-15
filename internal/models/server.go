package models

type Server struct {
	Uri             string   `json:"uri"`
	RedirectCount   int      `json:"redirectCount"`
	RedirectServers []string `json:"redirectServers"`
}
