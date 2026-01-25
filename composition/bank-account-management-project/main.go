package main

import "fmt"

type Account struct {
	AccountNumber string
	Balance       float64
	OwnerName     string
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	acc.Balance += amount
	fmt.Printf("Deposited: %.2f, New Balance: %.2f\n", amount, acc.Balance)
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	if amount > acc.Balance {
		return fmt.Errorf("insufficient funds: available balance is %.2f", acc.Balance)
	}
	acc.Balance -= amount
	fmt.Printf("Withdrawn: %.2f, New Balance: %.2f\n", amount, acc.Balance)
	return nil
}

func (acc *Account) GetBalance() float64 {
	return acc.Balance
}

func (acc *Account) String() string {
	return fmt.Sprintf("Account{Number: %s, Owner: %s, Balance: %.2f}", acc.AccountNumber, acc.OwnerName, acc.Balance)
}

type SavingsAccount struct {
	Account
	InterestRate float64
}

func (sa *SavingsAccount) AddInterest() {
	interest := sa.Balance * sa.InterestRate / 100
	if err := sa.Deposit(interest); err != nil {
		fmt.Printf("Error adding interest: %v\n", err)
	} else {
		fmt.Printf("Interest added: %.2f\n", interest)
	}
}

type OverDraftAccount struct {
	Account
	OverdraftLimit float64
}

func (oda *OverDraftAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdrawal amount must be positive")
	}
	availableBalance := oda.Balance + oda.OverdraftLimit
	if amount > availableBalance {
		return fmt.Errorf("insufficient funds: available balance with overdraft limit is %.2f", availableBalance)
	}
	oda.Balance -= amount
	if oda.Balance < 0 {
		fmt.Printf("Withdrawn: %.2f, New Balance: %.2f (Overdraft used: %.2f)\n", amount, oda.Balance, -oda.Balance)
	} else {
		fmt.Printf("Withdrawn: %.2f, New Balance: %.2f\n", amount, oda.Balance)
	}
	return nil
}

func main() {
	// Create a regular account
	acc := &Account{
		AccountNumber: "ACC001",
		OwnerName:     "John Doe",
		Balance:       1000.00,
	}
	fmt.Println("=== Regular Account ===")
	fmt.Println(acc)
	if err := acc.Deposit(500); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if err := acc.Withdraw(200); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Current Balance: %.2f\n\n", acc.GetBalance())

	// Create a savings account
	sa := &SavingsAccount{
		Account: Account{
			AccountNumber: "SAV001",
			OwnerName:     "Jane Smith",
			Balance:       5000.00,
		},
		InterestRate: 3.5,
	}
	fmt.Println("=== Savings Account ===")
	fmt.Println(&sa.Account)
	if err := sa.Deposit(1000); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	sa.AddInterest()
	fmt.Printf("Current Balance: %.2f\n\n", sa.GetBalance())

	// Create an overdraft account
	oda := &OverDraftAccount{
		Account: Account{
			AccountNumber: "ODA001",
			OwnerName:     "Bob Johnson",
			Balance:       2000.00,
		},
		OverdraftLimit: 500.00,
	}
	fmt.Println("=== Overdraft Account ===")
	fmt.Println(&oda.Account)
	if err := oda.Deposit(500); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if err := oda.Withdraw(2300); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	if err := oda.Withdraw(300); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("Current Balance: %.2f\n", oda.GetBalance())
}
