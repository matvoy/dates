package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/matvoy/dates/models"
)

func (a *App) registerRoutes() {
	a.router.Use(middleware.Logger())
	a.router.GET("/date", a.handler)
}

func (a *App) handler(c echo.Context) error {
	query := new(models.QueryDate)
	if err := c.Bind(query); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	resDate, err := getDateFromAPI(query.Month, query.Day)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	date := mapToDomain(resDate)
	if err := a.processDate(date); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, date)
}

func getDateFromAPI(month, day int) (*models.DateResponseBody, error) {
	if month == 0 || day == 0 {
		now := time.Now()
		day = now.Day()
		month = int(now.Month())
	}
	resp, err := http.Get(fmt.Sprintf("https://history.muffinlabs.com/date/%v/%v", month, day))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	date := new(models.DateResponseBody)
	if err := json.Unmarshal(body, date); err != nil {
		return nil, err
	}
	return date, nil
}

func mapToDomain(date *models.DateResponseBody) *models.Date {
	return &models.Date{
		URL:        date.URL,
		DateString: date.Date,
		Events:     date.Data.Events,
		Births:     date.Data.Births,
		Deaths:     date.Data.Deaths,
	}
}
