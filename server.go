/**
 * Created by chaolinding on 2018/6/12.
 */

package main

import (
	"fmt"
	"net/http"
)


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello go + docker !")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server start on 8888!")
	http.ListenAndServe(":8888", nil)
}
