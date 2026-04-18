package stockapi

import (
	"autoJoosik-mcp-server/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Result struct {
	Data    string
	Mapping string
}

type response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Mapping interface{} `json:"mapping"`
}

type BuySellResult struct {
	StkCd   string
	Qty     string
	Message string
	Error   string
}

type BuySellResponse struct {
	Status string `json:"status"`
	StkCd  string `json:"stkCd"`
	Qty    string `json:"qty"`
	Error  string `json:"error"`
}

func CallStockInfoAPI() (Result, error) {
	resp, err := http.Get("http://localhost:6070/market/stockInfos")
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	logger.Info("service:CallStockAPI")

	var apiResp response
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return Result{}, err
	}

	dataBytes, err := json.Marshal(apiResp.Data)
	if err != nil {
		return Result{}, err
	}

	mappingBytes, err := json.Marshal(apiResp.Mapping)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Data:    string(dataBytes),
		Mapping: string(mappingBytes),
	}, nil
}

func CallStockBuyAPI(stkCd string, qty int) (BuySellResult, error) {
	// 1. 요청 body 생성(POST)
	reqBody := map[string]interface{}{
		"StkCd": stkCd,
		"Qty":   qty,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return BuySellResult{
			Message: "매수 완료",
			StkCd:   stkCd,
			Qty:     fmt.Sprint(qty),
			Error:   err.Error(),
		}, err
	}

	// 2. 요청 생성
	req, err := http.NewRequest("POST", "http://localhost:6070/market/buy", bytes.NewBuffer(jsonData))
	if err != nil {
		return BuySellResult{
			Message: "매수 완료",
			StkCd:   stkCd,
			Qty:     fmt.Sprint(qty),
			Error:   err.Error(),
		}, err
	}

	// 3. 헤더 설정
	client := &http.Client{}

	// 4. 요청 실행
	resp, err := client.Do(req)
	if err != nil {
		return BuySellResult{
			Message: "매수 실패",
			StkCd:   stkCd,
			Qty:     fmt.Sprint(qty),
			Error:   err.Error(),
		}, err
	}
	defer resp.Body.Close()

	// 5. 응답 처리
	var apiResp BuySellResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return BuySellResult{}, err
	}

	return BuySellResult{
		Message: "매수 완료",
		StkCd:   apiResp.StkCd,
		Qty:     apiResp.Qty,
		Error:   apiResp.Error,
	}, nil
}

func CallStockSellAPI(stkCd string, qty int) (BuySellResult, error) {
	// 1. 요청 body 생성(POST)
	reqBody := map[string]interface{}{
		"StkCd": stkCd,
		"Qty":   qty,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return BuySellResult{
			Message: "매매 완료",
			StkCd:   stkCd,
			Qty:     fmt.Sprint(qty),
			Error:   err.Error(),
		}, err
	}

	// 2. 요청 생성
	req, err := http.NewRequest("POST", "http://localhost:6070/market/sell", bytes.NewBuffer(jsonData))
	if err != nil {
		return BuySellResult{
			Message: "매매 완료",
			StkCd:   stkCd,
			Qty:     fmt.Sprint(qty),
			Error:   err.Error(),
		}, err
	}

	// 3. 헤더 설정
	client := &http.Client{}

	// 4. 요청 실행
	resp, err := client.Do(req)
	if err != nil {
		return BuySellResult{
			Message: "매매 실패",
			StkCd:   stkCd,
			Qty:     fmt.Sprint(qty),
			Error:   err.Error(),
		}, err
	}
	defer resp.Body.Close()

	// 5. 응답 처리
	var apiResp BuySellResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return BuySellResult{}, err
	}

	return BuySellResult{
		Message: "매매 완료",
		StkCd:   apiResp.StkCd,
		Qty:     apiResp.Qty,
		Error:   apiResp.Error,
	}, nil
}
