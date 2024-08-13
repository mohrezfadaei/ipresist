package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mohrezfadaei/ipresist/internal/db"
)

type HealthCheckController struct{}

type PingRequestResponse struct {
	Nodes         map[string][]string `json:"nodes"`
	Ok            int                 `json:"ok"`
	PermanentLink string              `json:"permanent_link"`
	RequestID     string              `json:"request_id"`
}

type PingResultResponse map[string][][][]interface{}

func (ctrl HealthCheckController) FetchAndStorePingResults(ip string, nodes []string) error {
	requestID, err := ctrl.SendPingRequest(ip, nodes)
	if err != nil {
		return err
	}

	pingResults, err := ctrl.getPingResults(requestID)
	if err != nil {
		return err
	}

	ctrl.storePingResults(ip, pingResults)
	return nil
}

func (ctrl HealthCheckController) SendPingRequest(ip string, nodes []string) (string, error) {
	nodesParam := ""
	for _, n := range nodes {
		nodesParam += fmt.Sprintf("%s,", n)
	}

	url := fmt.Sprintf("https://check-host.net/check-ping?host=%s%s", ip, nodesParam)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create ping request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send ping request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	var pingRequestResponse PingRequestResponse
	if err := json.NewDecoder(resp.Body).Decode(&pingRequestResponse); err != nil {
		return "", fmt.Errorf("failed to decode ping request response: %v", err)
	}

	return pingRequestResponse.RequestID, nil
}

func (ctrl HealthCheckController) getPingResults(requestID string) (PingResultResponse, error) {
	url := fmt.Sprintf("https://check-host.net/check-ping/results/%s", requestID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create get ping results request: %v", err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send get ping results request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}

	var pingResultResponse PingResultResponse
	if err := json.NewDecoder(resp.Body).Decode(&pingResultResponse); err != nil {
		return nil, fmt.Errorf("failed to decode get ping results response: %v", err)
	}

	return pingResultResponse, nil
}

func (ctrl HealthCheckController) storePingResults(ip string, results PingResultResponse) {
	for node, pings := range results {
		for _, ping := range pings[0] {
			status, _ := ping[0].(string)
			latency, _ := ping[1].(float64)
			db.WritePingResult(ip, node, status, latency)
		}
	}
}
