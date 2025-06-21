package main

import (
	"html/template"
	"net/http"
	"time"

	"webpractise/session"
)

var globalSessions *session.Manager

// 然後在 init 函式中初始化
func init() {
	globalSessions, _ = session.NewManager("memory", "gosessionid", 3600)
	go globalSessions.GC()
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}

// func main() {
// 	db, err := sql.Open("mysql", "root:123123@tcp(127.0.0.1:3306)/will?charset=utf8")
// 	checkErr(err)

// 	stmt, err := db.Prepare("INSERT userinfo SET username=?,department=?,created=?")
// 	checkErr(err)

// 	res, err := stmt.Exec("cute cat", "SWEETY", "2020-06-20")
// 	checkErr(err)

// 	id, err := res.LastInsertId()
// 	checkErr(err)

// 	row, _ := db.Query("SELECT * FROM userinfo WHERE uid=?", id)
// 	fmt.Println("row:")
// 	for row.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string
// 		err = row.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println(uid, username, department, created)
// 	}
// 	rows, _ := db.Query("SELECT * FROM userinfo")
// 	fmt.Println("rows:")
// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created string
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println(uid, username, department, created)
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
