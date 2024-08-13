package db

import (
	"context"
	"log"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/mohrezfadaei/ipresist/config"
)

var (
	InfluxClient *influxdb3.Client
)

func ConnectInfluxDB() {
	var err error
	InfluxClient, err = influxdb3.New(influxdb3.ClientConfig{
		Host:     config.INFLUXDB_HOST,
		Token:    config.INFLUXDB_TOKEN,
		Database: config.INFLUXDB_BUCKET,
	})
	if err != nil {
		log.Fatalf("Failed to connect to InfluxDB: %v", err)
	}

	log.Println("Connected to InfluxDB successfully")
}

func WritePingResult(ip, node, status string, latency float64) {
	annoted := struct {
		Measurement string    `lp:"measurement"`
		IP          string    `lp:"tag,ip"`
		Node        string    `lp:"tag,node"`
		Status      string    `lp:"field,status"`
		Latency     float64   `lp:"field,latency"`
		Time        time.Time `lp:"timestamp"`
	}{
		"ping_results",
		ip,
		node,
		status,
		latency,
		time.Now(),
	}

	data := []any{annoted}
	if err := InfluxClient.WriteData(context.Background(), data); err != nil {
		log.Printf("Error writing to InfluxDB: %v", err)
	}
}
