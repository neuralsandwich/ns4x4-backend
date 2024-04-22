package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/halfsystems/ns4x4-backend/motorpart"
	"github.com/halfsystems/ns4x4-backend/inmemory"
)

const (
	dbFileName= "db"
)

func readFromBackup(filePath string) ([]motorpart.Part, error) {
	log.Println("Loading backup")
	_, err := os.Stat(filePath)

	if err != nil {
		_, err := os.Create(filePath)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	jsonData, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	var parts []motorpart.Part
	if len(jsonData) == 0 {
		log.Println("No parts found")
		return parts, nil
	}

	err = json.Unmarshal(jsonData, &parts)
	if err != nil {
		return nil, err
	}

	return parts, nil
}

func writeToBackup(filePath string, parts []motorpart.Part) error {
	jsonData, err := json.Marshal(parts)
	if err != nil {
		return err
	}

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var partRepository motorpart.Repository
	{
		partRepository = inmemory.New()
	}

	parts, err := readFromBackup(dbFileName)
	log.Println("Loaded ", len(parts), "parts")
	if err != nil {
		fmt.Print(err)
		return
	}

	for _, part := range parts {
		partRepository.AddPart(part)
	}

	var partService motorpart.Service
	{
		partService = motorpart.New(partRepository)
	}

	var loggingService motorpart.Service
	{
		loggingService = motorpart.NewLoggingService(partService)
	}

	var h http.Handler
	{
		h = motorpart.MakeHTTPHandler(loggingService)
	}

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening")
		errs <- http.ListenAndServe(
			"0.0.0.0:8080",
			h,
		)
	}()

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("terminated: ", <-errs)

    parts, err = partRepository.ListParts()
	if err != nil {
	    log.Fatal("Failed to get parts")
	}
	err = writeToBackup(dbFileName, parts)
	if err != nil {
		log.Fatal("Failed to write database")
	}
	log.Println("Wrote", len(parts), "parts")
}

