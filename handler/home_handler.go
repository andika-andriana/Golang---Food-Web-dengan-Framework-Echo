package handler

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
)

var baseURL = "http://localhost:1323"

func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go

	var datax, err = ambil_data()
	var datapopuler, errx = baca_populer()
	if err != nil {

	}
	if errx != nil {

	}
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name":        "HOME",
		"msg":         "Framework Echo Worked!",
		"data":        datax,
		"datapopuler": datapopuler,
	})
}

func ambil_data() ([]menu, error) {
	var err error
	var client = &http.Client{}
	var data []menu

	request, err := http.NewRequest("GET", baseURL+"/baca_menu", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func baca_populer() ([]menu, error) {
	var err error
	var client = &http.Client{}
	var data []menu

	request, err := http.NewRequest("GET", baseURL+"/baca_populer", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}
	return data, nil
}
