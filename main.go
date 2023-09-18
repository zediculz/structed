package main

import (
	"fmt"
)

// bank management

type account struct {
	name     string
	number   float64
	balance  int
	limit    int
}

func (a account) CreateAcct() {
	fmt.Println("account created successfully, account details => : ", a)
}


func (a account) Transfer(amt int) {
	if amt >= a.limit {
		fmt.Println("Limit exceeded, limit: ", a.limit)
	}else if amt >= a.balance {
		fmt.Println("Balance exceeded, Balance: ", a.balance)
	}else {
		a.limit = a.limit - amt
		a.balance = a.balance - amt
		fmt.Println("Transfer successful, Balance: ", a)
	}

}

func (a account) _() {
	fmt.Println("account created successfully, account details => : ", a)
}

func main() {
	var user = account{
		name: "johndoe",
		number: 12345,
		limit: 100 ,
		balance: 56,
	}

	user.CreateAcct()

	user.Transfer(16)
}
