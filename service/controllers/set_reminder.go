package controllers

import (
	"net/http"
    "database/sql"
    "fmt"
	"log"

    _ "github.com/mattn/go-sqlite3"
)

type reminder struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Time string `json:"time"`
}

// ReminderController controls reminder operations
func ReminderController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	r.ParseForm()

	request := reminder{
		Title: r.FormValue("title"),
		Description: r.FormValue("description"),
		Time: r.FormValue("time"),
	}
	fmt.Println(request)

	addReminder(request, w)

}

func addReminder(reminderObject reminder,  w http.ResponseWriter) {

	db, err := sql.Open("sqlite3", "./jarvis.db")
	checkErr(err)
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS reminder (id INTEGER PRIMARY KEY AUTOINCREMENT, title TEXT, description TEXT, time TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	checkErr(err)

	stmt, err := tx.Prepare("insert into reminder(title, description, time) values(?, ?, ?)")
	checkErr(err)
	defer stmt.Close()

	_, err = stmt.Exec(reminderObject.Title, reminderObject.Description, reminderObject.Time)
	checkErr(err)
	tx.Commit()
	
	w.Write([]byte(`{"status": "success", "message": "Reminder has been set !"}`))
}

func ShowReminder() []reminder{
	var result []reminder
	
	db, err := sql.Open("sqlite3", "./jarvis.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("select id, title, description, time from reminder")
	checkErr(err)

	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var description string
		var time string
		err = rows.Scan(&id, &title, &description, &time)
		checkErr(err)

		r := reminder {
			Id: id,
			Title: title,
			Description: description,
			Time: time,
		}
		result = append(result, r) 
	}
	err = rows.Err()
	checkErr(err)

	return result
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}