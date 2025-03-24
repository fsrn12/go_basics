package lib

import (
	st "beginnerGo/lib/structs"
	f "fmt"
	"reflect"
)

type Car struct { // Defined Struct
	color    string
	model    string
	make     string
	year     int
	doors    int
	isHybrid bool
}

type Address struct {
	Street string
	City   string
}

type Contact struct {
	Name    string
	Phone   string
	Email   string
	Address Address
}

func AdvancedDS() {
	pf := f.Printf
	pln := f.Println
	//Structs
	pln("=========Advanced Data Structures - Struct, Pointers=========")

	kim := st.Student{
		Name:    "Kim Chang",
		Age:     13,
		Year:    8,
		Address: "123 Home Road",
		Contact: "123456789",
		Catchup: false,
	}
	ella := st.Student{
		Name:    "Ella Jan",
		Age:     14,
		Year:    8,
		Address: "124 Home Road",
		Contact: "123456788",
		Catchup: true,
	}
	bee := st.Student{
		Name:    "Bellatrix La Strange",
		Age:     10,
		Year:    5,
		Address: "60 Home Road",
		Contact: "236541879",
		Catchup: true,
	}

	pf("First position in class %+v\n", ella)
	pf("Second position in class %+v\n", kim)
	pf("Third position in class %+v\n", bee)

	honda := Car{
		color:    "Pearl White",
		model:    "Civic",
		make:     "Honda",
		year:     2024,
		doors:    4,
		isHybrid: true,
	}

	pf("Car details: %v\n", honda)

	// Anonymous Struct
	ship := struct {
		name  string
		speed int
	}{
		name:  "NCC1701-C",
		speed: 11,
	}
	pf("Ship model: %s\nShip speed: %d\n", ship.name, ship.speed)

	// Nested Structs
	liz := Contact{
		Name:  "Liz Xing",
		Phone: "193847857",
		Email: "liz@email.com",
		Address: Address{
			Street: "111 Home Park Lane",
			City:   "Asgard",
		},
	}

	emma := Contact{
		Name:  "Emma Bostian",
		Email: "emma@mail.com",
		Phone: "02877272722",
		Address: Address{
			Street: "65 The Main Street",
			City:   "Manchester",
		},
	}

	emma.PrintContact()
	liz.PrintContact()
	// pf("Emma's Details: %+v\n", emma)
	// pf("Emma's Details: %+v\n", liz)

}

func (c Contact) PrintContact() {
	values := reflect.ValueOf(c)
	types := values.Type()
	// f.Println("Number of fields (Length of struct): ", values.NumField())        // Number of fields
	// f.Println("Index of Name Key in struct lizContact", types.Field(0).Index[0]) // Key of the struct at given index
	// f.Printf("Values; %v\nTypes: %v\n", values, types)

	for i := 0; i < values.NumField(); i++ {
		f.Println(types.Field(i).Name, values.Field(i))
	}
}
