package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual : $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Name: %s, Hourly Rate: $%.2f, Hours Worked: %.2f", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommisionEmployee struct {
	Name          string
	BaseSalary    float64
	CommisionRate float64
	SalesAmount   float64
}

func (ce CommisionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommisionRate * ce.SalesAmount)
}

func (ce CommisionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base Salary: $%.2f, Commission Rate: %.2f, Sales Amount: $%.2f)", ce.Name, ce.BaseSalary, ce.CommisionRate, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf(" - Processing : %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	totalPay := 0.0
	for _, employee := range employees {
		pay := employee.CalculatePay()
		PrintEmployeeSummary(employee)
		fmt.Printf(" - Pay: $%.2f\n", pay)
		totalPay += pay
	}
	fmt.Printf("Total Pay for all employees: $%.2f\n", totalPay)
}

func main() {
	employees := []Payable{
		SalariedEmployee{"Alice", 60000},
		HourlyEmployee{"Bob", 25.50, 160},
		CommisionEmployee{"Charlie", 5000, 0.08, 50000},
	}
	ProcessPayroll(employees)
}
