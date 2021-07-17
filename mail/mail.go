package mail

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"text/template"
)

type Dest struct {
	Name  string
	Msg   string
	Reclu string
}

func checker(err error) {
	if err != nil {
		log.Panic(err.Error())
	}
}

func Mail(Reclu, Msg string) {
	mailFrom := os.Getenv("mailfrom")
	mailTo := os.Getenv("mailto")
	passw := os.Getenv("psw")
	from := mail.Address{"sebastian trabajo", mailFrom}
	to := mail.Address{"sebastian", mailTo}
	subject := "trabajos"
	dest := Dest{Name: to.Address, Msg: Msg, Reclu: Reclu}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject
	headers["Content-Type"] = `text/html; charset="UTF-0"`
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s \r\n", k, v)

	}
	t, err := template.ParseFiles("temp.html")
	checker(err)
	buf := new(bytes.Buffer)
	err = t.Execute(buf, dest)
	checker(err)
	message += buf.String()
	host := "smtp.gmail.com"
	servername := host + ":465"
	auth := smtp.PlainAuth("", mailFrom, passw, host)
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	checker(err)
	client, err := smtp.NewClient(conn, host)
	checker(err)
	err = client.Auth(auth)
	checker(err)
	err = client.Mail(from.Address)
	checker(err)
	err = client.Rcpt(to.Address)
	w, err := client.Data()
	checker(err)
	_, err = w.Write([]byte(message))
	checker(err)
	err = w.Close()
	checker(err)
	client.Quit()
}
