package retailshop

import (
	"bytes"
	"encoding/json"
)

type Trow map[string]interface{}

type Transaction struct {
	Id           string `json:"id"`
	StartTime    uint64 `json:"startTime"`
	CompleteTime uint64 `json:"completeTime"`
	CashRecieved uint64 `json:"cashRecieved"`
	TRowStack    []Trow `json:"tRowsStack"`
}

func DecodeTransactionJson(jsonData []byte) (*Transaction, error) {
	var t Transaction
	d := json.NewDecoder(bytes.NewBuffer(jsonData))
	return &t, d.Decode(&t)
}



