package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Wildan",
		"address": "Jakarta",
	}

	person["title"] = "Junior Software Engineer"
	person["age"] = "23"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["id"] = "B001"
	book["title"] = "Book of All-Book"
	book["author"] = "You"
	book["pages"] = "300"
	book["released"] = "2021"
	fmt.Println("New Book")
	fmt.Println(book["id"])
	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["pages"])
	fmt.Println(book["released"])

	delete(book, "id")
	fmt.Println("After ID is Deleted")
	fmt.Println(book["id"], "ID is Deleted :p")
	fmt.Println(book["title"])
	fmt.Println(book["author"])
	fmt.Println(book["pages"])
	fmt.Println(book["released"])
}
