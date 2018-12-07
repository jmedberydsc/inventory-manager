package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Weapon struct {
	Id           int
	Name         string
	WeaponType   string
	Description  string
	Puncture     string
	Slice        string
	Impact       string
	Foi          string
	Balance      string
	Suitedness   string
	Construction string
	Weight       string
	Cost         string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "inventoryuser"
	dbPass := "123098pa"
	dbName := "inventory_manager"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM weapon ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	wea := Weapon{}
	res := []Weapon{}
	for selDB.Next() {
		var id int
		var name, wtype, description, puncture, slice, impact, foi, balance, suitedness, construction, weight, cost string
		err = selDB.Scan(&id, &name, &wtype, &description, &puncture, &slice, &impact, &foi, &balance, &suitedness, &construction, &weight, &cost)
		if err != nil {
			panic(err.Error())
		}
		wea.Id = id
		wea.Name = name
		wea.WeaponType = wtype
		wea.Description = description
		wea.Puncture = puncture
		wea.Slice = slice
		wea.Impact = impact
		wea.Foi = foi
		wea.Balance = balance
		wea.Suitedness = suitedness
		wea.Construction = construction
		wea.Weight = weight
		wea.Cost = cost
		res = append(res, wea)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM weapon WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	wea := Weapon{}
	for selDB.Next() {
		var id int
		var name, wtype, description, puncture, slice, impact, foi, balance, suitedness, construction, weight, cost string
		err = selDB.Scan(&id, &name, &wtype, &description, &puncture, &slice, &impact, &foi, &balance, &suitedness, &construction, &weight, &cost)
		if err != nil {
			panic(err.Error())
		}
		wea.Id = id
		wea.Name = name
		wea.WeaponType = wtype
		wea.Description = description
		wea.Puncture = puncture
		wea.Slice = slice
		wea.Impact = impact
		wea.Foi = foi
		wea.Balance = balance
		wea.Suitedness = suitedness
		wea.Construction = construction
		wea.Weight = weight
		wea.Cost = cost
	}
	tmpl.ExecuteTemplate(w, "Show", wea)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM weapon WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	wea := Weapon{}
	for selDB.Next() {
		var id int
		var name, wtype, description, puncture, slice, impact, foi, balance, suitedness, construction, weight, cost string
		err = selDB.Scan(&id, &name, &wtype, &description, &puncture, &slice, &impact, &foi, &balance, &suitedness, &construction, &weight, &cost)
		if err != nil {
			panic(err.Error())
		}
		wea.Id = id
		wea.Name = name
		wea.WeaponType = wtype
		wea.Description = description
		wea.Puncture = puncture
		wea.Slice = slice
		wea.Impact = impact
		wea.Foi = foi
		wea.Balance = balance
		wea.Suitedness = suitedness
		wea.Construction = construction
		wea.Weight = weight
		wea.Cost = cost
	}
	tmpl.ExecuteTemplate(w, "Edit", wea)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		wtype := r.FormValue("wtype")
		description := r.FormValue("description")
		puncture := r.FormValue("puncture")
		slice := r.FormValue("slice")
		impact := r.FormValue("impact")
		foi := r.FormValue("foi")
		balance := r.FormValue("balance")
		suitedness := r.FormValue("suitedness")
		construction := r.FormValue("construction")
		weight := r.FormValue("weight")
		cost := r.FormValue("cost")

		insForm, err := db.Prepare("INSERT INTO weapon(name, type, description, puncture, slice, impact, foi, balance, suitedness, construction, weight, cost) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, wtype, description, puncture, slice, impact, foi, balance, suitedness, construction, weight, cost)
		log.Println("INSERT: name: " + name + " | wtype: " + wtype + "| description: " + description + "| puncture: " + puncture + "| slice: " + slice + "| impact: " + impact + "| foi: " + foi + "| balance: " + balance + "| suitedness: " + suitedness + "| construction: " + construction + "| weight: " + weight + "| cost: " + cost)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		wtype := r.FormValue("wtype")
		description := r.FormValue("description")
		puncture := r.FormValue("puncture")
		slice := r.FormValue("slice")
		impact := r.FormValue("impact")
		foi := r.FormValue("foi")
		balance := r.FormValue("balance")
		suitedness := r.FormValue("suitedness")
		construction := r.FormValue("construction")
		weight := r.FormValue("weight")
		cost := r.FormValue("cost")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE weapon SET name=?, type=?, description=?, puncture=?, slice=?, impact=?, foi=?, balance=?, suitedness=?, construction=?, weight=?, cost=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, wtype, description, puncture, slice, impact, foi, balance, suitedness, construction, weight, cost, id)
		log.Println("UPDATE: Name: " + name + " | type: " + wtype + "| description: " + description + "| puncture: " + puncture + "| slice: " + slice + "| impact: " + impact + "| foi: " + foi + "| balance: " + balance + "| suitedness: " + suitedness + "| construction: " + construction + "| weight: " + weight + "| cost: " + cost + "| Id: " + id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	wea := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM weapon WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(wea)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
