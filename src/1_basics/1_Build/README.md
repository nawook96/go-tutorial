# Create Go source file
Make **hello.go** file like /GOPATH/src/hello.go

```GOPATH/src/hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Execute the following command in terminal
```
go build hello.go
go run hello.go
```

**Golang** consists of all parts of the package. Therefore, package setting is placed on the top of source codes. Also, the main package is the first to run.

```
package main
```

**import** is a keyword that uses packages. We use `fmt` package to print strings.

```
import "fmt"
```

**Like** C, C++, Java, and C#, start with the main function. And using the Println function of the fmt package, Outputs a string "Hello, world!".

```
func main() {
    fmt.Println("Hello, world!")
}
```