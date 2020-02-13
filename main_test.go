package main

import (
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"testing"
)

var jar *cookiejar.Jar
var client *http.Client

func init() {
	var err error
	jar, err = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client = &http.Client{
		Jar: jar,
	}
}

func TestLogin(t *testing.T) {
	u, err := url.Parse("http://localhost:1337/api/auth/login")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Get(u.String())
	if err != nil {
		t.Fatal(err)
	}

	got := resp.StatusCode
	want := http.StatusOK
	if got != want {
		t.Fatalf("resp.StatusCode = %d; want %d", got, want)
	}

	if jar.Cookies(u)[0].Name != "sid" {
		t.Fatalf("Cookie.Name = %s; want %s", jar.Cookies(u)[0].Name, "sid")
	}
}

func TestSecret(t *testing.T) {
	u, err := url.Parse("http://localhost:1337/secret")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Get(u.String())
	if err != nil {
		t.Fatal(err)
	}

	got := resp.StatusCode
	want := http.StatusOK
	if got != want {
		t.Fatalf("resp.StatusCode = %d; want %d", got, want)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	gotBody := string(body)
	wantBody := "My little Secret"
	if gotBody != wantBody {
		t.Fatalf("string(body) = %s; want %s", gotBody, wantBody)
	}
}
