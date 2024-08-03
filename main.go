package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/skratchdot/open-golang/open"
	"net/http"
	"time"
)

func main() {
	setupWhatsAppClient()

	http.HandleFunc("/", selectContactsHandler)
	http.HandleFunc("/select", selectContactsHandler)
	http.HandleFunc("/status", statusHandler)
	http.HandleFunc("/api/status-updates", statusUpdateHandler)

	go func() {
		fmt.Println("WhatsApp Status Monitor by 0xagil starting on http://localhost:8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()
	time.Sleep(2 * time.Second)

	err := open.Run("http://localhost:8080")
	if err != nil {
		fmt.Printf("Error opening browser: %s\n", err)
	}

	select {}
}
