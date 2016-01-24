package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	betfairLogin    = "type-your-login-here"
	betfairPassword = "type-your-password-here"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {

	// you have to have these files pre-generated
	cert, err := tls.LoadX509KeyPair("client-2048.crt", "client-2048.key")
	if err != nil {
		fmt.Println("Key load error. Details: ", err)
		return 1
	}

	config := tls.Config{Certificates: []tls.Certificate{cert}}

	tr := &http.Transport{TLSClientConfig: &config}

	client := http.Client{Transport: tr, Timeout: time.Second}

	req, err := http.NewRequest(
		"POST",
		"https://identitysso-api.betfair.com/api/certlogin",
		bytes.NewReader([]byte(fmt.Sprintf("username=%s&password=%s", betfairLogin, betfairPassword))))

	if err != nil {
		fmt.Println("HTTP POST request creation error. Details: ", err)
		return 2
	}

	req.Header.Add("X-Application", "appkey")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Server response error. Details: ", err)
		return 3
	}

	defer resp.Body.Close()

	fmt.Print("Response:")
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Response body reading error. Details: ", err)
		return 4
	}

	fmt.Printf("\nResponse status code: %d\n", resp.StatusCode)
	return 0
}
