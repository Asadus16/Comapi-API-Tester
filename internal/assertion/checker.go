

package assertion

import (
	"fmt"
	"strings"

	"github.com/Asadus16/comapi/pkg/types"
	"github.com/tidwall/gjson"
)

// CheckAssertions validates all assertions for a test result
func CheckAssertions(testCase types.TestCase, result *types.TestResult) {
	var assertionResults []types.AssertionResult
	allPassed := true

	for _, assertion := range testCase.Assertions {
		assertionResult := checkSingleAssertion(assertion, result)
		assertionResults = append(assertionResults, assertionResult)
		
		if !assertionResult.Passed {
			allPassed = false
		}
	}

	result.Assertions = assertionResults
	if allPassed {
		result.Status = types.StatusPass
	} else {
		result.Status = types.StatusFail
	}
}

// checkSingleAssertion validates a single assertion
func checkSingleAssertion(assertion types.Assertion, result *types.TestResult) types.AssertionResult {
	assertionResult := types.AssertionResult{
		Type:     assertion.Type,
		Target:   assertion.Target,
		Expected: assertion.Expected,
		Passed:   false,
	}

	switch assertion.Type {
	case "status":
		assertionResult = checkStatusAssertion(assertion, result)
	case "json_path":
		assertionResult = checkJSONPathAssertion(assertion, result)
	case "header":
		assertionResult = checkHeaderAssertion(assertion, result)
	case "response_time":
		assertionResult = checkResponseTimeAssertion(assertion, result)
	default:
		assertionResult.Message = fmt.Sprintf("Unknown assertion type: %s", assertion.Type)
	}

	return assertionResult
}

// checkStatusAssertion validates HTTP status code
func checkStatusAssertion(assertion types.Assertion, result *types.TestResult) types.AssertionResult {
	expected, ok := assertion.Expected.(int)
	if !ok {
		// Try to convert from float64 (common in JSON parsing)
		if expectedFloat, ok := assertion.Expected.(float64); ok {
			expected = int(expectedFloat)
		} else {
			return types.AssertionResult{
				Type:     assertion.Type,
				Expected: assertion.Expected,
				Actual:   result.Response.StatusCode,
				Passed:   false,
				Message:  "Expected value must be an integer",
			}
		}
	}

	actual := result.Response.StatusCode
	passed := actual == expected

	return types.AssertionResult{
		Type:     assertion.Type,
		Expected: expected,
		Actual:   actual,
		Passed:   passed,
		Message:  fmt.Sprintf("Expected status %d, got %d", expected, actual),
	}
}

// checkJSONPathAssertion validates JSON path expressions
func checkJSONPathAssertion(assertion types.Assertion, result *types.TestResult) types.AssertionResult {
	jsonData := result.Response.Body
	path := assertion.Target
	
	// Remove $.  prefix if present (gjson doesn't use it)
	if strings.HasPrefix(path, "$.") {
		path = path[2:]
	}
	
	// Use gjson to extract value from JSON path
	value := gjson.Get(jsonData, path)
	
	if !value.Exists() {
		return types.AssertionResult{
			Type:     assertion.Type,
			Target:   assertion.Target,
			Expected: assertion.Expected,
			Actual:   nil,
			Passed:   false,
			Message:  fmt.Sprintf("JSON path '%s' not found", path),
		}
	}

	var actual interface{}
	switch value.Type {
	case gjson.String:
		actual = value.String()
	case gjson.Number:
		actual = value.Num
	case gjson.True, gjson.False:
		actual = value.Bool()
	default:
		actual = value.Value()
	}

	// Check the assertion based on operator
	operator := assertion.Operator
	if operator == "" {
		operator = "equals" // default
	}

	var passed bool
	var message string

	switch operator {
	case "equals":
		passed = compareValues(actual, assertion.Expected)
		message = fmt.Sprintf("Expected %v, got %v", assertion.Expected, actual)
	case "contains":
		actualStr := fmt.Sprintf("%v", actual)
		expectedStr := fmt.Sprintf("%v", assertion.Expected)
		passed = strings.Contains(actualStr, expectedStr)
		message = fmt.Sprintf("Expected '%s' to contain '%s'", actualStr, expectedStr)
	case "not_equals":
		passed = !compareValues(actual, assertion.Expected)
		message = fmt.Sprintf("Expected not %v, got %v", assertion.Expected, actual)
	case "greater_than":
		passed = compareNumeric(actual, assertion.Expected, ">")
		message = fmt.Sprintf("Expected %v > %v", actual, assertion.Expected)
	case "less_than":
		passed = compareNumeric(actual, assertion.Expected, "<")
		message = fmt.Sprintf("Expected %v < %v", actual, assertion.Expected)
	default:
		passed = false
		message = fmt.Sprintf("Unknown operator: %s", operator)
	}

	return types.AssertionResult{
		Type:     assertion.Type,
		Target:   assertion.Target,
		Expected: assertion.Expected,
		Actual:   actual,
		Passed:   passed,
		Message:  message,
	}
}

// checkHeaderAssertion validates response headers
func checkHeaderAssertion(assertion types.Assertion, result *types.TestResult) types.AssertionResult {
	headerName := assertion.Target
	expectedValue := fmt.Sprintf("%v", assertion.Expected)
	
	actualValue, exists := result.Response.Headers[headerName]
	if !exists {
		return types.AssertionResult{
			Type:     assertion.Type,
			Target:   assertion.Target,
			Expected: assertion.Expected,
			Actual:   nil,
			Passed:   false,
			Message:  fmt.Sprintf("Header '%s' not found", headerName),
		}
	}

	operator := assertion.Operator
	if operator == "" {
		operator = "equals"
	}

	var passed bool
	var message string

	switch operator {
	case "equals":
		passed = actualValue == expectedValue
		message = fmt.Sprintf("Expected header '%s' = '%s', got '%s'", headerName, expectedValue, actualValue)
	case "contains":
		passed = strings.Contains(actualValue, expectedValue)
		message = fmt.Sprintf("Expected header '%s' to contain '%s', got '%s'", headerName, expectedValue, actualValue)
	default:
		passed = false
		message = fmt.Sprintf("Unknown operator: %s", operator)
	}

	return types.AssertionResult{
		Type:     assertion.Type,
		Target:   assertion.Target,
		Expected: assertion.Expected,
		Actual:   actualValue,
		Passed:   passed,
		Message:  message,
	}
}

// checkResponseTimeAssertion validates response time
func checkResponseTimeAssertion(assertion types.Assertion, result *types.TestResult) types.AssertionResult {
	expectedMs, ok := assertion.Expected.(float64)
	if !ok {
		if expectedInt, ok := assertion.Expected.(int); ok {
			expectedMs = float64(expectedInt)
		} else {
			return types.AssertionResult{
				Type:     assertion.Type,
				Expected: assertion.Expected,
				Actual:   result.Duration.Milliseconds(),
				Passed:   false,
				Message:  "Expected value must be a number (milliseconds)",
			}
		}
	}

	actualMs := float64(result.Duration.Milliseconds())
	operator := assertion.Operator
	if operator == "" {
		operator = "less_than"
	}

	var passed bool
	var message string

	switch operator {
	case "less_than":
		passed = actualMs < expectedMs
		message = fmt.Sprintf("Expected response time < %gms, got %gms", expectedMs, actualMs)
	case "greater_than":
		passed = actualMs > expectedMs
		message = fmt.Sprintf("Expected response time > %gms, got %gms", expectedMs, actualMs)
	case "equals":
		passed = actualMs == expectedMs
		message = fmt.Sprintf("Expected response time = %gms, got %gms", expectedMs, actualMs)
	default:
		passed = false
		message = fmt.Sprintf("Unknown operator: %s", operator)
	}

	return types.AssertionResult{
		Type:     assertion.Type,
		Expected: expectedMs,
		Actual:   actualMs,
		Passed:   passed,
		Message:  message,
	}
}

// compareValues compares two values for equality, handling type conversions
func compareValues(actual, expected interface{}) bool {
	// Handle numeric comparisons with type conversion
	if actualNum, ok := actual.(float64); ok {
		if expectedNum, ok := expected.(float64); ok {
			return actualNum == expectedNum
		}
		if expectedInt, ok := expected.(int); ok {
			return actualNum == float64(expectedInt)
		}
	}
	
	if actualInt, ok := actual.(int); ok {
		if expectedNum, ok := expected.(float64); ok {
			return float64(actualInt) == expectedNum
		}
		if expectedInt, ok := expected.(int); ok {
			return actualInt == expectedInt
		}
	}

	// String comparison
	return fmt.Sprintf("%v", actual) == fmt.Sprintf("%v", expected)
}

// compareNumeric compares two numeric values
func compareNumeric(actual, expected interface{}, operator string) bool {
	actualNum, actualOk := toFloat64(actual)
	expectedNum, expectedOk := toFloat64(expected)
	
	if !actualOk || !expectedOk {
		return false
	}
	
	switch operator {
	case ">":
		return actualNum > expectedNum
	case "<":
		return actualNum < expectedNum
	case ">=":
		return actualNum >= expectedNum
	case "<=":
		return actualNum <= expectedNum
	default:
		return false
	}
}

// toFloat64 converts various numeric types to float64
func toFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case float32:
		return float64(v), true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case int32:
		return float64(v), true
	default:
		return 0, false
	}
}