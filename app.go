package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

const port = 8080

func indexHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	url := os.Getenv("DISCORD_WEBHOOK_URL")
	botname := os.Getenv("DISCORD_WEBHOOK_BOTNAME")

	message := `{"username": "` + botname + `", "content": "` + jsonEscape(string(reqDump)) + `"}`
	fmt.Println("message: ", message)
	// + string(reqDump) +jsonEscape(string(reqDump))

	var jsonStr = []byte(string(message))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	fmt.Printf("REQUEST:\n%s", string(reqDump))
	w.Write([]byte("Hello World\n"))
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	s := string(b)
	return s[1 : len(s)-1]
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Printf("Starting HTTP server at port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
