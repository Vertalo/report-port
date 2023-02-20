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
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func checkenv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("Required variable ", key, " is not set! Aborting...")
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
			log.Println("Error connecting to:", host, "on port:", port)
			return false
		}
		if conn != nil {
			defer conn.Close()
		}
	}
	return true
}

func main() {

	sh := &StatusHandler{status: false, host: checkenv("CHECKHOST"), ports: strings.FieldsFunc(checkenv("PORTS"), ports_check)}
	http.Handle("/", sh)
	log.Fatal(http.ListenAndServe(getenv("LISTEN", "")+":"+getenv("PORT", "8080"), nil))
}
