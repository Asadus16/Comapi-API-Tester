package config

import (
	"fmt"
	"os"

	"github.com/Asadus16/comapi/pkg/types"
	"gopkg.in/yaml.v2"
)

// LoadTestSuite reads and parses a YAML test configuration file
func LoadTestSuite(filename string) (*types.TestSuite, error) {
	// Read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	// Parse YAML into our TestSuite struct
	var suite types.TestSuite
	err = yaml.Unmarshal(data, &suite)
	if err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	// Basic validation
	if suite.Name == "" {
		return nil, fmt.Errorf("test suite name is required")
	}
	
	if suite.BaseURL == "" {
		return nil, fmt.Errorf("base_url is required")
	}
	
	if len(suite.Tests) == 0 {
		return nil, fmt.Errorf("at least one test is required")
	}

	// Validate each test case
	for i, test := range suite.Tests {
		if test.Name == "" {
			return nil, fmt.Errorf("test %d: name is required", i+1)
		}
		if test.Method == "" {
			return nil, fmt.Errorf("test '%s': method is required", test.Name)
		}
		if test.Path == "" {
			return nil, fmt.Errorf("test '%s': path is required", test.Name)
		}
		if len(test.Assertions) == 0 {
			return nil, fmt.Errorf("test '%s': at least one assertion is required", test.Name)
		}
	}

	return &suite, nil
}

// ValidateAssertion checks if an assertion is properly formatted
func ValidateAssertion(assertion types.Assertion) error {
	switch assertion.Type {
	case "status":
		if assertion.Expected == nil {
			return fmt.Errorf("status assertion requires 'expected' field")
		}
	case "json_path":
		if assertion.Target == "" {
			return fmt.Errorf("json_path assertion requires 'target' field")
		}
		if assertion.Expected == nil {
			return fmt.Errorf("json_path assertion requires 'expected' field")
		}
	case "header":
		if assertion.Target == "" {
			return fmt.Errorf("header assertion requires 'target' field (header name)")
		}
		if assertion.Expected == nil {
			return fmt.Errorf("header assertion requires 'expected' field")
		}
	case "response_time":
		if assertion.Expected == nil {
			return fmt.Errorf("response_time assertion requires 'expected' field")
		}
		if assertion.Operator == "" {
			assertion.Operator = "less_than" // default
		}
	default:
		return fmt.Errorf("unsupported assertion type: %s", assertion.Type)
	}
	
	return nil
}