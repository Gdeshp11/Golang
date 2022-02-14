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
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	RWLock.Lock()
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
	RWLock.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	RWLock.Lock()
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
		RWLock.Unlock()
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		RWLock.Unlock()
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	//convert string to float and get error status ok
	priceStr := req.URL.Query().Get("price")
	priceFloat, ok := strconv.ParseFloat(priceStr, 32)

	if ok == nil {
		if _, ok := db[item]; ok {
			fmt.Fprintf(w, "Item already exist in list, please use 'update' request")
		} else {
			RWLock.Lock()
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
	RWLock.Lock()
	if _, itemExist := db[item]; itemExist {
		if ok == nil {
			db[item] = dollars(priceFloat)
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
	item := req.URL.Query().Get("item")
	RWLock.Lock()
	if _, itemExist := db[item]; itemExist {
		delete(db, item)
		RWLock.Unlock()
		fmt.Fprintf(w, "deleted item %s", item)
	} else {
		fmt.Fprintf(w, "%s does not exist in list", item)
		RWLock.Unlock()
	}
}
