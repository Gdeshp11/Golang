package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodbEndpoint = "mongodb://172.28.187.46:32103" // Find this from the Mongo container
)

type Post struct {
	ID           primitive.ObjectID `bson:"_id"`
	Product_name string             `bson:"product_name"`
	Price        dollars            `bson:"price"`
	CreatedAt    time.Time          `bson:"created_at"`
	Tags         string             `bson:"tags"`
}

var RWLock sync.RWMutex
var ctx context.Context
var col *mongo.Collection

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(mongodbEndpoint),
	)
	checkError(err)
	// Connect to mongo
	ctx = context.Background()
	err = client.Connect(ctx)
	checkError(err)
	// Disconnect
	defer client.Disconnect(ctx)

	// select collection from database
	col = client.Database("blog").Collection("posts")

	mux := http.NewServeMux()
	mux.HandleFunc("/list", list)
	mux.HandleFunc("/price", price)
	mux.HandleFunc("/create", create)
	mux.HandleFunc("/update", update)
	mux.HandleFunc("/delete", delete)
	log.Fatal(http.ListenAndServe(":8000", mux)) // Listens for curl communication of localhost
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// type database map[string]dollars // database of items with their polar prices

func list(w http.ResponseWriter, req *http.Request) {

	// filter posts tagged as golang
	filter := bson.M{"tags": "products"}

	// find all documents
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	// iterate through all documents
	for cursor.Next(ctx) {
		var p Post
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "Item:%s, Price:%s\n", p.Product_name, p.Price)
	}

	// check if the cursor encountered any errors while iterating
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

}

//TODO get price of item from mongodb
func price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item") //Gets the item from the query
	fmt.Fprintf(w, "Item:%s\n", item)

	//filter based on product name
	filter := bson.M{"product_name": item}

	// find one document
	var p Post
	if err := col.FindOne(ctx, filter).Decode(&p); err != nil {
		fmt.Fprint(w, "ERROR:", err)
		// fmt.Fprintf(w, "no such item: %q\n", item) // if the item does not exist write and error

	} else {
		fmt.Printf("post: %+v\n", p)
		fmt.Fprintf(w, "Price of %s:%s\n", item, p.Price)
	}
}

func create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item") // gets item and price
	//convert string to float and get error status ok
	priceStr := req.URL.Query().Get("price")
	priceFloat, ok := strconv.ParseFloat(priceStr, 32) // converts price into float32, if it cannot, it throws an error

	if ok == nil {
		fmt.Fprintf(w, "Adding item: %s ", item)
		fmt.Fprintf(w, "\nprice: %f", priceFloat)
		// Insert one
		res, err := col.InsertOne(ctx, &Post{
			ID:           primitive.NewObjectID(),
			Product_name: item,
			Price:        dollars(priceFloat),
			CreatedAt:    time.Now(),
			Tags:         "products",
		})

		if err == nil {
			fmt.Printf("inserted id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
		}

	} else {
		fmt.Fprintf(w, "Invalid price")
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	//convert string to float and get error status ok
	priceStr := req.URL.Query().Get("price")
	priceFloat, ok := strconv.ParseFloat(priceStr, 32)

	// adds item

	if ok == nil {
		fmt.Fprintf(w, "Updating item: %s ", item)
		fmt.Fprintf(w, "\nprice: %f", priceFloat)
		filter := bson.D{{"product_name", item}}
		update := bson.D{{"$set",
			bson.D{
				{"price", priceFloat},
			},
		}}

		res, err := col.UpdateOne(ctx,
			filter,
			update)

		if err == nil {
			fmt.Println("update count: ", res.ModifiedCount)
			fmt.Fprintf(w, "updated price of item:%s-%s\n", item, priceStr)
		} else {
			fmt.Fprint(w, "Error:\n", err)
		}
	} else {
		fmt.Fprintf(w, "Invalid price")
	}
}

func delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item") //gets the item name
	res, err := col.DeleteMany(ctx, bson.M{"product_name": item})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintln(w, "delete count: ", res.DeletedCount)
	}
}
