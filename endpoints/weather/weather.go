package weather

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/graniticio/granitic/v2/logging"
	"github.com/graniticio/granitic/v2/ws"
)

type GetLogic struct {
	Log    logging.Logger
	ApiKey string
	City   string
}

func (gl *GetLogic) Process(ctx context.Context, req *ws.Request, res *ws.Response) {
	body, err := gl.fetchWeather()

	if err != nil {
		gl.Log.LogErrorf("Error fetching data : Error %v", err)
		res.HTTPStatus = 500
		return
	}
	var values map[string]interface{}
	err = json.Unmarshal(body, &values)
	if err != nil {
		gl.Log.LogErrorf("Error unmarshalling response body data : Error %v", err)
		res.HTTPStatus = 500
		return
	}
	res.Body = values
}

func (gl GetLogic) fetchWeather() ([]byte, error) {
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + gl.City + "&APPID=" + gl.ApiKey
	resp, err := http.Get(url)

	if err != nil {
		gl.Log.LogErrorf("Failed make http request, Err: %+v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		gl.Log.LogErrorf("Error: http status code %d", resp.StatusCode)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading body : Error %v", err)
	}
	return body, err
}
