package app

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	app := &App{Message: "hello"}
	s := httptest.NewServer(app) // http provides its own test objects
	defer s.Close()
	resp, err := http.Get(s.URL) // test from the outside
	if err != nil {
		t.Error(err)
	}
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		t.Error(err)
	} else if string(body) != "hello" {
		t.Error("expected", "hello", "got", body)
	}
}

func fakeApp(msg string) *httptest.Server {
	app := &App{Message: msg}
	return httptest.NewServer(app)
}
func get(t *testing.T, s *httptest.Server, path string) string {
	resp, err := http.Get(s.URL + path)
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return string(body)
}

func TestHelpedHello(t *testing.T) {
	s := fakeApp("hello")
	defer s.Close()
	if body := get(t, s, "/"); body != "hello" {
		t.Error("expected", "hello", "got", body)
	}
}
