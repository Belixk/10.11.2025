package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Belixk/10.11.2025/internal/models"
	"github.com/Belixk/10.11.2025/internal/services"
	"github.com/Belixk/10.11.2025/internal/storage"
)

// Для проверки ссылок
func HandleCheckLinks(w http.ResponseWriter, r *http.Request) {
	var request models.CheckLinksRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	linkStatus, linksNum, err := services.GetLinks(request.Links)
	if err != nil {
		fmt.Printf("ошибка при получении данных: %v", err)
		return
	}
	response := models.CheckLinksResponse{
		Links:    linkStatus,
		LinksNum: linksNum,
	}
	json.NewEncoder(w).Encode(response)

	storage.Save(&response)
}

// Для запроса со списком номеров ранее отправленных ссылок
func HandleReport(w http.ResponseWriter, r *http.Request) {
	var request models.ReportRequest
	json.NewDecoder(r.Body).Decode(&request)
	var allData []*models.CheckLinksResponse
	for _, num := range request.LinksList {
		data := storage.GetByNumber(num)
		if data != nil {
			allData = append(allData, data)
		}
	}
	pdfData := services.GeneratePDF(allData)
	w.Header().Set("Content-Type", "application/pdf")
	w.Write(pdfData)
}
