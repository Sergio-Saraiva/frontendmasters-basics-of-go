package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"frontendmasters.com/courses/go-basics/cryptomasters/models"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*models.Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("currency must be 3 characters long")
	}

	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var cexResponse CEXResponse
	err = json.Unmarshal(bodyBytes, &cexResponse)
	if err != nil {
		return nil, err
	}

	rate := &models.Rate{
		Currency: upCurrency,
		Price:    cexResponse.Ask,
	}

	return rate, nil
}
