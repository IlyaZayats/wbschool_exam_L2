package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"task11/internal/handlers"
	"task11/internal/repositories"
	"task11/internal/services"
	"task11/internal/storage"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	repo, err := repositories.NewCalendarRepository(storage.NewCalendarStorage(&sync.RWMutex{}))
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	service, err := services.NewCalendarService(repo)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	handler, err := handlers.NewCalendarHandlers(http.NewServeMux(), service)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	if err := handler.Run("localhost:8080"); err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	//mux := http.NewServeMux()
	//rh := http.RedirectHandler("http://example.org", 307)
	//mux.Handle("/foo", rh)
	//
	//if err := http.ListenAndServe("localhost:8080", mux); err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}

	//s, err := server.NewServer("8080", handler.mux)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}
	//if err := s.Run(); err != nil {
	//	fmt.Println(err.Error())
	//	os.Exit(1)
	//}
}
