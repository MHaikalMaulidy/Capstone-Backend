package service

import (
	"backend-summarizer/database"
	"backend-summarizer/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ProcessSummary(text string, n int) string {
	summary := callMLModel(text, n)
	database.SaveSummary(text, summary)
	return summary
}

func callMLModel(text string, n int) string {
	payload := map[string]interface{}{
		"text": text,
		"n":    n,
	}
	jsonVal, _ := json.Marshal(payload)

	resp, err := http.Post("https://web-production-6ed8.up.railway.app/summarize", "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		return "Error contacting summarization service"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)
	return result["summary"]
}

func FetchSummaries() []model.Summary {
	return database.GetAllSummaries()
}
