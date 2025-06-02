// package cmd

// import (
// 	"fmt"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/cobra"
// 	"github.com/Asadus16/comapi/internal/runner"
// 	"github.com/Asadus16/comapi/pkg/types"
// )

// // serverCmd represents the server command
// var serverCmd = &cobra.Command{
// 	Use:   "server",
// 	Short: "Start Comapi web server",
// 	Long:  `Start the Comapi web server to provide a REST API for the frontend.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		port, _ := cmd.Flags().GetString("port")
// 		startServer(port)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(serverCmd)
// 	serverCmd.Flags().StringP("port", "p", "8080", "Port to run the server on")
// }

// func startServer(port string) {
// 	// Set Gin to release mode for cleaner output
// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.Default()
	
// 	// Enable CORS for React frontend
// 	r.Use(func(c *gin.Context) {
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		c.Header("Access-Control-Allow-Headers", "Content-Type")
		
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
		
// 		c.Next()
// 	})

// 	// API Routes
// 	api := r.Group("/api/v1")
// 	{
// 		api.POST("/tests/run", runTestsEndpoint)
// 		api.POST("/tests/validate", validateTestsEndpoint)
// 		api.GET("/health", healthCheckEndpoint)
// 	}

// 	fmt.Printf("🧭 Comapi server starting on port %s\n", port)
// 	fmt.Printf("🌐 API: http://localhost:%s/api/v1\n", port)
// 	fmt.Printf("📡 Health Check: http://localhost:%s/api/v1/health\n", port)
// 	fmt.Printf("🚀 Ready to accept requests from React frontend!\n\n")
	
// 	r.Run(":" + port)
// }

// // API Endpoints

// // runTestsEndpoint - POST /api/v1/tests/run
// func runTestsEndpoint(c *gin.Context) {
// 	var request struct {
// 		TestSuite types.TestSuite `json:"test_suite"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid request format: " + err.Error()})
// 		return
// 	}

// 	// Validate basic requirements
// 	if request.TestSuite.Name == "" {
// 		c.JSON(400, gin.H{"error": "Test suite name is required"})
// 		return
// 	}
	
// 	if request.TestSuite.BaseURL == "" {
// 		c.JSON(400, gin.H{"error": "Base URL is required"})
// 		return
// 	}
	
// 	if len(request.TestSuite.Tests) == 0 {
// 		c.JSON(400, gin.H{"error": "At least one test is required"})
// 		return
// 	}

// 	fmt.Printf("📋 Running test suite: %s\n", request.TestSuite.Name)
// 	fmt.Printf("🌐 Base URL: %s\n", request.TestSuite.BaseURL)
// 	fmt.Printf("🧪 Tests: %d\n", len(request.TestSuite.Tests))

// 	// Create HTTP client
// 	httpClient := runner.NewHTTPClient(request.TestSuite.BaseURL, request.TestSuite.Headers)

// 	// Run all tests
// 	var results []types.TestResult
// 	for i, test := range request.TestSuite.Tests {
// 		fmt.Printf("  Running test %d/%d: %s\n", i+1, len(request.TestSuite.Tests), test.Name)
		
// 		result := httpClient.ExecuteTest(test)
// 		results = append(results, result)
		
// 		// Log result
// 		if result.Status == types.StatusPass {
// 			fmt.Printf("    ✅ PASS - %dms\n", result.Duration.Milliseconds())
// 		} else {
// 			fmt.Printf("    ❌ FAIL - %s\n", result.Error)
// 		}
// 	}

// 	// Calculate summary
// 	passed := 0
// 	for _, result := range results {
// 		if result.Status == types.StatusPass {
// 			passed++
// 		}
// 	}

// 	fmt.Printf("📊 Summary: %d/%d passed\n\n", passed, len(results))

// 	response := gin.H{
// 		"suite_name":    request.TestSuite.Name,
// 		"total_tests":   len(results),
// 		"passed_tests":  passed,
// 		"failed_tests":  len(results) - passed,
// 		"results":       results,
// 	}

// 	c.JSON(200, response)
// }

// // validateTestsEndpoint - POST /api/v1/tests/validate
// func validateTestsEndpoint(c *gin.Context) {
// 	var request struct {
// 		TestSuite types.TestSuite `json:"test_suite"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid request format"})
// 		return
// 	}

// 	// Validate the test suite
// 	errors := validateTestSuite(request.TestSuite)

// 	if len(errors) > 0 {
// 		c.JSON(400, gin.H{
// 			"valid": false,
// 			"errors": errors,
// 		})
// 		return
// 	}

// 	c.JSON(200, gin.H{
// 		"valid": true,
// 		"message": "Test suite is valid",
// 	})
// }

// // healthCheckEndpoint - GET /api/v1/health
// func healthCheckEndpoint(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"status": "ok",
// 		"service": "comapi",
// 		"version": "1.0.0",
// 		"timestamp": time.Now().Unix(),
// 	})
// }

// // Helper function to validate test suite
// func validateTestSuite(suite types.TestSuite) []string {
// 	var errors []string

// 	if suite.Name == "" {
// 		errors = append(errors, "Test suite name is required")
// 	}

// 	if suite.BaseURL == "" {
// 		errors = append(errors, "Base URL is required")
// 	}

// 	if len(suite.Tests) == 0 {
// 		errors = append(errors, "At least one test is required")
// 	}

// 	for i, test := range suite.Tests {
// 		if test.Name == "" {
// 			errors = append(errors, fmt.Sprintf("Test %d: name is required", i+1))
// 		}
// 		if test.Method == "" {
// 			errors = append(errors, fmt.Sprintf("Test '%s': method is required", test.Name))
// 		}
// 		if test.Path == "" {
// 			errors = append(errors, fmt.Sprintf("Test '%s': path is required", test.Name))
// 		}
// 	}

// 	return errors
// }



package cmd

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/Asadus16/comapi/internal/runner"
	"github.com/Asadus16/comapi/pkg/types"
)


// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Comapi web server",
	Long:  `Start the Comapi web server to provide a REST API for the frontend.`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		startServer(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("port", "p", "8080", "Port to run the server on")
}

func startServer(port string) {
	// Set Gin to release mode for cleaner output
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	
	// Enable CORS for React frontend
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// API Routes
	api := r.Group("/api/v1")
	{
		api.POST("/tests/run", runTestsEndpoint)
		api.POST("/tests/validate", validateTestsEndpoint)
		api.GET("/health", healthCheckEndpoint)
	}

	fmt.Printf("🧭 Comapi server starting on port %s\n", port)
	fmt.Printf("🌐 API: http://localhost:%s/api/v1\n", port)
	fmt.Printf("📡 Health Check: http://localhost:%s/api/v1/health\n", port)
	fmt.Printf("🚀 Ready to accept requests from React frontend!\n\n")
	
	r.Run(":" + port)
}

// API Endpoints

// runTestsEndpoint - POST /api/v1/tests/run
func runTestsEndpoint(c *gin.Context) {
	var request struct {
		TestSuite types.TestSuite `json:"test_suite"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format: " + err.Error()})
		return
	}

	// For single test mode, we don't need suite name validation
	if len(request.TestSuite.Tests) == 0 {
		c.JSON(400, gin.H{"error": "Test is required"})
		return
	}

	// Get the single test
	test := request.TestSuite.Tests[0]
	
	if test.Name == "" {
		c.JSON(400, gin.H{"error": "Test name is required"})
		return
	}
	
	if test.URL == "" {
		c.JSON(400, gin.H{"error": "Complete URL is required"})
		return
	}

	fmt.Printf("🧪 Running test: %s\n", test.Name)
	fmt.Printf("🌐 URL: %s\n", test.URL)
	fmt.Printf("📡 Method: %s\n", test.Method)

	// For single URL mode, we extract base URL and path
	// Create a dummy base URL and set the full URL as the path
	httpClient := runner.NewHTTPClient("", map[string]string{})

	// Run the single test
	result := httpClient.ExecuteTestWithFullURL(test)
	
	// Log result
	if result.Status == types.StatusPass {
		fmt.Printf("  ✅ PASS - %dms\n", result.Duration.Milliseconds())
	} else {
		fmt.Printf("  ❌ FAIL - %s\n", result.Error)
	}

	fmt.Printf("📊 Test completed\n\n")

	response := gin.H{
		"suite_name":    test.Name,
		"total_tests":   1,
		"passed_tests":  func() int { if result.Status == types.StatusPass { return 1 }; return 0 }(),
		"failed_tests":  func() int { if result.Status == types.StatusFail { return 1 }; return 0 }(),
		"results":       []types.TestResult{result},
	}

	c.JSON(200, response)
}

// validateTestsEndpoint - POST /api/v1/tests/validate
func validateTestsEndpoint(c *gin.Context) {
	var request struct {
		TestSuite types.TestSuite `json:"test_suite"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	// Validate the test suite
	errors := validateTestSuite(request.TestSuite)

	if len(errors) > 0 {
		c.JSON(400, gin.H{
			"valid": false,
			"errors": errors,
		})
		return
	}

	c.JSON(200, gin.H{
		"valid": true,
		"message": "Test suite is valid",
	})
}

// healthCheckEndpoint - GET /api/v1/health
func healthCheckEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"service": "comapi",
		"version": "1.0.0",
		"timestamp": time.Now().Unix(),
	})
}

// Helper function to validate test suite
func validateTestSuite(suite types.TestSuite) []string {
	var errors []string

	if suite.Name == "" {
		errors = append(errors, "Test suite name is required")
	}

	if suite.BaseURL == "" {
		errors = append(errors, "Base URL is required")
	}

	if len(suite.Tests) == 0 {
		errors = append(errors, "At least one test is required")
	}

	for i, test := range suite.Tests {
		if test.Name == "" {
			errors = append(errors, fmt.Sprintf("Test %d: name is required", i+1))
		}
		if test.Method == "" {
			errors = append(errors, fmt.Sprintf("Test '%s': method is required", test.Name))
		}
		if test.Path == "" {
			errors = append(errors, fmt.Sprintf("Test '%s': path is required", test.Name))
		}
	}

	return errors
}