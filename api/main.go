package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/graphql-go/graphql"
	"log"
	"net/http"
	"os"
	"todoapi/models"
	"todoapi/schema"
)

type Payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func handleRequest(res http.ResponseWriter, req *http.Request) {
	payload := Payload{}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	switch req.Method {
	case "POST":
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&payload)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
	case "GET":
		payload.Query = req.URL.Query().Get("query")
		variables := map[string]interface{}{}
		err := json.Unmarshal([]byte(req.URL.Query().Get("variables")), &variables)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		payload.Variables = variables
	}

	debug, _ := json.Marshal(payload)
	log.Println(string(debug))

	result := graphql.Do(graphql.Params{
		Schema:         schema.Schema,
		RequestString:  payload.Query,
		VariableValues: payload.Variables,
	})
	err := json.NewEncoder(res).Encode(result)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
}

func main() {

	dsn := flag.String("dsn", os.Getenv("DSN"), "connection data source name")
	bind := flag.String("bind", os.Getenv("BIND"), "http server binding")
	flag.Parse()
	if len(*dsn) == 0 {
		log.Fatal("DNS was not provided")
	}
	if len(*bind) == 0 {
		log.Fatal("Bind information was not provided")
	}
	models.InitDatabase(*dsn)

	http.HandleFunc("/", handleRequest)
	fmt.Println("Listening on ", *bind)
	log.Fatal(http.ListenAndServe(*bind, nil))
}
