package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	welcomeMessage := "Welcome to the DevOps challenge! This is your opportunity to showcase your expertise in automation, networking and security. Whether you are an experienced DevOps professional or just starting out, this challenge will test your ability to work in a fast-paced and dynamic environment. Your task will be to demonstrate your skills in the areas of infrastructure management, deployment and automation, network security, and system performance. So, are you ready to take on this challenge and prove that you're a true DevOps professional?"
	fmt.Fprintln(w, welcomeMessage)
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
	go http.ListenAndServe(":8080", nil)

	http.HandleFunc("/challenge", challengeHandler)
	http.HandleFunc(page, finalHandler)
	http.ListenAndServe(":8573", nil)
}
