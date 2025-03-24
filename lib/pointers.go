package lib

import (
	st "beginnerGo/lib/structs"
	f "fmt"
)

func Pointers() {
	passport2 := st.Student{
		Name:    "Passport2",
		Age:     13,
		Year:    8,
		Address: "124 Main Road",
		Contact: "123456788",
		Catchup: true,
	}
	teddy := st.Student{
		Name:    "Teddy",
		Age:     13,
		Year:    8,
		Address: "123 Main Road",
		Contact: "123456789",
		Catchup: true,
	}

	teddyPtr := &teddy
	f.Printf("Pointer of Teddy: %p\n", teddyPtr)
	f.Printf("Name Before: %s\n", teddy.Name)
	teddy.ChangeName("John Cleese")
	f.Printf("Name After: %s\n", teddy.Name)

	f.Println("-------------------------------")
	f.Printf("Name Before: %s\n", passport2.Name)
	passport2.ChangeName("Filios Fog")
	f.Printf("Name After: %s\n", passport2.Name)
	juice := "apple juice"
	juicePointer := &juice
	f.Printf("Value of Juice: %s\nMemory address of Juice: %p\n", juice, juicePointer)
	*juicePointer = "Blackcurrent Juice"
	f.Printf("Value of Juice: %s\nMemory address of Juice: %p\n", juice, juicePointer)

}
