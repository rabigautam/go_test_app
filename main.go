package main

import (
	// "fmt"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
)

/**
*to load the env variable
 */

func envVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error occurs while loading .env")
	}
	return os.Getenv((key))
}

func main() {
	baseUrl := envVariable("BASE_URL")
	url := baseUrl + "list/"
	
	token := envVariable(("AUTHENTICATION_TOKEN"))

	//request to given url
	req, err := http.NewRequest("GET", url, nil)

	//add headers to request

	req.Header.Add("Authorization",token)
	req.Header.Add("Content_type","application/json")

	//https request with bad certificate
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	//send request using httpclient

	client := &http.Client{Transport:tr}

	resp,err := client.Do(req)

	if err != nil {
		log.Fatalf(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err!=nil{
		log.Println("Error in response",err.Error())
	}
	log.Println(string([]byte(body)))
}
