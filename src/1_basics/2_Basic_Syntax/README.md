# Curly bracket({})
---
You must start a square bracket on the same line when you start a function, condition statement, or loop statement, such as:

```
func main() {
	i := 10
	if i > 5 {
		fmt.Println("above 5")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
```

# Semicolon
---
Unlike C,C++, you usually omit semicolon at the end of the phrase in **Golang**.

```
fmt.Println("Hello, world!")
```

If you want to use multiple phrases per line, you can separate them with semicolons.

```
fmt.Println("Hello,");fmt.Println("world!")
```

# Comment out
---
The Comment out has two type, one line and range.

```
// fmt.Println("Hello, world!")
```
```
/*
fmt.Println("Hello,")
fmt.Println("world!")
*/
```