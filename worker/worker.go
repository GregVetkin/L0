package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	
	
	"l0/models/database"
	"l0/models/nats"
	"l0/models/order"
	"github.com/nats-io/stan.go"
)


func workerStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}




type App struct {
	sc	stan.Conn
	db	database.Orders
}

func (a *App) addOrder(m *stan.Msg) {
	var err error
	var o order.Order
	
	err = json.Unmarshal(m.Data, &o)
	if err != nil {
		log.Println("Error unmarshal:", err)
		return
	}
	
	err = a.db.AddOrder(o.OrderUID, m.Data)
	if err != nil {
		log.Println("Error database:", err)
	}
	fmt.Println(o.OrderUID)
	
}




func main() {
	sc, err := nats.NewServer("worker")
	if err != nil {
		log.Fatal(err)
	}
	
	db, err := database.OpenDataBase()
	if err != nil {
		log.Fatal(err)
	}
	
	
	var Application App = App{sc: sc, db: db}
	Application.sc.Subscribe("addOrder", Application.addOrder, stan.DurableName("addOrder-durable"))


	fmt.Println("Server listening on port 8181...")
	http.HandleFunc("/status", workerStatus)
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Fatal(err)
	}
}