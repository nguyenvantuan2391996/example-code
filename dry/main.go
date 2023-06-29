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
	time.Now().AddDate(0, 1, 1)
	time.Now().AddDate(0, 0, NumberOfDaysAdditional)

	// Sử dụng kỹ thuật tái sử dụng
}

func GetElementRandomFromArray(arr []string) string {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := len(arr) - 1

	return arr[rand.Intn(max-min+1)+min]
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
