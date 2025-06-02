package types

import (
	"time"
)

// TestSuite represents a collection of API tests
type TestSuite struct {
	Name        string            `json:"name" yaml:"name"`
	BaseURL     string            `json:"base_url" yaml:"base_url"`
	Headers     map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Environment map[string]string `json:"environment,omitempty" yaml:"environment,omitempty"`
	Tests       []TestCase        `json:"tests" yaml:"tests"`
}

// TestCase represents a single API test
type TestCase struct {
	Name        string            `json:"name" yaml:"name"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Method      string            `json:"method" yaml:"method"`
	Path        string            `json:"path" yaml:"path"`           // For backward compatibility
	URL         string            `json:"url" yaml:"url"`             // New: complete URL
	Headers     map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Body        string            `json:"body,omitempty" yaml:"body,omitempty"`
	Assertions  []Assertion       `json:"assertions" yaml:"assertions"`
}

// Assertion represents a test assertion
type Assertion struct {
	Type     string      `json:"type" yaml:"type"`         // "status", "header", "json_path", "response_time"
	Target   string      `json:"target,omitempty" yaml:"target,omitempty"`   // JSON path, header name, etc.
	Expected interface{} `json:"expected" yaml:"expected"` // Expected value
	Operator string      `json:"operator,omitempty" yaml:"operator,omitempty"` // "equals", "contains", "less_than", etc.
}

// TestResult represents the result of a single test
type TestResult struct {
	TestName     string              `json:"test_name"`
	Status       TestStatus          `json:"status"`
	Duration     time.Duration       `json:"duration"`
	Request      RequestInfo         `json:"request"`
	Response     ResponseInfo        `json:"response"`
	Assertions   []AssertionResult   `json:"assertions"`
	Error        string              `json:"error,omitempty"`
}

// TestStatus represents the status of a test
type TestStatus string

const (
	StatusPass TestStatus = "PASS"
	StatusFail TestStatus = "FAIL"
	StatusSkip TestStatus = "SKIP"
)

// RequestInfo contains information about the HTTP request
type RequestInfo struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body,omitempty"`
}

// ResponseInfo contains information about the HTTP response
type ResponseInfo struct {
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Size       int64             `json:"size"`
}

// AssertionResult represents the result of a single assertion
type AssertionResult struct {
	Type     string      `json:"type"`
	Target   string      `json:"target,omitempty"`
	Expected interface{} `json:"expected"`
	Actual   interface{} `json:"actual"`
	Passed   bool        `json:"passed"`
	Message  string      `json:"message,omitempty"`
}

// SuiteResult represents the overall result of a test suite
type SuiteResult struct {
	SuiteName    string        `json:"suite_name"`
	TotalTests   int           `json:"total_tests"`
	PassedTests  int           `json:"passed_tests"`
	FailedTests  int           `json:"failed_tests"`
	SkippedTests int           `json:"skipped_tests"`
	Duration     time.Duration `json:"duration"`
	Results      []TestResult  `json:"results"`
}

// Config represents the application configuration
type Config struct {
	Timeout         time.Duration `json:"timeout" yaml:"timeout"`
	MaxRedirects    int           `json:"max_redirects" yaml:"max_redirects"`
	FollowRedirects bool          `json:"follow_redirects" yaml:"follow_redirects"`
	VerifySSL       bool          `json:"verify_ssl" yaml:"verify_ssl"`
	OutputFormat    string        `json:"output_format" yaml:"output_format"` // "console", "json", "html"
	OutputFile      string        `json:"output_file,omitempty" yaml:"output_file,omitempty"`
}