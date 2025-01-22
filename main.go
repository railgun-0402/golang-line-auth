package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

// LINE認証に必要な値を管理
// TODO: 環境変数化
const (
	clientID = ""
	redirectURI = ""
	leinAuthURL = ""
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// 認証URLの構築
	authURL, _ := url.Parse(leinAuthURL)
	query := authURL.Query()
	query.Set("response_type", "code")
	query.Set("client_id", clientID)
	query.Set("redirect_uri", redirectURI)
	query.Set("state", "random_state_string")
	query.Set("scope", "profile openid email")
	authURL.RawQuery = query.Encode()

	// redirect
	http.Redirect(w, r, authURL.String(), http.StatusFound)
}

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Start Line Auth\n")
	}

	http.HandleFunc("/", helloHandler)

	log.Println("Server Start at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}