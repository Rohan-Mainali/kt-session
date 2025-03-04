package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var timeOut string
	http.HandleFunc("/test", test)
	http.HandleFunc("/read-config", readConfig)

	flag.StringVar(&timeOut, "time-out", "10s", "Set the timeout duration")
	flag.Parse()
	fmt.Println("time-out", timeOut)
	fmt.Println("ENV_TYPE", os.Getenv("ENV"))
	fmt.Println("API_TOKEN", os.Getenv("API_TOKEN"))
	fmt.Println("Server is running on port 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hey there!, welcome to your first k8s session"))
}

func readConfig(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)

	token := os.Getenv("TOKEN")

	w.Write([]byte(token))
}
