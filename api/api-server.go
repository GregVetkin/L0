package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"

	"l0/models/nats"
	"l0/models/cache"
	"github.com/nats-io/stan.go"
	"l0/models/order"
	"io"
	"encoding/json"
)


type App struct {
	sc	stan.Conn
	ch	cache.Cache
}

func (a *App) addOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Only POST method allowed", 405)
		return
	}
	bodyBytes, err := io.ReadAll(r.Body)
    if err != nil {
        log.Println("Error:", err)
		return
    }
	
	var o order.Order
	err = json.Unmarshal(bodyBytes, &o)
	if err != nil {
		log.Println("Error!:", err)
		http.Error(w, "Bad data", 400)
		return
	}
	
	a.ch.StoreOrder(o.OrderUID, bodyBytes)
	a.sc.Publish("addOrder", bodyBytes)
}


func (a *App) getOrder(w http.ResponseWriter, r *http.Request) {
	uid := r.URL.Query().Get("uid")
	
	data, err := a.ch.TakeOrder(uid)
	if err != nil {
		http.Error(w, "No data with this uid", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}




func main() {
	sc, err := nats.NewServer("api")
	if err != nil {
		log.Fatal(err)
	}
	
	ch := cache.NewCache()
	err = ch.Fill()
	if err != nil {
		log.Fatal(err)
	}
	
	
	var Application App = App{sc: sc, ch: ch}
	

	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
		
	http.HandleFunc("/", ShowMainPage)
	http.HandleFunc("/order", Application.getOrder)
	http.HandleFunc("/order/new", Application.addOrder)

	fmt.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}


func ShowMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./static/html/mainpage.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
	

}