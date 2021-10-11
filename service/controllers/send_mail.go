package controllers

import (
	"crypto/tls"
	"fmt"
	"github.com/Harkishen-Singh/Jarvis-personal-assistant/service/logger"
	"net/http"
	"net/smtp"
	"strings"
)

// Mail sets the mailing options.
type Mail struct {
	Sender  string
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

// SMTPServer sets SMTP server options.
type SMTPServer struct {
	Host      string
	Port      string
	TLSConfig *tls.Config
}

// ServerName returns the name of the server.
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

// EmailController controls reminder operations.
func EmailController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	to := r.FormValue("to")
	toArr := strings.Split(to, ";")
	cc := r.FormValue("cc")
	ccArr := strings.Split(cc, ";")
	bcc := r.FormValue("bcc")
	bccArr := strings.Split(bcc, ";")

	request := Mail{
		Sender:  r.FormValue("sender"),
		To:      toArr,
		Cc:      ccArr,
		Bcc:     bccArr,
		Subject: r.FormValue("subject"),
		Body:    r.FormValue("body"),
	}

	send(request, w)

}

// BuildMessage returns the resultant message built.
func (mail *Mail) BuildMessage() string {
	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.Sender)
	if len(mail.To) > 0 {
		header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}

	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += "\r\n" + mail.Body

	return header
}

func send(mailObject Mail, w http.ResponseWriter) {

	mail := Mail{}
	mail.Sender = mailObject.Sender
	mail.To = mailObject.To
	mail.Subject = mailObject.Subject
	mail.Body = mailObject.Body

	messageBody := mail.BuildMessage()

	smtpServer := SMTPServer{Host: "smtp.gmail.com", Port: "465"}
	smtpServer.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	auth := smtp.PlainAuth("", mail.Sender, "Jarvis@123", smtpServer.Host)

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), smtpServer.TLSConfig)
	if err != nil {
		logger.Error(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		logger.Error(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		logger.Error(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		logger.Error(err)
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		logger.Info("sending to: " + k)
		if err = client.Rcpt(k); err != nil {
			logger.Error(err)
		}
	}

	// Data
	wr, err := client.Data()
	if err != nil {
		logger.Error(err)
	}

	_, err = wr.Write([]byte(messageBody))
	if err != nil {
		logger.Error(err)
	}

	err = wr.Close()
	if err != nil {
		logger.Error(err)
	}

	client.Quit()

	logger.Info("Mail sent successfully")
	w.Write([]byte(`{"status":"success", "message": "Mail sent Successfully"}`))

}
