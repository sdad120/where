package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

func hexStringToBytes(s string) [20]byte {
	bs := make([]byte, 0)
	for i := 0; i < len(s); i = i + 2 {
		b, _ := strconv.ParseInt(s[i:i+2], 16, 16)
		bs = append(bs, byte(b))
	}

	var fixBs [20]byte
	copy(fixBs[:], bs[:20])

	return fixBs
}

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
	signatureCode := hexStringToBytes(r.Form.Get("signature"))
	if hashCode == signatureCode {
		fmt.Fprint(w, r.Form.Get("echostr"))
	} else {
		fmt.Fprint(w, "")
	}
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	bodyString, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", bodyString)
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	// http.HandleFunc("/wx", AuthHandler)
	http.HandleFunc("/wx", MessageHandler)
	http.ListenAndServe(":80", nil)
}
