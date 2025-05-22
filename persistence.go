
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const dataFile = "quotes.json"

func LoadQuotes() {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		log.Fatalf("Ошибка чтения файла данных: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Ошибка чтения данных: %v", err)
	}

	err = json.Unmarshal(data, &quotes)
	if err != nil {
		log.Fatalf("Ошибка разбора JSON: %v", err)
	}

	for _, q := range quotes {
		if q.ID >= nextID {
			nextID = q.ID + 1
		}
	}
}

func SaveQuotes() {
	data, err := json.MarshalIndent(quotes, "", "  ")
	if err != nil {
		log.Printf("Ошибка сериализации: %v", err)
		return
	}
	err = ioutil.WriteFile(dataFile, data, 0644)
	if err != nil {
		log.Printf("Ошибка записи в файл: %v", err)
	}
}
