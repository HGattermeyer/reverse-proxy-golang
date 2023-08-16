package models

type Server struct {
	ID             uint
	Uri            string
	RedirectCount  int
	RedirectServer []RedirectServer
}

// type Server struct {
// 	Uri             string   `json:"uri"`
// 	RedirectCount   int      `json:"redirectCount"`
// 	RedirectServers []string `json:"redirectServers"`
// }
