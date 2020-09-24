package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/send-sms/", sendSms)
	p := getPort()
	fmt.Println("listen on", p)
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	p := os.Getenv("PORT")
	if p == "" {
		p = "8080"
	}
	return ":" + p
}

func sendSms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == http.MethodPost {
		s := new(Sms)
		json.NewDecoder(r.Body).Decode(&s)
		err := Send(s)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			// bs, _ := json.Marshal(err)
			// fmt.Fprintf(w, "%s", bs)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}
