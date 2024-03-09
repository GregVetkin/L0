package database

import (
	"log"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)


const (
	dbhost string = "pgdb-container"
	dbuser string = "l0user"
	dbpass string = "l0pass"
	dbname string = "l0"
)

type Order struct {
	Uid		string
	Data	[]byte
}


type Postgres struct {
	DB *sql.DB
}

func (pg *Postgres) AddOrder(uid string, data []byte) error {
	_, err := pg.DB.Exec("INSERT INTO public.orders (uid, orderjson) VALUES ($1, $2)", uid, data)
	return err
}

func (pg *Postgres) AllOrders() ([]Order, error) {
	orders := []Order{}
	
	rows, err := pg.DB.Query("SELECT uid, orderjson FROM public.orders")
    if err != nil {
		log.Print(err)
        return orders, fmt.Errorf("Error")
    }
    defer rows.Close()
    
    
    for rows.Next(){
        o := Order{}
        err := rows.Scan(&o.Uid, &o.Data)
        if err != nil{
            log.Print(err)
            continue
        }
        orders = append(orders, o)
    }
	return orders, nil
}



func OpenDataBase() (*Postgres, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbhost, dbuser, dbpass, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return &Postgres{DB: db}, nil
}


type Orders interface {
	AddOrder(uid string, data []byte) error
	AllOrders() ([]Order, error)
}






	