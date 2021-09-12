package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const port = "8080"

type Host struct {
	IPAddress string `json:"ip_address"`
}

func getPublicIp(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get("https://curlmyip.org")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	string_body := string(body)
	w.Header().Set("Content-Type", "application/json")
	log.Println(string_body)
	host := Host{
		IPAddress: string_body,
	}

	json.NewEncoder(w).Encode(host)
}

func main() {
	http.HandleFunc("/", getPublicIp)
	log.Printf("Server started at port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
