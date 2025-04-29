package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchAPI(endpoint string) ([]byte, error) {
	client := &http.Client{Timeout: timeout}
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("request creation failed: %w", err)
	}

	req.Header.Set("User-Agent", "eol-cli/1.0")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	return io.ReadAll(resp.Body)
}

func ListProducts() ([]string, error) {
	body, err := fetchAPI("/all.json")
	if err != nil {
		return nil, err
	}

	var products []string
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, fmt.Errorf("JSON parsing failed: %w", err)
	}

	return products, nil
}

func GetProductCycles(product string) ([]Cycle, error) {
	body, err := fetchAPI(fmt.Sprintf("/%s.json", product))
	if err != nil {
		return nil, err
	}

	var cycles []Cycle
	if err := json.Unmarshal(body, &cycles); err != nil {
		return nil, fmt.Errorf("cycles parsing failed: %w", err)
	}

	return cycles, nil
}

func GetCycleDetails(product, cycle string) (*Cycle, error) {
	body, err := fetchAPI(fmt.Sprintf("/%s/%s.json", product, cycle))
	if err != nil {
		return nil, err
	}

	var details Cycle
	if err := json.Unmarshal(body, &details); err != nil {
		return nil, fmt.Errorf("cycle details parsing failed: %w", err)
	}

	return &details, nil
}
