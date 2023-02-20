package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode"
)

type StatusHandler struct {
	status bool
	host   string
	ports  []string
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func ports_check(c rune) bool {
	return !unicode.IsNumber(c)
}

func (sh *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sh.status = raw_connect(sh.host, sh.ports)
	if sh.status {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "OK")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "NOT OK")
	}
}

func raw_connect(host string, ports []string) bool {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			log.Println("Error connecting to: ", host, port)
			return false
		}
		if conn != nil {
			defer conn.Close()
		}
	}
	return true
}

func main() {

	sh := &StatusHandler{status: false, host: os.Getenv("HOST"), ports: strings.FieldsFunc(os.Getenv("PORTS"), ports_check)}
	http.Handle("/", sh)
	log.Fatal(http.ListenAndServe(getenv("LISTEN", "")+":"+getenv("PORT", "8080"), nil))
}
