package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/acom21/pet-proj-dsandbox/config"
	"io"
	"net/http"
	"net/url"
)

type Converter struct {
	Logger config.Logger
	API    config.CurrencyRatesAPI
}

func (c *Converter) ConvertCurrency(ctx context.Context) (*ConverterResp, error) {
	const currencyRates = "https://api.currencyapi.com/v3/latest"
	currencies := []string{"USD,RUB,AMD,RUB,EUR"}

	params := url.Values{
		"apikey":     {c.API.AccessToken},
		"currencies": currencies,
	}
	uri, err := url.JoinPath(currencyRates, params.Encode())
	var rates *ConverterResp

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("create GET request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Do: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ReadAll: %w", err)
	}

	if err = json.Unmarshal(data, &rates); err != nil {
		return nil, err
	}

	return rates, nil
}
