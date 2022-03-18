/*
	To implement crud operations as RPC in go, functions need to satisfy below criterion:
		- Functions need to be methods
		- All functions can be exported
		- Functions need to have two arguments both of them should be of exported types
		- Second argument of the function must be a pointer
		- The return type of the functions must be of error type
*/

package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

type Item struct {
	Title string
	Body  string
}

var database []Item

func (a *API) GetDB(name string, reply *[]Item) error {
	*reply = database
	return nil
}

func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item
	for _, item := range database {
		if item.Title == title {
			getItem = item
			break
		}
	}
	*reply = getItem
	return nil
}

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changedItem Item
	for i, item := range database {
		if item.Title == edit.Title {
			database[i] = Item{edit.Title, edit.Body}
			changedItem = database[i]
			break
		}
	}
	*reply = changedItem
	return nil
}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for i, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}

func main() {
	var api = new(API)

	err := rpc.Register(api)
	if err != nil {
		log.Fatal("Error registering the API", err)
	}
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}

	log.Printf("Serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving", err)
	}

	//fmt.Println("Initial database: ", database)
	// a := Item{"first", "a first item"}
	// b := Item{"second", "a second item"}
	// c := Item{"third", "a third item"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("Second database", database)

	// DeleteItem(b)
	// fmt.Println("Third database", database)

	// EditItem("third", Item{"fourth", "a new item"})
	// fmt.Println("Fourth database", database)

	// x := GetByName("second")
	// y := GetByName("fourth")
	// fmt.Println(x, y)
}
