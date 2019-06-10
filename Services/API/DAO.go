package apiService

import "net"

type dao interface {
	OpenLock(string) bool
	AddBadge(string, string, string) bool
	DeleteBadge(string) bool
	GetServerAddress() net.IP
	ChangeMode(string) bool
	GetCurrentMode() string
	GetLastLog() string
	GetBadgesList() []Badge
	SetNomPrenom(string, string, string) bool
}
