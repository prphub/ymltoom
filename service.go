package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type Message struct {
	Currencies []Valute `yaml:"currencies"`
}
type Valute struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func main() {
	filename, _ := filepath.Abs("./currencies.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	y := Message{}

	err = yaml.Unmarshal(yamlFile, &y)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%+v\n", y)

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, y)
	})
	fmt.Println("Server is listening...")
	http.ListenAndServe("localhost:8080", nil)

}

//   currency{name="usd"} 70
