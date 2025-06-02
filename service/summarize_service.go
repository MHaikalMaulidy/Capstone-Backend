package service

import (
	"backend-summarizer/database"
	"backend-summarizer/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ProcessSummary(text string) string {
	summary := callMLModel(text)
	database.SaveSummary(text, summary)
	return summary
}

func callMLModel(text string) string {
	payload := map[string]interface{}{
		"text": text,
	}
	jsonVal, _ := json.Marshal(payload)

	resp, err := http.Post("https://api-capstone-kappa.vercel.app/summarize", "application/json", bytes.NewBuffer(jsonVal))
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
