package gosercomp

//go:generate skyencoder -struct SkyGroup -no-test
type SkyGroup struct {
	Id     int32
	Name   string
	Colors []string
}
