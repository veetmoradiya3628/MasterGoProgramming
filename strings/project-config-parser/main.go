package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func parseConfig(content string) (map[string]string, error) {
	config := make(map[string]string)
	re := regexp.MustCompile(`^\s*([A-Z0-9_]+)\s*=\s*(?:"([^"]*)"|'([^']*)'|(.*))\s*$`)
	scanner := bufio.NewScanner(strings.NewReader(content))
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := scanner.Text()

		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" || strings.HasPrefix(trimmedLine, "#") {
			continue
		}
		matches := re.FindStringSubmatch(trimmedLine)
		if matches == nil {
			fmt.Printf("Line %d, '%s' - Is invalid\n", lineNo, line)
			continue
		}
		key := matches[1]
		var value string
		if matches[2] != "" {
			value = matches[2]
		} else if matches[3] != "" {
			value = matches[3]
		} else {
			value = matches[4]
		}
		config[key] = value
	}
	return config, nil
}

func main() {
	envFileContent := `
# App
APP_NAME=sample-service
APP_ENV=development
APP_PORT=8080
APP_DEBUG=true

# Database (Postgres)
DB_HOST=localhost
DB_PORT=5432
DB_NAME=sample_db
DB_USER=sample_user
DB_PASSWORD=sample_password

# Auth
JWT_SECRET=change_this_secret
JWT_EXPIRY_MINUTES=60

# Cache (Redis)
REDIS_HOST=localhost
REDIS_PORT=6379
CACHE_TTL_SECONDS=300

# External Service
PAYMENT_SERVICE_URL=https://payment.mock.local
PAYMENT_TIMEOUT_MS=5000

# Logging
LOG_LEVEL=info

# Feature flag
FEATURE_NEW_DASHBOARD=true
`
	config, err := parseConfig(envFileContent)
	if err != nil {
		os.Exit(1)
	}
	for k, v := range config {
		fmt.Printf("%s = %q\n", k, v)
	}
}
