package data

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"math/rand"
"time"
	"github.com/hashicorp/go-hclog"
)

type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

func NewRates(log hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: log, rates: map[string]float64{}}
	err := er.getRates()
	return er, err
}

func (e *ExchangeRates) GetRate(base, destination string) (float64, error) {
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}

	dr, ok := e.rates[destination]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", destination)
	}
	return dr / br, nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes

func(e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{}{
	ret := make (chan struct{})

	go func (){
		ticker := time.NewTicker(interval)
		for {
			    select {
				case <- ticker.C:
					for k , v := range e.rates {
						
						change := (rand.Float64() / 10)

						direction := rand.Intn(1)

						if direction == 0 {
							change = 1 - change
						} else {
							change = 1  + change
						}

						e.rates[k] = v * change


					}
					ret <-struct{}{}
				}
		}

	}()

return ret

}

func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {

	}
	defer resp.Body.Close()

	md := &Cubes{}

	err = xml.NewDecoder(resp.Body).Decode(&md)
	if err != nil {
		e.log.Info("Error unmarshaling xml ")
		fmt.Println("Error unmarshaling xml ")
		return err
	}
	e.log.Info("Rates", resp.Body)

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}
		e.rates[c.Currency] = r
		e.log.Info("Rates ", c)
	}
	e.rates["EUR"] = 1
	return nil
}

type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
