package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)

type Data struct {
	Weather []Weather `json:"weather"`
	Main Main `json:"main"`
}

type Weather struct {
	Id int `json:"id"`
	Main string `json:"main"`
	Description string `json:"description"`
	Icon string `json:"icon"`
}

type Main struct{
	Temp float32 `json:"temp"`
	FeelsLike float32 `json:"feels_like"`
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
	Pressure int `json:"pressure"`
	Humidity int `json:"humidity"` 
}

func main(){

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	key := os.Getenv("API_KEY")

	fmt.Printf("key:%s",key)
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=%s",key)

	resp,err := http.Get(url)

	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body,err := io.ReadAll(resp.Body)
	if err != nil{
		log.Fatal(err)
	}

	var data Data

	if err := json.Unmarshal(body,&data); err != nil {
		log.Fatal(err)
	}


	fmt.Println(data)
	// Ginエンジンのインスタンスを作成
	
}