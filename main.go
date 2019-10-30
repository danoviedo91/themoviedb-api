package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	type Genres struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	type Data struct {
		Genres []Genres `json:"genres"`
	}

	APIURL := "https://api.themoviedb.org/3/genre/movie/list?language=en-US&api_key=" + os.Getenv("API_KEY")
	fmt.Println(APIURL)

	req, errNewRequest := http.NewRequest(http.MethodGet, APIURL, nil)
	if errNewRequest != nil {
		panic(errNewRequest)
	}

	client := http.DefaultClient
	resp, errDo := client.Do(req)
	if errDo != nil {
		panic(errDo)
	}

	defer resp.Body.Close()

	body, errReadAll := ioutil.ReadAll(resp.Body)
	if errReadAll != nil {
		panic(errReadAll)
	}

	var data Data
	errUnmarshal := json.Unmarshal(body, &data)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	for _, genre := range data.Genres {
		fmt.Println("----------------------------")
		fmt.Println("ID of Genre is:", genre.ID)
		fmt.Println("Name of Genre is:", genre.Name)
	}

}
