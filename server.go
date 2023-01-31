package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Congratulations you pass the first test! Now go to /challenge to get the next challenge.")
}

func challengeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	currentYear := currentTime.Year()
	encoded := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("/devops%d", currentYear)))
	message := fmt.Sprintf("Congratulations again! Now go to %s to start your final challenge.", encoded)
	fmt.Fprintln(w, message)
}

func finalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Congratulations you pass the last test! This is a template of a DevOps test using Docker. Project: https://github.com/giovannirossini/devops-test.")
}

func main() {
	currentTime := time.Now()
	currentYear := currentTime.Year()
	page := fmt.Sprintf("/devops%d", currentYear)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/challenge", challengeHandler)
	http.HandleFunc(page, finalHandler)

	err := http.ListenAndServe(":8573", nil)
	if err != nil {
		fmt.Println(err)
	}
}
