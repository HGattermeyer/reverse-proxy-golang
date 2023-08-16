package models

type Server struct {
	ID             uint
	Uri            string
	RedirectCount  int
	Strategy       string
	RedirectServer []RedirectServer
}
