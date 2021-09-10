package main

import (
	"fmt"
	"strings"
)


func mergeTemplateWithContact(required_template Template, required_contact ContactResponse) {
	from := envVariable("MAIL")
	password := envVariable("PASWD")
	for key, contact := range required_contact.Result {
		fmt.Println(key, "_", contact)
		//didn't find replace element in template so , supposed only to replace for contact_name

		msg := strings.ReplaceAll(required_template.Template_text, "{{contact_name}}", contact.Contact_name)
		fmt.Println(msg, "msg")

		sendMail(from, password, []string{contact.Contact_name}, msg)
	}
}
