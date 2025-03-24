package lib

import (
	f "fmt"
)

func BasicGo() {
	f.Println("===============Go Basic Syntax=====================")
	// Basic Data Types
	var studentName string = "James"
	var totalMarks int = 884
	var isDistinction bool = true
	var temperature float32 = 9.5
	f.Printf("Student Name: %s\nTotal Marks: %d\nGot Distinction: %t\n", studentName, totalMarks, isDistinction)
	f.Printf("Temperature is %f", temperature)
	// Math Operators
	additionExample := 2 + 2
	subtractionExample := 5 - 3
	divisionExample := 8 / 4
	multiplicationExample := 10 * 5

	f.Printf("Addition: 2 + 2 = %d\nSubtraction: 5 - 3 = %d\nDivision: 8 / 4 = %d\nMultiplication: 10 * 5= %d\n", additionExample, subtractionExample, divisionExample, multiplicationExample)

	// variable declaration/use
	// no un-used variable allowed
	name := "Spock" // type inferred, shot-hand operator (:=)
	age := 20
	isAlien := true
	f.Printf("Hello, I am %s and I am %d years old\n", name, age)
	f.Printf("Is Alien:  %t\n", isAlien)

	// variable declaration without initialization
	// var city string
	// city = "Vulcanara"
	// f.Printf("I live in %s\n", city)

	// declaring multiple variables with same types all in one go
	var ship, race string = "Enterprise", "Vulcan"

	// declaring multiple variables with varied types
	var (
		isCommander  bool   = true
		numOfHonors  int    = 56
		favoriteFood string = "fish pie"
	)

	f.Printf("Ship: %s\nPlanet: %s\nCommand Status: %t\nNumber of Honors: %d\nFavorit Food: %s\n", ship, race, isCommander, numOfHonors, favoriteFood)

	var defaultInt int       // compiler will assign default value of 0
	var defaultFloat float64 // compiler will assign default value of 0.000000
	var defaultString string // compiler will assign default value of ""
	var defaultBool bool     // compiler will assign default value of 'false'
	f.Printf("Default Int: %d\nDefault Float: %f\nDefault String: %s\nDefault Bool: %t\n", defaultInt, defaultFloat, defaultString,
		defaultBool)

	// constant variables
	const color = "red" // a const value can be left unused

	// declaring multiple const variables with same type, type can be inferred
	const (
		OriginalSeries    = 1966
		TheNextGeneration = 1987
		DeepSpace9        = 1993
		Voyager           = 1995
	)

	f.Printf("OriginalSeries: %d\nTheNextGeneration: %d\nDeepSpace9: %d\nVoyager: %d\n", OriginalSeries, TheNextGeneration, DeepSpace9, Voyager)

	// Enum like structure
	const (
		Jan = iota + 1
		Feb
		Mar
	)
	f.Printf("Jan: %d\nFeb: %d\nMar: %d\n", Jan, Feb, Mar)

	// typed and untyped const are same
	const typedValue int = 10
	const untypedValue = 10
	println(typedValue == untypedValue)

	result := Add(10, 20)
	f.Printf("Result: %d\n", result)

	shoppingTotal := CalculateTotal(24.95, 65.66)
	f.Printf("Shopping Total: %f\n", shoppingTotal)

	sum, product := CalculateSumAndProduct(24, 36)
	f.Printf("Sum: %d\nProduct: %d\n", sum, product)

	_, product2 := CalculateSumAndProduct(2, 6)
	f.Printf("Product2: %d\n", product2)

	//String Manipulation
	firstName := "mark"
	f.Printf("Capitalized Word: %s\n", CapitalizeWord(firstName))
}
