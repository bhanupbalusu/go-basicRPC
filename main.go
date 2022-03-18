package main

import "fmt"

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item
	for _, item := range database {
		if item.title == title {
			getItem = item
			break
		}
	}
	return getItem
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, newItem Item) Item {
	var changedItem Item
	for i, item := range database {
		if item.title == title {
			database[i] = newItem
			changedItem = newItem
			break
		}
	}
	return changedItem
}

func DeleteItem(item Item) Item {
	var del Item
	for i, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:i], database[i+1:]...)
			del = item
			break
		}
	}
	return del
}

func main() {
	fmt.Println("Initial database: ", database)
	a := Item{"first", "a first item"}
	b := Item{"second", "a second item"}
	c := Item{"third", "a third item"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("Second database", database)

	DeleteItem(b)
	fmt.Println("Third database", database)

	EditItem("third", Item{"fourth", "a new item"})
	fmt.Println("Fourth database", database)

	x := GetByName("second")
	y := GetByName("fourth")
	fmt.Println(x, y)
}
