package lib

import (
	f "fmt"
)

func Iteration(){
		f.Println("==================LOOPS==================")
for i :=5; i>= 0; i--{
	f.Println("countdown...", i)
}

// For Loop like a while loop => there is no while keyword in Go
count := 0;
for count <= 5{
	f.Println("count is ", count)
	count++
}

// Infinite Loop
jumps := 0
for {
	if jumps > 5{
		f.Println("Jumps No: ", jumps)
		break
	}
	jumps++
}

}
