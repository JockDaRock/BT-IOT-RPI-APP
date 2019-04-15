package main

import (
	"fmt"
	"encoding/json"
	//"time"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	//"bytes"
	//"os"
	//"io"
	"os/exec"
	"io/ioutil"
	//"archive/tar"
	//"path/filepath"
	//"./commandModels"
)

func startSSH () {
	out, err := exec.Command("gotty", "-w", "/bin/sh").Output()

	if err != nil {
		log.Fatal(err)
	}

	output := string(out)

	fmt.Printf(output)
}

func getSSH(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	go startSSH()
	fmt.Fprint(w, "SSH started at Device on Port 8080 in your web browser.")
}

func main() {
	router := httprouter.New()
	router.GET("/getSSH", getSSH)
	log.Fatal(http.ListenAndServe(":5000", router))
}