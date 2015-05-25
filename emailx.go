package emailx

import (
	"errors"
	"net"
	"regexp"
	"strings"
)

var (
	ErrInvalidFormat    = errors.New("wrong email format")
	ErrUnresolvableHost = errors.New("wrong email format")

	userRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
	hostRegexp = regexp.MustCompile("^[^\\s]+\\.[^\\s]+$")
)

// Validate checks format of a given email and resolves its host name.
func Validate(email string) error {
	at := strings.LastIndex(email, "@")
	if at <= 0 || at > len(email)-3 {
		return ErrInvalidFormat
	}

	user := email[:at]
	host := email[at+1:]

	if !userRegexp.MatchString(user) || !hostRegexp.MatchString(host) {
		return ErrInvalidFormat
	}

	_, err := net.ResolveIPAddr("ip", host)
	if err != nil {
		return ErrUnresolvableHost
	}

	return nil
}

// Normalize normalizes email address.
func Normalize(email string) string {
	// Trim whitespaces.
	email = strings.TrimSpace(email)

	// Trim extra dot in hostname.
	email = strings.TrimRight(email, ".")

	// Lowercase.
	email = strings.ToLower(email)

	return email
}
