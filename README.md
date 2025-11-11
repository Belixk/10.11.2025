# 10.11.2025

# Сервис для проверки ссылок
Сервис для проверки доступонсти веб-ресурсов с генерацией в pdf файл

# Использованные паттерны
- **Хранилище** с сериализацией в json файл
- **Сохранение** сохранение состояния при перезагрузке
- **Разделение ответственности** handelrs, services, storage

# Структура проекта
internal/
├──handlers/ #http обработчики
├──models/ # структуры данных
├──services/ # бизнес-логика
├──storage/ # хранение данных

## Api Endpoints

### POST /check_links
Проверка доступа ссылок и возвращение номера links_num

**Запрос:curl -X POST http://localhost:8080/check_links -H "Content-Type: application/json" -d '{"links": ["google.com", "yandex.ru"]}'**
```json
{
    {"links":{"google.com":"available", "yandex.ru":"available"},"links_num":1}
}
```

**Ответ**
```json
{
    {
    "links":{
        "google.com":"available",
        "yandex.ru":"available"
    },
    "links_num":1
    }
}
```

### POST /report
Генерирует PDF отчет по ранее проверенным ссылкам

**Запрос: curl -X POST http://localhost:8080/report -H "Content-Type: application/json" -d '{"links_list": [1]}' --output report.pdf**

**Ответ**
pdf файл с ссылками

## Примеры использования

Проверка ссылок(/check_links):
```bash
curl -X POST http://localhost:8080/check_links -H "Content-Type:application/json" -d '{"links": ["", "", "", ""]}' 
```

Получение отчета в pdf
```bash
curl -X POST http://localhost:8080/report -H "Content-Type:application/json" -d '{"links_list": [1, ...., n+1]}' --output filename.pdf
```

# Запуск
```bash
go run cmd/api/main.go
```
Сервер запуститься на http://localhost:8080

# Особенности
- Сохранение в storage.json при перезагрузке
- Автоматическое добавление https:// к ссылкам, которые без него
- Потокобезопасные операции с sync.Mutex
- Генерация pdf файла с помощью go-pdf/fpdf