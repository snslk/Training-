package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234s"
)

var db *sql.DB
var tmpl = template.Must(template.ParseFiles("template/login.html"))
var hometmpl = template.Must(template.ParseFiles("template/home.html"))
var inserttmpl = template.Must(template.ParseFiles("template/insert.html"))
var deletetmpl = template.Must(template.ParseFiles("template/delete.html"))
var edittmpl = template.Must(template.ParseFiles("template/edit.html"))
var showtmpl = template.Must(template.ParseFiles("template/show.html"))
var signuptmpl = template.Must(template.ParseFiles("template/signup.html"))

//**********************************************************************************

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
	host, port, user, password)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}


	fmt.Println("Successfully connected to the database!")

	r := mux.NewRouter()
	r.HandleFunc("/", loginHandler).Methods("GET", "POST")
	r.HandleFunc("/home", homeHandler).Methods("GET")
	r.HandleFunc("/insert",insertHandler).Methods("GET","POST")
	r.HandleFunc("/signup",signupHandler).Methods("GET","POST")
	r.HandleFunc("/delete",deleteHandler).Methods("GET","POST")
	r.HandleFunc("/edit",editHandler).Methods("GET","POST")
	r.HandleFunc("/show",showHandler).Methods("GET","POST")

	log.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
//***************************************************************

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("name")
		password := r.FormValue("user_id")

		if checkUser(username, password)==true {
			http.Redirect(w,r,"/home?username="+username,http.StatusSeeOther)
		} else {
			fmt.Fprintf(w, "User not found or incorrect password %s.",username)
		}
	} else {
		tmpl.Execute(w, nil)
	}
}

//******************************************************************

func checkUser(username, password string) bool {
	var dbUsername, dbPassword string
	query := "select user_id,name from users where user_id=$1 and name=$2 ;"
	err:= db.QueryRow(query,password,username).Scan(&dbPassword,&dbUsername)
	if err != nil {
		if err == sql.ErrNoRows {
			return false 
		}
		log.Fatalf("Error querying database: %v", err)
	}

	return dbPassword == password
}

//*******************************************************************

//******************************************************************
func signupHandler(w http.ResponseWriter ,r *http.Request){
	
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		userid := r.FormValue("user_id")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")
		age :=r.FormValue("age")
	
		if insertUser(1,name,userid,email,mobile,age )==true {
			fmt.Fprintf(w,"<a href=/>login</a>")
			return
		}else{
			fmt.Fprintf(w,"Error signin\n")
		}

	}else {
		signuptmpl.Execute(w, nil)
	}
}

//************************************************************************
func insertUser(num int,name,userid,email,mobile,age string) bool {
	var query string 
	if (num == 1){
		query = "insert into users values($1,$2,$3,$4,$5);"
	}else if (num == 0){
		query = "insert into user_list values($1,$2,$3,$4,$5);"
	}		
	_,err:= db.Query(query,userid,name,email,mobile,age)
	if err != nil {
		if err == sql.ErrNoRows {
			return false 
		}
		log.Fatalf("Error querying database: %v", err)
	}

	return true
}

//************************************************************************
func insertHandler(w http.ResponseWriter,r *http.Request){
	
	if r.Method==http.MethodPost {
		name := r.FormValue("name")
		userid := r.FormValue("user_id")
		address := r.FormValue("add")
		qualification := r.FormValue("qual")
		age := r.FormValue("age")
	
		if insertUser(0,name,userid,address,qualification,age) {
			fmt.Fprintf(w,"<a href=/home>Home Page</a>")
			return
		}else{
			fmt.Fprintf(w,"Error signin\n")
		}

	}else {
		inserttmpl.Execute(w,nil)
	}
}
//****************************************************************************
func deleteHandler(w http.ResponseWriter ,r *http.Request){
	
	if r.Method == http.MethodPost {
		userId := r.FormValue("user_id")
		query := "delete from user_list where user_id=$1;"
		_,err:= db.Query(query,userId)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w,"Error \n")
			}
			log.Fatalf("Error querying database: %v", err)
		}
		fmt.Fprintf(w,"<a href=/home>Home Page</a>")
		
		}else { 
		deletetmpl.Execute(w,nil)
	}	
}

func showHandler(w http.ResponseWriter ,r *http.Request){
	var id,name,add,qual,age string
	fmt.Fprintf(w,"-----------------------------------------------------------------------------------------\n")
    	fmt.Fprintf(w,"|Id\t|Name\t\t|Address\t\t\t|Qualification\t\t|Age\t|\n")
	fmt.Fprintf(w,"-----------------------------------------------------------------------------------------\n")
	rows, err := db.Query("SELECT * FROM user_list")
	if err != nil {
    	log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
    	if err := rows.Scan(&id, &name ,&add ,&qual, &age); err != nil {
        	log.Fatal(err)
    	}
    	fmt.Fprintf(w,"%s\t%s\t\t%s\t\t\t\t%s\t\t\t%s\t\n",id, name,add,qual,age)
	}
}


func editHandler(w http.ResponseWriter ,r *http.Request ){
	var query string	
	if r.Method == http.MethodPost {
		userId := r.FormValue("user_id")
		entry := r.FormValue("entry")
		ty := r.FormValue("name")
		
		query = fmt.Sprintf("update user_list set %s = '%s' where user_id = '%s' ",ty,entry,userId)				
		_,err:=db.Query(query)
		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Fprintf(w,"Error \n")
			}
			log.Fatalf("Error querying database: %v", err)
		}
		fmt.Fprintf(w,"<a href=/home>Home Page</a>")
		
		}else { 
			edittmpl.Execute(w,nil)
	}
}
func homeHandler(w http.ResponseWriter ,r *http.Request ){
		hometmpl.Execute(w,nil)
		var id,name,add,qual,age string
	fmt.Fprintf(w,"-----------------------------------------------------------------------------------------\n")
    	fmt.Fprintf(w,"|Id\t|Name\t\t|Address\t\t\t|Qualification\t\t|Age\t|\n")
	fmt.Fprintf(w,"-----------------------------------------------------------------------------------------\n")
	rows, err := db.Query("SELECT * FROM user_list")
	if err != nil {
    	log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
    	if err := rows.Scan(&id, &name ,&add ,&qual, &age); err != nil {
        	log.Fatal(err)
    	}
    	fmt.Fprintf(w,"%s\t%s\t\t%s\t\t\t\t%s\t\t\t%s\t\n",id, name,add,qual,age)
	}


}
