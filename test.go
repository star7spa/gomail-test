package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	mail "github.com/xhit/go-simple-mail/v2"
)

const htmlBody = `<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
		<title>Hello!</title>
	</head>
	<body>
		<p>This is a <b>test</b>.</p>
	</body>
</html>`

var ()

func main() {

	var hostname string
	flag.StringVar(&hostname, "hostname", "localhost", "Hostname of the mysql server to connect to")

	var port int
	flag.IntVar(&port, "port", 2325, "Hostname of the mysql server to connect to")

	var username string
	flag.StringVar(&username, "username", os.Getenv("RELAY_USER"), "username of the mysql server to connect to")

	var password string
	flag.StringVar(&password, "password", os.Getenv("RELAY_PASS"), "password of the mysql server to connect to")

	flag.Parse()

	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = hostname
	server.Port = port
	server.Username = username
	server.Password = password
	server.Encryption = mail.EncryptionNone

	// Can be PLAIN, LOGIN, CRAM-MD5, None
	server.Authentication = mail.AuthLogin

	// Connect with SMTP
	smtpClient, err := server.Connect()

	// Check connection
	if err != nil {
		fmt.Println("Fatal")
		log.Fatal(err)
	}

	// New email
	email := mail.NewMSG()
	email.SetFrom("noreply <noreply@whatyouwant>").
		AddTo("testemail@test.test").
		SetSubject("New Email test")
	email.SetBody(mail.TextHTML, htmlBody)

	// Check before send email
	if email.Error != nil {
		log.Fatal(email.Error)
	}

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Println("Error during send process")
		log.Println(err)
	} else {
		log.Println("Email Sent")
	}
}
