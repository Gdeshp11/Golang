package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var RWLock sync.RWMutex

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", mux)) // Listens for curl communication of localhost
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars // database of items with their polar prices

func (db database) list(w http.ResponseWriter, req *http.Request) {
	RWLock.Lock() // Locks in case of multiple reads/writes
	for item, price := range db {
		fmt.Fprintf(w, "\n%s: %s", item, price) //iterates through the database and prints out each item and dollar price
	}
	RWLock.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item") //Gets the item from the query
	RWLock.Lock()                       //locks the system
	if price, ok := db[item]; ok {      // if the item exists
		fmt.Fprintf(w, "%s\n", price) // prints price
		RWLock.Unlock()
	} else {
		w.WriteHeader(http.StatusNotFound)         // 404
		fmt.Fprintf(w, "no such item: %q\n", item) // if the item does not exist write and error
		RWLock.Unlock()
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item") // gets item and price
	//convert string to float and get error status ok
	priceStr := req.URL.Query().Get("price")
	priceFloat, ok := strconv.ParseFloat(priceStr, 32) // converts price into float32, if it cannot, it throws an error

	if ok == nil {
		if _, ok := db[item]; ok { // checks if item alrady exists
			fmt.Fprintf(w, "Item already exist in list, please use 'update' request")
		} else {
			RWLock.Lock() // adds item
			fmt.Fprintf(w, "Adding item: %s ", item)
			fmt.Fprintf(w, "\nprice: %f", priceFloat)
			db[item] = dollars(priceFloat)
			RWLock.Unlock()
		}
	} else {
		fmt.Fprintf(w, "Invalid price")
	}

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	//convert string to float and get error status ok
	priceStr := req.URL.Query().Get("price")
	priceFloat, ok := strconv.ParseFloat(priceStr, 32)
	RWLock.Lock() //Works similarly to create, but instead it checks if the item already exists
	if _, itemExist := db[item]; itemExist {
		if ok == nil {
			db[item] = dollars(priceFloat) // changes the price
			RWLock.Unlock()
			fmt.Fprintf(w, "updated price of item %s: %s", item, dollars(priceFloat))

		} else {
			fmt.Fprintf(w, "Invalid price %s", priceStr)
			RWLock.Unlock()
		}
	} else {
		fmt.Fprintf(w, "%s does not exist in list", item)
		RWLock.Unlock()
	}
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")      //gets the item name
	RWLock.Lock()                            //locks system
	if _, itemExist := db[item]; itemExist { //checks if item exists
		delete(db, item) // deletes the item
		RWLock.Unlock()  // unlocks
		fmt.Fprintf(w, "deleted item %s", item)
	} else {
		fmt.Fprintf(w, "%s does not exist in list", item)
		RWLock.Unlock()
	}
}
