package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// LINE認証に必要な値を管理
const (
	clientID = ""
	redirectURI = ""
	lineAuthURL = ""
	tokenURL = ""
	clientSecret = ""
)

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// 認証コード
	code := r.URL.Query().Get("code")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", redirectURI)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)

	// Send Post Request
	response, err := http.PostForm(tokenURL, data)
	if err != nil {
		http.Error(w, "Failed to get token", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Fprintf(w, "Reponse: %s\n", body)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("start redirect...")

	// 認証URLの構築
	authURL, _ := url.Parse(lineAuthURL)
	query := authURL.Query()
	query.Set("response_type", "code")
	query.Set("client_id", clientID)
	query.Set("redirect_uri", redirectURI)
	query.Set("state", "random_state_string")
	query.Set("scope", "profile openid email")
	authURL.RawQuery = query.Encode()

	// redirect
	http.Redirect(w, r, authURL.String(), http.StatusFound)

	log.Println("end redirect...")
}

func main() {
	http.HandleFunc("/login", loginHandler)

	http.HandleFunc("/callback", callbackHandler)

	log.Println("Server Start at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}