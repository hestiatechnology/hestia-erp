package fiscal

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func ValidateSSN(ssn uint32) bool {
	return true
}

// CheckVatRequest represents the request body for checking a VAT number
type CheckVatRequest struct {
	CountryCode              string `json:"countryCode"`
	VatNumber                string `json:"vatNumber"`
	RequesterMemberStateCode string `json:"requesterMemberStateCode"`
	RequesterNumber          string `json:"requesterNumber"`
	TraderName               string `json:"traderName,omitempty"`
	TraderStreet             string `json:"traderStreet,omitempty"`
	TraderPostalCode         string `json:"traderPostalCode,omitempty"`
	TraderCity               string `json:"traderCity,omitempty"`
	TraderCompanyType        string `json:"traderCompanyType,omitempty"`
}

// CheckVatResponse represents the response from the VAT checking service
type CheckVatResponse struct {
	CountryCode       string `json:"countryCode"`
	VatNumber         string `json:"vatNumber"`
	RequestDate       string `json:"requestDate"`
	Valid             bool   `json:"valid"`
	RequestIdentifier string `json:"requestIdentifier"`
	Name              string `json:"name,omitempty"`
	Address           string `json:"address,omitempty"`
	TraderName        string `json:"traderName,omitempty"`
	TraderStreet      string `json:"traderStreet,omitempty"`
	TraderPostalCode  string `json:"traderPostalCode,omitempty"`
	TraderCity        string `json:"traderCity,omitempty"`
	TraderCompanyType string `json:"traderCompanyType,omitempty"`
}

// CheckVatNumber sends a POST request to check the validity of a VAT number
func CheckVatNumber(ctx context.Context, req CheckVatRequest) (*CheckVatResponse, error) {
	// Prepare the URL
	url := "https://your-vat-checking-api.com/check-vat-number"

	// Serialize the request body into JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	// Create a new HTTP POST request
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	httpReq.Header.Set("Content-Type", "application/json")

	// Send the request using the default HTTP client
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("received non-OK response: %v, body: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse the response body into the CheckVatResponse struct
	var vatResponse CheckVatResponse
	err = json.NewDecoder(resp.Body).Decode(&vatResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &vatResponse, nil
}
