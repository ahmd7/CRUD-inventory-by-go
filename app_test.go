package main

import (
	"log"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(BDuser, DBpasswd, DBhost, "test")
	if err != nil {
		log.Fatal("Error occured while initializing the database")
	}
	m.Run()
}
