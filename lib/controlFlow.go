package lib

import f "fmt"

func ControlFlow(){
	f.Println("===============Control Flow=====================")
adult, child, senior := 20, 12, 13
age := 30
if age <= 10 {
	f.Printf("You are a child, You Pay: $%d\n", child)
} else if age >= 60 {
	f.Printf("You are a senior, You Pay: $%d\n", senior)
}else {
	f.Printf("You are an adult, You Pay: $%d\n", adult)
}


// Switch Statements
day := "Tuesday"
switch day{
	case "Monday", "Tuesday", "Wednesday":
	f.Printf("Today is %s, Sports Club\n", day)
	case "Thursday":
	f.Printf("Today is %s, Boxing and Circuit Club\n", day)
	case "Friday":
	f.Printf("Today is %s, Coding Club\n", day)
	default:
	f.Printf("Today is %s, Homework\n", day)
}
}
