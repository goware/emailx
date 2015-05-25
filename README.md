# URLx
[Golang](http://golang.org/) pkg for email validation and normalization.

[![GoDoc](https://godoc.org/github.com/goware/emailx?status.png)](https://godoc.org/github.com/goware/emailx)
[![Travis](https://travis-ci.org/goware/emailx.svg?branch=master)](https://travis-ci.org/goware/emailx)

## Email validation

```go
import "github.com/goware/emailx"

func main() {
    if !emailx.IsValid("My+Email@example.com") {
        // Email is not valid!
    }
}
```

## Email normalization

```go
import "github.com/goware/emailx"

func main() {
    normalized, err := emailx.Normalize(" My+Email@example.com. ")
    if err != nil {
        // Handle error.
    }

    fmt.Print(normalized)
    // Prints my+email@example.com
}
```

## License
Emailx is licensed under the [MIT License](./LICENSE).
