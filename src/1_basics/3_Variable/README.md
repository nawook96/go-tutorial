# Declaration
---
There are two ways to declare a variable in the **Golang**: using the `var` keyword and omitting the template.

```
var i int
var s string

var age1 int = 10
var name1 string = "Shin"
```

If you omit data-type, the variable is determined by data-type of value that you enter. However, if you omit the data-type, you must define the initial value.

```
var age2 = 20
var name2 = "Kim"
var address // compiled error
```

# Short declaration
---
`:=` allows you to simply declare and initialize variables without using `var` and `type` keywords.

```
age3 := 30
name3 := "Park"
```

# Multple declaration
---
you can declare multiple variables separated by `,`(comma). Then, values are subsituted in the order in which variables are declared and must be equal to number of variables.
```
var x1, y1 int = 30, 50       // x = 30, y = 50
var age4, name4 = 40, "Choi" // age = 10, name = "Choi"
a, b, c, d := 1, 3.4, "Hello, world!", false // a = 1, b = 3.4, c = "Hello, world!", d = false.
```

You can enter values for multiple variables even after a variable is declared. In **Golang**, this feature is called `Parallel assignment`.

```
var x3, y3 int
var age6 int

x4, y4, age = 10, 20, 5  // x = 10, y = 20, age = 5: Parallel assignment
```

`var` keyword and `()` allow you to declare and initialize multiple variables.

```
var (
	x2, y2      int = 30, 50      // x and y : int.
	age5, name5     = 10, "Maria" // age : int, name : string
)
```