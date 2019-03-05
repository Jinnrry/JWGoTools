package tools

import (
	"net/smtp"
	"strings"
)

func SendToMail(senderuser, senderpassword, senderhost, to, title, body, mailtype string) error {
	hp := strings.Split(senderhost, ":")
	auth := smtp.PlainAuth("", senderuser, senderpassword, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + senderuser + ">\r\nSubject: " + title + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(senderhost, auth, senderuser, send_to, msg)
	return err
}
