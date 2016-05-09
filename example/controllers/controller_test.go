package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

//for tests need first run http example server

func TestMainIndex(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:7000/")
	if err != nil {
		t.Error(err)
	}
	io.Copy(os.Stdout, resp.Body)

	//for send some arguments need specify full path - I KNOW, IT`S UGLY
	resp, err = http.Get("http://127.0.0.1:7000/Main/Index/arg0/arg1")
	if err != nil {
		t.Error(err)
	}
	io.Copy(os.Stdout, resp.Body)
}

func TestUsersGetJson(t *testing.T) {
	body := `{"name":"UserName","ip":"127.0.0.1"}`

	resp, err := http.Post("http://127.0.0.1:7000/Users/GetMixJson", "application/json", strings.NewReader(body))
	if err != nil {
		t.Error(err)
	}
	io.Copy(os.Stdout, resp.Body)
}

func TestCityGetJson(t *testing.T) {
	body := `{"name":"Moscow","country":"Russia"}`

	resp, err := http.Post("http://127.0.0.1:7000/City/GetFieldJson", "application/json", strings.NewReader(body))
	if err != nil {
		t.Error(err)
	}
	io.Copy(os.Stdout, resp.Body)
}

func BenchmarkMainIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://127.0.0.1:7000/")
		if err != nil {
			fmt.Println(err)
			// b.Error(err)
			continue
		}
		resp.Body.Close()
	}
}

func BenchmarkCity(b *testing.B) {
	body := strings.NewReader(`{"name":"Moscow","country":"Russia"}`)

	for i := 0; i < b.N; i++ {
		resp, err := http.Post("http://127.0.0.1:7000/City/GetFieldJson", "application/json", body)
		body.Seek(0, 0)

		if err != nil {
			fmt.Println(err)
			// b.Error(err)
			continue
		}
		resp.Body.Close()
	}
}
