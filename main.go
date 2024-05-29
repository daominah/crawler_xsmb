package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/daominah/crawler_xsmb/pkg/core"
	"github.com/daominah/crawler_xsmb/pkg/driver/source_xsktcomvn"
	"github.com/joho/godotenv"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	refreshDataIntervalStr := os.Getenv("REFRESH_DATA_INTERVAL_SECONDS")
	refreshDataIntervalInt, err := strconv.Atoi(refreshDataIntervalStr)
	if err != nil {
		log.Fatalf("bad env REFRESH_DATA_INTERVAL_SECONDS, strconv.Atoi error: %v", err)
	}
	refreshDataInterval := time.Duration(refreshDataIntervalInt) * time.Second
	log.Printf("recognized env REFRESH_DATA_INTERVAL_SECONDS: %v", refreshDataInterval)

	var dataSource core.DataSource = source_xsktcomvn.DataSourceXsktcomvn{}
	job := func() {
		freshData, err := dataSource.DownloadHTML()
		if err != nil {
			log.Printf("error sourceXsktcomvn.DownloadHTML: %v", err)
			return
		}
		result, err := dataSource.ParseToResultXSMB(string(freshData))
		if err != nil {
			log.Printf("error ParseXsktcomvn: %v", err)
			return
		}
		result.DownloadedAt = time.Now()
		log.Printf("result: %+v", result)
	}

	for {
		log.Printf("begin job")
		job()
		log.Printf("end job")
		time.Sleep(refreshDataInterval)
	}
}
