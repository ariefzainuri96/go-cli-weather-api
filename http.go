package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	// "github.com/redis/go-redis/v9"
)

func TestFunction() {
	fmt.Println("test function")
}

func SetupHttp() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("GET /weather", getWeather)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to go-cli-weather-api")
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	redis := NewRedisService()

	address := r.URL.Query().Get("address")

	cacheWeather, err := GetRedisValue[[]byte](redis.Client, address)

	if err == nil {
		fmt.Println("cache found")
		w.WriteHeader(http.StatusOK)
		w.Write(cacheWeather)

		return
	}

	err = godotenv.Load()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	currDate := time.Now().Local().Format("2006-01-02")

	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s/%s?key=%s", address, currDate, weatherAPIKey)

	resp, err := http.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = SetRedisValue(redis.Client, address, body, time.Hour*1)

	if err != nil {
		// log error to analytics use case
		fmt.Println("err save redis value: ", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
