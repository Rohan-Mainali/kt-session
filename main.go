package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/read-config", readConfig)
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

	filePath := "/config/ca.crt"

	file, _ := os.Open(filePath)
	readFile, _ := io.ReadAll(file)

	w.Write(readFile)
}
