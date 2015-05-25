package emailx_test

import (
	"fmt"
	"testing"

	"github.com/goware/emailx"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in  string
		out string
		err bool
	}{
		// Invalid format.
		{in: "", err: true},
		{in: "email@", err: true},
		{in: "email@x", err: true},
		{in: "email@@example.com", err: true},
		{in: "email@at@example.com", err: true},
		{in: "some whitespace@example.com", err: true},
		{in: "email@whitespace example.com", err: true},

		// Unresolvable domain.
		{in: "email+extra@wrong.example.com", err: true},

		// Valid.
		{in: "email@example.com"},
		{in: "email+extra@example.com"},
	}

	for _, tt := range tests {
		err := emailx.Validate(tt.in)
		if err != nil {
			if !tt.err {
				t.Errorf(`"%s": unexpected error \"%v\"`, tt.in, err)
			}
			continue
		}
		if tt.err && err == nil {
			t.Errorf(`"%s": expected error`, tt.in)
			continue
		}
	}
}

func ExampleValidate() {
	err := emailx.Validate("My+Email@example.com")
	if err != nil {
		fmt.Print("Email is not valid.")

		if err == emailx.ErrInvalidFormat {
			fmt.Print("Wrong format.")
		}

		if err == emailx.ErrUnresolvableHost {
			fmt.Print("Unresolvable host.")
		}
	}
}
