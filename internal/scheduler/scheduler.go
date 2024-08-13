package scheduler

import (
	"log"
	"time"

	"github.com/mohrezfadaei/ipresist/internal/db"
	"github.com/mohrezfadaei/ipresist/resource/apiv1/controllers"
)

func StartHealthCheckScheduler(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			performHealthChecks()
		}
	}()
}

func performHealthChecks() {
	var ips []db.IP
	if err := db.DB.Find(&ips).Error; err != nil {
		log.Printf("Failed to fetch IPs for healt check: %v", err)
		return
	}

	nodes := []string{
		"ir5.node.check-host.net",
		"ir6.node.check-host.net",
		"ir3.node.check-host.net",
		"ir1.node.check-host.net",
	}

	HealthCheckController := controllers.HealthCheckController{}
	for _, ip := range ips {
		log.Printf("Perfoming health check for IP: %s", ip.IPAddress)
		if err := HealthCheckController.FetchAndStorePingResults(ip.IPAddress, nodes); err != nil {
			log.Printf("Health check failed for IP: %s: %v", ip.IPAddress, err)
		}
	}
}
