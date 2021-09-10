
App written in Go Lang to perform some critsend Api call and solve the following problems

You can find the documentation from https://beta.critsend.com/docs

The token to test is `as Provided`

You need to :

Retrieve the template 17 from the api

Retrieve the list 11 from contact_list

Retrieve the data of each contact in list 11

Replace the variable into template with the data retrieved :

var:17 is the content of data field 17
Send the email to an smtp server :`as provided`


use following key in .env file
BASE_URL=https://api.critsend.io/
AUTHENTICATION_TOKEN={{token}}
MAIL={{from_mail}}
PASWD={{mail_password}}

To work well with dotenv package export GO111MODULE=on from cli or include in .bashrc related file
