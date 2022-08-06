package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

const port = 8080

func indexHandler(w http.ResponseWriter, r *http.Request) {
	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}

	// url := "https://discordapp.com/api/webhooks/1005315978202198137/6uKnusikcv7__bDSp2BN0O8hosATuI7kYNlm1g4obKqqnkPQnM-grLPkRkVQlUNspQKb"

	// message := `{"username": "test", "content": "` + string(reqDump) + `"}`
	// // + string(reqDump) +

	// var jsonStr = []byte(string(message))
	// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	//
	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	fmt.Println("raw req dump:", string(reqDump))

	fmt.Printf("REQUEST:\n%s", string(reqDump))
	w.Write([]byte("Hello World\n"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Printf("Starting HTTP server at port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
