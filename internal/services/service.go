package services

import (
	"net/http"
	"strings"

	"github.com/Belixk/10.11.2025/internal/storage"
)

func GetLinks(links []string) (map[string]string, int, error) {
	result := make(map[string]string)
	// В цикле будем проходиться по ссылкам из массива links
	for _, link := range links {
		url := link
		if !strings.HasPrefix(link, "http") { // если они буду идти без префикса https:// просто добавляем их
			url = "https://" + link
		}
		resp, err := http.Get(url) // отправляем запрос
		if err != nil {
			result[link] = "not available"
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			result[link] = "not available" // не доступен
		} else {
			result[link] = "available" // доступен
		}

	}
	linkNum := storage.GetLinksNum()
	return result, linkNum, nil
}
