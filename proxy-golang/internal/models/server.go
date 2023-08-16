package models

type ServerFile struct {
	Uri             string
	RedirectCount   int
	RedirectServers []string
}

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
