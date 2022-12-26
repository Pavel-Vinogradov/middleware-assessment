package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

// main Функция main
func main() {

	// Echo instance
	s := echo.New()
	s.Use(MW)
	s.GET("/day", Handler)
	err := s.Start(":8080")
	//Логируем ошибку
	if err != nil {
		log.Fatal(err)
	}

}

// Handler  Обработчик запроса
func Handler(context echo.Context) error {
	d := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	during := time.Until(d)
	s := fmt.Sprintf("Количество дней: %d", int64(during.Hours()/24))
	err := context.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}

// MW Промежуточный обработчик для запроса
func MW(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		val := c.Request().Header.Get("User-Role")
		if val != "admin" {
			return c.String(http.StatusForbidden, "У вас нет прав")
		}
		err := next(c)
		if err != nil {
			return err
		}
		return nil
	}
}
