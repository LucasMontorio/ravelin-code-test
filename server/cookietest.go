package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.ListenAndServe(":8181", http.HandlerFunc(Test))
}

func Test(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:		"foo",
		Value:		"bar",
		Expires:	time.Now().Add(1 * time.Hour),
		Domain:		".godev.local",	// edit (or omit)
		Path:		"/",		// ^ ditto
		HttpOnly:	true,
	}
	fmt.Fprintln(w, "Hello world")
	http.SetCookie(w, c)
}
