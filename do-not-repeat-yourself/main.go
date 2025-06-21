package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type User struct {
	ID   int
	Name string
}

type Product struct {
	ID          int
	ProductName string
}

const (
	NumberOfDaysAdditional = 1
)

func main() {
	// Modularization
	user, product := GetData(1, 1)
	fmt.Println(user)
	fmt.Println(product)

	user2 := GetInformationFromUserService(1)
	product2 := GetInformationFromProductService(2)

	fmt.Println(user2)
	fmt.Println(product2)

	// Functions và Classes
	fmt.Println(GetElementRandomFromArray([]string{"1", "2", "3"}))

	// Libraries và Frameworks
	fmt.Println(strings.Contains("Tuan Nguyen", "tu"))

	// Sử dụng biến và hằng số
	fmt.Println(time.Now().AddDate(0, 1, 1))
	fmt.Println(time.Now().AddDate(0, 0, NumberOfDaysAdditional))
}

func GetElementRandomFromArray(arr []string) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	minValue := 0
	maxvalue := len(arr) - 1

	return arr[r.Intn(maxvalue-minValue+1)+minValue]
}

func GetInformationFromUserService(id int) *User {
	return &User{
		ID:   id,
		Name: fmt.Sprintf("%v-%v", "Tuan Nguyen", id),
	}
}

func GetInformationFromProductService(id int) *Product {
	return &Product{
		ID:          id,
		ProductName: fmt.Sprintf("%v-%v", "iPhone 14 pro max", id),
	}
}

func GetData(userID, productID int) (*User, *Product) {
	user := &User{
		ID:   userID,
		Name: fmt.Sprintf("%v-%v", "Tuan Nguyen", userID),
	}

	product := &Product{
		ID:          productID,
		ProductName: fmt.Sprintf("%v-%v", "iPhone 14 pro max", productID),
	}

	return user, product
}
