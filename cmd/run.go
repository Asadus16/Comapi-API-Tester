// package cmd

// import (
// 	"fmt"
// 	"os"
// 	"github.com/Asadus16/comapi/internal/config"
// 	"github.com/Asadus16/comapi/internal/runner" 

// 	"github.com/Asadus16/comapi/pkg/types"
// 	"github.com/spf13/cobra"
// )

// // runCmd represents the run command
// var runCmd = &cobra.Command{
// 	Use:   "run [test-file]",
// 	Short: "Run API tests from a YAML file",
// 	Long: `Run API tests defined in a YAML configuration file.

// Example:
//   comapi run tests.yaml
//   comapi run examples/sample.yaml`,
// 	Args: cobra.ExactArgs(1),
// 	Run: func(cmd *cobra.Command, args []string) {
// 		testFile := args[0]
		
// 		// Check if file exists
// 		if _, err := os.Stat(testFile); os.IsNotExist(err) {
// 			fmt.Printf("❌ Test file not found: %s\n", testFile)
// 			os.Exit(1)
// 		}
		
// 		fmt.Printf("🧭 Running tests from: %s\n", testFile)
		
// 		// Load and parse the test configuration
// 		suite, err := config.LoadTestSuite(testFile)
// 		if err != nil {
// 			fmt.Printf("❌ Failed to load test suite: %v\n", err)
// 			os.Exit(1)
// 		}
		
// 		fmt.Printf("📋 Test Suite: %s\n", suite.Name)
// 		fmt.Printf("🌐 Base URL: %s\n", suite.BaseURL)
// 		fmt.Printf("🧪 Running %d test(s)...\n\n", len(suite.Tests))
		
// 		// Create HTTP client
// 		httpClient := runner.NewHTTPClient(suite.BaseURL, suite.Headers)
		
// 		// Run each test
// 		var results []types.TestResult
// 		for i, test := range suite.Tests {
// 			fmt.Printf("Running test %d/%d: %s\n", i+1, len(suite.Tests), test.Name)
			
// 			result := httpClient.ExecuteTest(test)
// 			results = append(results, result)
			
// 			// Show basic result
// 			if result.Status == types.StatusPass {
// 				fmt.Printf("  ✅ %s - %dms\n", result.Status, result.Duration.Milliseconds())
// 			} else {
// 				fmt.Printf("  ❌ %s - %s\n", result.Status, result.Error)
// 			}
// 		}
		
// 		fmt.Printf("\n🎯 Test Summary:\n")
// 		passed := 0
// 		for _, result := range results {
// 			if result.Status == types.StatusPass {
// 				passed++
// 			}
// 		}
// 		fmt.Printf("  ✅ Passed: %d/%d\n", passed, len(results))
// 		if passed < len(results) {
// 			fmt.Printf("  ❌ Failed: %d/%d\n", len(results)-passed, len(results))
// 		}
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(runCmd)
	
// 	// Add flags for output format, verbose mode, etc.
// 	runCmd.Flags().StringP("output", "o", "console", "Output format (console, json, html)")
// 	runCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
// 	runCmd.Flags().StringP("env", "e", "", "Environment file for variable substitution")
// }

package cmd

import (
	"fmt"
	"os"
	"github.com/Asadus16/comapi/internal/config"
	"github.com/Asadus16/comapi/internal/runner" 

	"github.com/Asadus16/comapi/pkg/types"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [test-file]",
	Short: "Run API tests from a YAML file",
	Long: `Run API tests defined in a YAML configuration file.

Example:
  comapi run tests.yaml
  comapi run examples/sample.yaml`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		testFile := args[0]
		
		// Check if file exists
		if _, err := os.Stat(testFile); os.IsNotExist(err) {
			fmt.Printf("❌ Test file not found: %s\n", testFile)
			os.Exit(1)
		}
		
		fmt.Printf("🧭 Running tests from: %s\n", testFile)
		
		// Load and parse the test configuration
		suite, err := config.LoadTestSuite(testFile)
		if err != nil {
			fmt.Printf("❌ Failed to load test suite: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("📋 Test Suite: %s\n", suite.Name)
		fmt.Printf("🌐 Base URL: %s\n", suite.BaseURL)
		fmt.Printf("🧪 Running %d test(s)...\n\n", len(suite.Tests))
		
		// Create HTTP client
		httpClient := runner.NewHTTPClient(suite.BaseURL, suite.Headers)
		
		// Run each test
		var results []types.TestResult
		for i, test := range suite.Tests {
			fmt.Printf("Running test %d/%d: %s\n", i+1, len(suite.Tests), test.Name)
			
			result := httpClient.ExecuteTest(test)
			results = append(results, result)
			
			// Show basic result
			if result.Status == types.StatusPass {
				fmt.Printf("  ✅ %s - %dms\n", result.Status, result.Duration.Milliseconds())
			} else {
				fmt.Printf("  ❌ %s - %dms\n", result.Status, result.Duration.Milliseconds())
				if result.Error != "" {
					fmt.Printf("    Error: %s\n", result.Error)
				}
			}
			
			// Show assertion details
			for _, assertion := range result.Assertions {
				if assertion.Passed {
					fmt.Printf("    ✅ %s: %s\n", assertion.Type, assertion.Message)
				} else {
					fmt.Printf("    ❌ %s: %s\n", assertion.Type, assertion.Message)
				}
			}
			
			// Show response body for failed tests (first 200 characters)
			if result.Status == types.StatusFail {
				responsePreview := result.Response.Body
				if len(responsePreview) > 200 {
					responsePreview = responsePreview[:200] + "..."
				}
				fmt.Printf("    📄 Response: %s\n", responsePreview)
			}
			
			fmt.Println() // Add blank line between tests
		}
		
		fmt.Printf("\n🎯 Test Summary:\n")
		passed := 0
		for _, result := range results {
			if result.Status == types.StatusPass {
				passed++
			}
		}
		fmt.Printf("  ✅ Passed: %d/%d\n", passed, len(results))
		if passed < len(results) {
			fmt.Printf("  ❌ Failed: %d/%d\n", len(results)-passed, len(results))
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	
	// Add flags for output format, verbose mode, etc.
	runCmd.Flags().StringP("output", "o", "console", "Output format (console, json, html)")
	runCmd.Flags().BoolP("verbose", "v", false, "Verbose output")
	runCmd.Flags().StringP("env", "e", "", "Environment file for variable substitution")
}