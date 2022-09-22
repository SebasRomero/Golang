package main

import (
	"fmt"

	"github.com/example/users"
)

func main() {
	var frank users.Cashier = users.NewEmployee("Frank")
	var robert users.Admin = users.NewEmployee(("Robert"))

	total := frank.CalcTotal(10, 20, 30)
	fmt.Println(total)

	robert.DesactivateEmployee(frank)
	fmt.Println(frank.CalcTotal(40, 50, 60))
}
