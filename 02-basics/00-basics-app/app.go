// every program must have a package.
// the main package is the entry point of the program.
package main

// this allows us to use Print function from the fmt package.
import "fmt"

// must be named main. Entry point of the program.
func main() {
	fmt.Print("Hello World")
}

// go build command will compile and create an executable file.
// one module can have multiple packages.
// to create project run `go mod init <module_name>`
// to run the program use `go run <file_name.go>`
// to build the program use `go build <file_name.go>`

// to create exe for windows use the following command:
// GOOS=windows GOARCH=amd64 go build -o myapp.exe app.go

// when we build 3rd party package we don't need main package.
