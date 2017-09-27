package main

import "fmt"

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

func (e Employee) EmpInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)

	for _, info := range e.MonthlySalary {
		fmt.Println("==============")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}

	return "----------------"
}

func main() {

	e := Employee{
		FirstName: "Leonardo",
		LastName:  "Souza",
		Email:     "leogsouza@gmail.com",
		Age:       35,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 12000.00,
				HRA:   4000.00,
				TA:    1000.00,
			},
			Salary{
				Basic: 11000.00,
				HRA:   2500.00,
				TA:    2100.00,
			},
		},
	}

	fmt.Println(e.EmpInfo())
}
