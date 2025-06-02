package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [filename]",
	Short: "Create a sample test configuration file",
	Long: `Generate a sample YAML test configuration file to get started quickly.

Example:
  comapi init                    # Creates sample-tests.yaml
  comapi init my-api-tests.yaml  # Creates my-api-tests.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		filename := "sample-tests.yaml"
		if len(args) > 0 {
			filename = args[0]
		}
		
		// Check if file already exists
		if _, err := os.Stat(filename); err == nil {
			fmt.Printf("⚠️  File %s already exists. Use --force to overwrite.\n", filename)
			return
		}
		
		// Create sample test file
		sampleContent := `name: "Sample API Tests"
base_url: "https://jsonplaceholder.typicode.com"
headers:
  Content-Type: "application/json"
  User-Agent: "Comapi/1.0"

tests:
  - name: "Get single post"
    description: "Fetch a specific post by ID"
    method: "GET"
    path: "/posts/1"
    assertions:
      - type: "status"
        expected: 200
      - type: "json_path"
        target: "$.userId"
        expected: 1
      - type: "json_path"
        target: "$.title"
        operator: "contains"
        expected: "sunt"

  - name: "Create new post"
    description: "Create a new post via POST request"
    method: "POST"
    path: "/posts"
    headers:
      Content-Type: "application/json"
    body: |
      {
        "title": "Test Post",
        "body": "This is a test post created by Comapi",
        "userId": 1
      }
    assertions:
      - type: "status"
        expected: 201
      - type: "json_path"
        target: "$.title"
        expected: "Test Post"

  - name: "Get all posts"
    description: "Fetch all posts and check response time"
    method: "GET"
    path: "/posts"
    assertions:
      - type: "status"
        expected: 200
      - type: "response_time"
        operator: "less_than"
        expected: 2000
`

		err := os.WriteFile(filename, []byte(sampleContent), 0644)
		if err != nil {
			fmt.Printf("❌ Failed to create file: %v\n", err)
			return
		}
		
		fmt.Printf("✅ Created sample test file: %s\n", filename)
		fmt.Printf("🚀 Run your tests with: comapi run %s\n", filename)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	
	// Add flag to force overwrite existing files
	initCmd.Flags().BoolP("force", "f", false, "Overwrite existing file")
}