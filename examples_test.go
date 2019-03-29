package errors_test

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/adrg/errors"
)

func ExampleNew() {
	err := errors.New("something bad happened")

	fmt.Printf("%s\n", err)  // Print error message only.
	fmt.Printf("%v\n", err)  // Print error chain without stack details.
	fmt.Printf("%+v\n", err) // Print error chain with stack details.
}

func ExampleErrorf() {
	err := errors.Errorf("invalid user ID: %d", 0)

	fmt.Printf("%s\n", err)  // Print error message only.
	fmt.Printf("%v\n", err)  // Print error chain without stack details.
	fmt.Printf("%+v\n", err) // Print error chain with stack details.
}

func ExampleAnnotate() {
	_, err := strconv.Atoi("this will fail")

	// No need to check if the original error is nil.
	// If err is nil, Annotate returns nil.
	err = errors.Annotate(err, "conversion failed")

	fmt.Printf("%s\n", err)  // Print error message only.
	fmt.Printf("%v\n", err)  // Print error chain without stack details.
	fmt.Printf("%+v\n", err) // Print error chain with stack details.
}

func ExampleAnnotatef() {
	var err error

	email := "alice @example.com"
	if strings.Contains(email, " ") {
		err = errors.New("email address cannot contain spaces")
	}

	// No need to check if the original error is nil.
	// If err is nil, Annotatef returns nil.
	err = errors.Annotatef(err, "invalid email adddress %s", email)

	fmt.Printf("%s\n", err)  // Print error message only.
	fmt.Printf("%v\n", err)  // Print error chain without stack details.
	fmt.Printf("%+v\n", err) // Print error chain with stack details.
}

func ExampleWrap() {
	errInvalid := errors.New("invalid data")

	_, err := strconv.Atoi("this will fail")
	if err != nil {
		err = errors.Wrap(errInvalid, err)
	}

	fmt.Printf("%s\n", err)  // Print error message only.
	fmt.Printf("%v\n", err)  // Print error chain without stack details.
	fmt.Printf("%+v\n", err) // Print error chain with stack details.
}

func ExampleUnwrap() {
	err := errors.Annotate(errors.New("first error"), "second error")
	if next := errors.Unwrap(err); next != nil {
		fmt.Printf("%s\n", next)  // Print error message only.
		fmt.Printf("%v\n", next)  // Print error chain without stack details.
		fmt.Printf("%+v\n", next) // Print error chain with stack details.
	}
}

func ExampleIs() {
	errInvalid := errors.New("invalid data")

	err := errors.Annotate(errInvalid, "validation failed")
	if errors.Is(err, errInvalid) {
		fmt.Println("err is invalidErr")
	}
}

func ExampleAs() {
	_, err := strconv.Atoi("this will fail")

	// No need to check if the original error is nil.
	// If err is nil, Annotate returns nil.
	err = errors.Annotate(err, "conversion failed")

	nErr := &strconv.NumError{}
	if ok := errors.As(err, &nErr); ok {
		fmt.Printf("error %s: parsing `%s`: %s\n", nErr.Func, nErr.Num, nErr.Err)
	}
}

func ExampleAs_interface() {
	err := errors.NewHTTP(nil, 500, "something bad happened")
	err = errors.Annotate(err, "annotated HTTP error interface")

	var nErr errors.HTTP
	if ok := errors.As(err, &nErr); ok {
		fmt.Printf("%s\n", err)        // Print error message only.
		fmt.Printf("%v\n", err)        // Print error chain without stack details.
		fmt.Printf("%+v\n", err)       // Print error chain with stack details.
		fmt.Println(errors.Code(nErr)) // Print error code
	}
}

func ExampleCode() {
	err := errors.HTTPf(nil, http.StatusNotFound, "invalid endpoint")

	// Print error code.
	fmt.Println(errors.Code(err))
}

func ExampleNewHTTP() {
	_, err := strconv.Atoi("this will fail")
	if err != nil {
		err = errors.NewHTTP(err, http.StatusBadRequest, "invalid data")

		fmt.Printf("%s\n", err)       // Print error message only.
		fmt.Printf("%v\n", err)       // Print error chain without stack details.
		fmt.Printf("%+v\n", err)      // Print error chain with stack details.
		fmt.Println(errors.Code(err)) // Print error code.
	}
}

func ExampleHTTPf() {
	validateEmail := func(email string) error {
		if strings.Contains(email, " ") {
			return errors.New("email address cannot contain spaces")
		}

		return nil
	}

	email := "alice @example.com"
	if err := validateEmail(email); err != nil {
		err = errors.HTTPf(err, http.StatusBadRequest, "invalid email %s", email)

		fmt.Printf("%s\n", err)       // Print error message only.
		fmt.Printf("%v\n", err)       // Print error chain without stack details.
		fmt.Printf("%+v\n", err)      // Print error chain with stack details.
		fmt.Println(errors.Code(err)) // Print error code.
	}
}
