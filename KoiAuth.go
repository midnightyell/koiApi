package koiApi

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type koiAuth struct {
	ServerURL *string `json:"server"`
	Username  *string `json:"user"`
	Password  *string `json:"password"`
}

var (
	verbose    bool
	configFile string = ".koiauth"
	auth       koiAuth
	dprefix    string = "koi-"
)

func init() {
	// Temporary variables for flag parsing
	var serverURL, username, password string

	// Define flags
	flag.BoolVar(&verbose, dprefix+"verbose", false, "Enable verbose output")
	flag.StringVar(&configFile, dprefix+"config", configFile, "Path to config file")
	flag.StringVar(&serverURL, dprefix+"server", "", "Server URL")
	flag.StringVar(&username, dprefix+"user", "", "Username")
	flag.StringVar(&password, dprefix+"password", "", "Password")

	// Parse os.Args
	flag.CommandLine.Parse(os.Args[1:])

	// Resolve config file path: use home directory if no '/' in path
	configPath := configFile
	if configPath != "" && !strings.Contains(configPath, "/") {
		home, err := os.UserHomeDir()
		if err != nil && verbose {
			fmt.Printf("Error getting home directory: %v\n", err)
		} else {
			configPath = filepath.Join(home, configPath)
		}
	}

	// Read config file (if it exists) to set auth fields
	data, err := os.ReadFile(configPath)
	if err == nil {
		if err := json.Unmarshal(data, &auth); err != nil && verbose {
			fmt.Printf("Error parsing config %s: %v\n", configPath, err)
		}
	} else if verbose && !os.IsNotExist(err) {
		fmt.Printf("Error reading config %s: %v\n", configPath, err)
	}

	// Override with environment variables (KOI_SERVER, KOI_USER, KOI_PASSWORD)
	if val, ok := os.LookupEnv("KOI_SERVER"); ok && val != "" {
		auth.ServerURL = &val
	}
	if val, ok := os.LookupEnv("KOI_USER"); ok && val != "" {
		auth.Username = &val
	}
	if val, ok := os.LookupEnv("KOI_PASSWORD"); ok && val != "" {
		auth.Password = &val
	}

	// Override with command-line values if provided
	if serverURL != "" {
		auth.ServerURL = &serverURL
	}
	if username != "" {
		auth.Username = &username
	}
	if password != "" {
		auth.Password = &password
	}
	if verbose {
		fmt.Printf("Koi auth: %s / %s @ %s\n", *auth.Username, *auth.Password, *auth.ServerURL)
	}
}

// Usage prints configuration options for koiAuth
func KoiAuthUsage() {
	fmt.Fprintf(os.Stderr, "Configuration options for koiAuth:\n")
	fmt.Fprintf(os.Stderr, "\nConfig file:\n")
	fmt.Fprintf(os.Stderr, "  - %s (default: %s in home directory if no '/')\n", dprefix+"config", configFile)
	fmt.Fprintf(os.Stderr, "    JSON file with fields: server, user, password\n")
	fmt.Fprintf(os.Stderr, "    Example: ~/.koiauth with {\"server\": \"http://example.com\", \"user\": \"user\", \"password\": \"pass\"}\n")
	fmt.Fprintf(os.Stderr, "\nEnvironment variables (override config file):\n")
	fmt.Fprintf(os.Stderr, "  - KOI_SERVER: Server URL\n")
	fmt.Fprintf(os.Stderr, "  - KOI_USER: Username\n")
	fmt.Fprintf(os.Stderr, "  - KOI_PASSWORD: Password\n")
	fmt.Fprintf(os.Stderr, "\nCommand-line flags (override config file and env vars):\n")
	fmt.Fprintf(os.Stderr, "  - %s: Enable verbose error logging\n", dprefix+"verbose")
	fmt.Fprintf(os.Stderr, "  - %s: Server URL\n", dprefix+"server")
	fmt.Fprintf(os.Stderr, "  - %s: Username\n", dprefix+"user")
	fmt.Fprintf(os.Stderr, "  - %s: Password\n", dprefix+"password")
	fmt.Fprintf(os.Stderr, "\nPrecedence: config file < environment variables < command-line flags\n")
}
