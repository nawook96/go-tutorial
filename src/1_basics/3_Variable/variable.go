package main

func main() {
	// Declaration
	var i int
	var s string

	var age1 int = 10
	var name1 string = "Shin"

	// Short declaration
	var age2 = 20
	var name2 = "Kim"
	// var address // compiled error

	age3 := 30
	name3 := "Park"

	// Multiple declaration
	var x1, y1 int = 30, 50                      // x = 30, y = 50
	var age4, name4 = 40, "Choi"                 // age = 10, name = "Choi"
	a, b, c, d := 1, 3.4, "Hello, world!", false // a = 1, b = 3.4, c = "Hello, world!", d = false.

	var (
		x2, y2      int = 30, 50    // x and y : int.
		age5, name5     = 50, "Lee" // age : int, name : string
	)

	// Parallel assignment
	var x3, y3 int
	var age6 int

	x4, y4, age6 = 10, 20, 5 // x = 10, y = 20, age = 5: Parallel assignment

}
