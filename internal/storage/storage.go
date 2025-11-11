package storage

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/Belixk/10.11.2025/internal/models"
)

var (
	mu        sync.Mutex
	allResult = make(map[int]*models.CheckLinksResponse)
	counter   int
)

func Save(rs *models.CheckLinksResponse) error {
	allResult[rs.LinksNum] = rs
	return nil
}

func GetByNumber(num int) *models.CheckLinksResponse {
	return allResult[num]
}

func SaveToFile() error {
	mu.Lock()
	defer mu.Unlock()

	data, err := json.Marshal(allResult)
	if err != nil {
		return err
	}

	return os.WriteFile("storage.json", data, 0644)
}

func GetLinksNum() int {
	mu.Lock()
	defer mu.Unlock()
	counter++
	return counter
}

func LoadFromFile() error {
	if _, err := os.Stat("storage.json"); err != nil {
		return nil
	}
	data, err := os.ReadFile("storage.json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &allResult); err != nil {
		return err
	}
	for num := range allResult {
		if num > counter {
			counter = num
		}
	}
	return nil
}
