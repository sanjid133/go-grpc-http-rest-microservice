package main

import (
	"flag"
	"time"
	"net/http"
	"strings"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
)

func main()  {
	address := flag.String("server", "http://localhost:9090", "http url")
	flag.Parse()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)

	var body string

	resp, err := http.Post(*address+"/v1/todo", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"api": "v1",
			"toDo": {
                "title":"title (%s)",
				"description":"description",
				"reminder":"%s"
 			}

		}
	`, pfx, pfx)))
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		body = fmt.Sprintf("%v", err)
	} else {
		body = string(bodyBytes)
	}

	log.Println(body)

	var created struct{

		Api string `json:"api"`
		Id string `json:"id"`
	}

	err = json.Unmarshal(bodyBytes, &created)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err = http.Get(fmt.Sprintf("%s%s/%s", *address, "/v1/todo", created.Id))
	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		body = fmt.Sprintf("%v", err)
	} else {
		body = string(bodyBytes)
	}

	log.Println(body)
}
