package main

import (
	"fmt"
	env "github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

/**
*to load the env variable
 */

func envVariable(key string) string {
	err := env.Load(".env")
	if err != nil {
		log.Fatalf("Error occurs while loading .env")
	}
	return os.Getenv((key))
}

type Template struct {
	Template_id         string
	Template_name       string
	Template_create     string
	Template_update     string
	Template_html       string
	Template_text       string
	Is_removable        string
	Template_screenshot string
}
type Contact struct {
	Contact_id     string
	Contact_name   string
	Contact_create string
	Contact_update string
	Is_removable   string
}
type TemplateResponse struct {
	Result []Template
}
type ContactResponse struct {
	Result []Contact
}

func main() {

	baseUrl := envVariable("BASE_URL")
	token := envVariable(("AUTHENTICATION_TOKEN"))

	from := envVariable("MAIL")
	password := envVariable("PASWD")

	// toList:=[]string{}

	fmt.Println("from :" + from)

	required_template, err := getTemplate(baseUrl, token)

	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(required_template, "required template")

	required_contact, err := getContact(baseUrl, token)
	resultContact := required_contact.Result

	for key, contact := range resultContact {
		fmt.Println(key, "_", contact)
		msg := strings.ReplaceAll(required_template.Template_text, "{{contact_name}}", contact.Contact_name)
		fmt.Println(msg, "msg")

		sendMail(from, password, []string{contact.Contact_name}, msg)
	}

}
