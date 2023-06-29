package service

import "time"

type (
	ConverterResp struct {
		Meta struct {
			LastUpdatedAt time.Time `json:"last_updated_at"`
		} `json:"meta"`
		Data map[string]CurencyData `json:"data"`
	}

	CurencyData struct {
		Code  string  `json:"code"`
		Value float64 `json:"value"`
	}
)
