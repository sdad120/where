package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	token := "catlaw"
	hashList := []string{
		token,
		r.Form.Get("timestamp"),
		r.Form.Get("nonce"),
	}
	sort.Strings(hashList)
	hashCode := sha1.Sum([]byte(hashList[0] + hashList[1] + hashList[2]))
	hashString := string(hashCode[:])
	fmt.Printf("%s, %s, %s", r.Form.Get("timestamp"), r.Form.Get("nounce"), hashString)
	if hashString == r.Form.Get("signature") {
		fmt.Fprint(w, r.Form.Get("echostr"))
	} else {
		fmt.Fprint(w, "")
	}
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/wx", AuthHandler)
	http.ListenAndServe(":80", nil)
}
