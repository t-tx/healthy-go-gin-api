package test

import (
	"net/http"
	"os"
	"testing"
)

var (
	username = "trungtvq"
	password = "AA1234aa"
	host     = "http://localhost:8080"
	client   http.Client
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
