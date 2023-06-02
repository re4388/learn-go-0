package account

import (
	"fmt"
	"strconv"
)

type Account struct {
	id  int
	bal float64
}

// go way to impl the OOP concept, using struct + ptr method
// Structs (objects) can also have methods.
// These are called "pointer receiver" methods
// because the first argument is a pointer to a struct of type Account denoted by a *Account.
func (account *Account) String() string {
	// get struct level stuff via account
	// String is the name of the method
	return fmt.Sprintf("Account[%d]: $%0.2f", account.id, account.bal)
}

func (account *Account) Deposit(amt float64) float64 {
	// amt is the arg, float64 is the type
	// the second float64 is the ret type
	account.bal += amt
	return account.bal
}

func (account *Account) Withdraw(amt float64) float64 {
	account.bal -= amt
	return account.bal
}

func (account *Account) Balance() float64 {
	return account.bal
}

func RunAccount1() {
	acc := Account{id: 1, bal: 1000.12}
	acc.Deposit(100)
	fmt.Println(acc.String())

	acc.Withdraw(100)
	fmt.Println(acc.String())

	str := strconv.FormatFloat(acc.Balance(), 'f', -1, 64)
	fmt.Println(str)
}
