package controllers

import (
	"net/http"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/smtp"
	"strings"
	_"github.com/mattn/go-sqlite3"
	"time"
)

// Mail ....
type Mail struct {
	ID      int       `json:"id"`
	Sender  string    `json: "sender"`
	To      []string  `json: "to"`
	Cc      []string  `json: "cc"`	
	Bcc     []string  `json: "bcc"`
	Subject string	  `json: "subject"`
	Body    string    `json: "body"`
	Time    string	  `json: "time"`
}

// SMTPServer ...
type SMTPServer struct {
	Host      string
	Port      string
	TLSConfig *tls.Config
}

// ServerName ...
func (s *SMTPServer) ServerName() string {
	return s.Host + ":" + s.Port
}

// EmailController controls reminder operations
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
		Sender: r.FormValue("sender"),
		To: toArr,
		Cc: ccArr,
		Bcc: bccArr,
		Subject: r.FormValue("subject"),
		Body: r.FormValue("body"),
		Time: r.FormValue("time"),
	}

	fmt.Println("request: ", request)

	if request.Time == "" {
		fmt.Println("Sending Mail")
		send(request, w)
	} else {
		fmt.Println("Sending Mail1111")
		addMail(request, w)
	}

}

// BuildMessage ...
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
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		log.Panic(err)
	}

	// step 1: Use Auth
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	// step 2: add all from and to
	if err = client.Mail(mail.Sender); err != nil {
		log.Panic(err)
	}
	receivers := append(mail.To, mail.Cc...)
	receivers = append(receivers, mail.Bcc...)
	for _, k := range receivers {
		log.Println("sending to: ", k)
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	// Data
	wr, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = wr.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = wr.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	log.Println("Mail sent successfully")
	w.Write([]byte(`{"status":"success", "message": "Mail sent Successfully"}`))

}

func addMail(mailObject Mail,  w http.ResponseWriter) {

	db, err := sql.Open("sqlite3", "./mail.db")
	checkErr(err)
	defer db.Close()

	// sqlStmt := `CREATE TABLE IF NOT EXISTS mail (id INTEGER PRIMARY KEY AUTOINCREMENT, sender TEXT, to TEXT, cc TEXT, bcc TEXT, subject TEXT, body TEXT, time TEXT);`
	sqlStmt := `CREATE TABLE IF NOT EXISTS mail (id INTEGER PRIMARY KEY AUTOINCREMENT, sender TEXT, time TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("insert into mail(sender, time) values(?, ?)")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(mailObject.Sender, mailObject.Time )
	checkErr(err)
	tx.Commit()
	
	w.Write([]byte(`{"status": "success", "message": "Mail has been set !"}`))
}

// CheckMail to check the pending mails in the database
func CheckMail() []Mail{
	var result []Mail
	
	db, err := sql.Open("sqlite3", "./mail.db")
	checkErr(err)
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS mail (id INTEGER PRIMARY KEY AUTOINCREMENT, sender TEXT, time TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return result
	}

	rows, err := db.Query("select id, sender, time from mail")
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var sender string
		var time string
		err = rows.Scan(&id, &sender, &time)
		checkErr(err)

		r := Mail {
			ID: id,
			Sender: sender,
			Time: time,
		}
		result = append(result, r) 
	}
	err = rows.Err()
	checkErr(err)

	fmt.Println("result:", result )
	return result
}

// CheckTime to check the time of the pending mails
func CheckTime(n time.Duration, data []Mail) bool {
	fmt.Println("data: ", data)
	fmt.Println("data.Time", data[0].Time)

	timestamp := time.Now().Local()
	// count := 0
	fmt.Println(strings.timestamp)

	for range time.Tick(n *time.Second) {
		// str := "Polling remote terminal data at <some remote terminal name> at "+ timestamp.String()
		//  fmt.Println(str, "count: ", count)
		//  count ++
		// for 
	} 

	return false
}