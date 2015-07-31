# emailx
[Golang](http://golang.org/) pkg for email validation and normalization.

[![GoDoc](https://godoc.org/github.com/goware/emailx?status.png)](https://godoc.org/github.com/goware/emailx)
[![Travis](https://travis-ci.org/goware/emailx.svg?branch=master)](https://travis-ci.org/goware/emailx)

## Email validation

- Simple email format check (not a complicated regexp, [this is why](http://davidcel.is/posts/stop-validating-email-addresses-with-regex/))
- Resolve the host name

```go
import "github.com/goware/emailx"

func main() {
    err := emailx.Validate("My+Email@wrong.example.com")
    if err != nil {
        fmt.Println("Email is not valid.")

        if err == emailx.ErrInvalidFormat {
            fmt.Println("Wrong format.")
        }

        if err == emailx.ErrUnresolvableHost {
            fmt.Println("Unresolvable host.")
        }
    }
}
```

## Email normalization

```go
import "github.com/goware/emailx"

func main() {
    fmt.Print(emailx.Normalize(" My+Email@example.com. "))
    // Prints my+email@example.com
}
```

## License
Emailx is licensed under the [MIT License](./LICENSE).
