package runner

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Asadus16/comapi/internal/assertion"
	"github.com/Asadus16/comapi/pkg/types"
)

// HTTPClient handles making HTTP requests for tests
type HTTPClient struct {
	client  *http.Client
	baseURL string
	headers map[string]string
}

// NewHTTPClient creates a new HTTP client for testing
func NewHTTPClient(baseURL string, defaultHeaders map[string]string) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: strings.TrimRight(baseURL, "/"),
		headers: defaultHeaders,
	}
}

// ExecuteTest runs a single test case and returns the result (legacy method)
func (h *HTTPClient) ExecuteTest(testCase types.TestCase) types.TestResult {
	startTime := time.Now()
	
	result := types.TestResult{
		TestName: testCase.Name,
		Status:   types.StatusFail, // Default to fail, change to pass if all assertions pass
		Request: types.RequestInfo{
			Method: testCase.Method,
			URL:    h.baseURL + testCase.Path,
		},
	}

	// Make the HTTP request
	resp, err := h.makeRequest(testCase)
	if err != nil {
		result.Error = fmt.Sprintf("Request failed: %v", err)
		result.Duration = time.Since(startTime)
		return result
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Error = fmt.Sprintf("Failed to read response body: %v", err)
		result.Duration = time.Since(startTime)
		return result
	}

	// Fill in response info
	result.Response = types.ResponseInfo{
		StatusCode: resp.StatusCode,
		Headers:    convertHeaders(resp.Header),
		Body:       string(bodyBytes),
		Size:       int64(len(bodyBytes)),
	}
	
	// Fill in request info
	result.Request.Headers = h.mergeHeaders(testCase.Headers)
	result.Request.Body = testCase.Body
	
	result.Duration = time.Since(startTime)

	// Run assertions to determine if test passes or fails
	assertion.CheckAssertions(testCase, &result)
	
	return result
}

// ExecuteTestWithFullURL runs a single test case with a complete URL
func (h *HTTPClient) ExecuteTestWithFullURL(testCase types.TestCase) types.TestResult {
	startTime := time.Now()
	
	result := types.TestResult{
		TestName: testCase.Name,
		Status:   types.StatusFail, // Default to fail, change to pass if all assertions pass
		Request: types.RequestInfo{
			Method: testCase.Method,
			URL:    testCase.URL, // Use the complete URL
		},
	}

	// Make the HTTP request with full URL
	resp, err := h.makeRequestWithFullURL(testCase)
	if err != nil {
		result.Error = fmt.Sprintf("Request failed: %v", err)
		result.Duration = time.Since(startTime)
		return result
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		result.Error = fmt.Sprintf("Failed to read response body: %v", err)
		result.Duration = time.Since(startTime)
		return result
	}

	// Fill in response info
	result.Response = types.ResponseInfo{
		StatusCode: resp.StatusCode,
		Headers:    convertHeaders(resp.Header),
		Body:       string(bodyBytes),
		Size:       int64(len(bodyBytes)),
	}
	
	// Fill in request info
	result.Request.Headers = h.mergeHeaders(testCase.Headers)
	result.Request.Body = testCase.Body
	
	result.Duration = time.Since(startTime)

	// Run assertions to determine if test passes or fails
	assertion.CheckAssertions(testCase, &result)
	
	return result
}

// makeRequest creates and executes the HTTP request (legacy method)
func (h *HTTPClient) makeRequest(testCase types.TestCase) (*http.Response, error) {
	url := h.baseURL + testCase.Path
	
	// Create request body
	var body io.Reader
	if testCase.Body != "" {
		body = bytes.NewBufferString(testCase.Body)
	}
	
	// Create request
	req, err := http.NewRequest(testCase.Method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Add headers (merge default headers with test-specific headers)
	headers := h.mergeHeaders(testCase.Headers)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	
	// Make the request
	return h.client.Do(req)
}

// makeRequestWithFullURL creates and executes the HTTP request using complete URL
func (h *HTTPClient) makeRequestWithFullURL(testCase types.TestCase) (*http.Response, error) {
	url := testCase.URL
	
	// Create request body
	var body io.Reader
	if testCase.Body != "" {
		body = bytes.NewBufferString(testCase.Body)
	}
	
	// Create request
	req, err := http.NewRequest(testCase.Method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	
	// Add headers (merge default headers with test-specific headers)
	headers := h.mergeHeaders(testCase.Headers)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	
	// Make the request
	return h.client.Do(req)
}

// mergeHeaders combines default headers with test-specific headers
func (h *HTTPClient) mergeHeaders(testHeaders map[string]string) map[string]string {
	merged := make(map[string]string)
	
	// Start with default headers
	for key, value := range h.headers {
		merged[key] = value
	}
	
	// Override with test-specific headers
	for key, value := range testHeaders {
		merged[key] = value
	}
	
	return merged
}

// convertHeaders converts http.Header to map[string]string
func convertHeaders(headers http.Header) map[string]string {
	result := make(map[string]string)
	for key, values := range headers {
		if len(values) > 0 {
			result[key] = values[0] // Take the first value if multiple exist
		}
	}
	return result
}