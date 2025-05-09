package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"time"
)

type ViaCEP struct {
	Localidade string `json:"localidade"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

type Temperature struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func main() {
	http.HandleFunc("/weather", weatherHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if !isValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	city, err := fetchCityByCEP(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	tempC, err := fetchTemperatureByCity(city)
	if err != nil {
		http.Error(w, "failed to fetch temperature", http.StatusInternalServerError)
		return
	}

	res := Temperature{
		TempC: round(tempC, 1),
		TempF: round(tempC*1.8+32, 1),
		TempK: round(tempC+273, 1),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func isValidCEP(cep string) bool {
	matched, _ := regexp.MatchString(`^[0-9]{8}$`, cep)
	return matched
}

func fetchCityByCEP(cep string) (string, error) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil || resp.StatusCode != 200 {
		return "", fmt.Errorf("CEP not found")
	}
	defer resp.Body.Close()
	var viaCEP ViaCEP
	if err := json.NewDecoder(resp.Body).Decode(&viaCEP); err != nil || viaCEP.Localidade == "" {
		return "", fmt.Errorf("invalid response")
	}
	return viaCEP.Localidade, nil
}

func fetchTemperatureByCity(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return 0, fmt.Errorf("API key not set")
	}
	client := http.Client{Timeout: 1 * time.Second}
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return 0, fmt.Errorf("weather fetch failed")
	}
	defer resp.Body.Close()
	var weather WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return 0, err
	}
	return weather.Current.TempC, nil
}

func round(val float64, places int) float64 {
	factor := math.Pow(10, float64(places))
	return math.Round(val*factor) / factor
}
