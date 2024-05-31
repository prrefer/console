package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

var setConsoleTitleW = syscall.NewLazyDLL("kernel32").NewProc("SetConsoleTitleW")

func setConsoleName[T string | []byte](name T) {
	utf16String := utf16.Encode([]rune(string(name)))
	setConsoleTitleW.Call(uintptr(unsafe.Pointer(&utf16String[0])))
}

func main() {
	fmt.Println("running")
	setConsoleName("https://github.com/prrefer/console")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /print", func(w http.ResponseWriter, r *http.Request) {
		if text, err := io.ReadAll(r.Body); err == nil {
			fmt.Print(string(text))
		}
	})
	mux.HandleFunc("POST /warn", func(w http.ResponseWriter, r *http.Request) {
		if text, err := io.ReadAll(r.Body); err == nil {
			fmt.Printf("[\u001B[33m*\u001B[0m] %s", string(text))
		}
	})
	mux.HandleFunc("POST /error", func(w http.ResponseWriter, r *http.Request) {
		if text, err := io.ReadAll(r.Body); err == nil {
			fmt.Printf("[\u001B[31m*\u001B[0m] %s", string(text))
		}
	})
	mux.HandleFunc("POST /clear", func(w http.ResponseWriter, r *http.Request) {
		clear := exec.Command("cmd", "/c", "cls")
		clear.Stdout = os.Stdout
		clear.Run()
	})
	mux.HandleFunc("POST /input", func(w http.ResponseWriter, r *http.Request) {
		if prompt, err := io.ReadAll(r.Body); err == nil {
			fmt.Print(string(prompt))
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				answer := scanner.Text()
				w.Write([]byte(answer))
			}
		}
	})
	mux.HandleFunc("POST /title", func(w http.ResponseWriter, r *http.Request) {
		if title, err := io.ReadAll(r.Body); err == nil {
			setConsoleName(title)
		}
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatal(err)
	}
}
