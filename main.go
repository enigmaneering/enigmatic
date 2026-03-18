package main

import (
	"flag"
	"fmt"
	"os"

	"git.enigmaneering.org/enigmatic/gpu"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "fetch":
		runFetch(os.Args[2:])
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("e - The Enigmaneering Guild CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  e <command> [flags]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  fetch    Download shader compilation toolchain binaries")
	fmt.Println("  help     Show this help message")
	fmt.Println()
	fmt.Println("Run 'e <command> -help' for more information on a command.")
}

func runFetch(args []string) {
	fetchCmd := flag.NewFlagSet("fetch", flag.ExitOnError)
	version := fetchCmd.String("version", "", "Specific version to download (e.g., v0.0.42). Defaults to latest.")
	dir := fetchCmd.String("dir", "external", "Directory to install libraries (default: ./external)")

	fetchCmd.Usage = func() {
		fmt.Println("Usage: e fetch [flags]")
		fmt.Println()
		fmt.Println("Download shader compilation toolchain binaries")
		fmt.Println()
		fmt.Println("Flags:")
		fetchCmd.PrintDefaults()
		fmt.Println()
		fmt.Println("Environment Variables:")
		fmt.Println("  ENIGMATIC_GOFETCH_DIRECTORY - Override installation directory")
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  e fetch                    # Install latest version to ./external")
		fmt.Println("  e fetch -version v0.0.42   # Install specific version")
		fmt.Println("  e fetch -dir /opt/shaders  # Install to custom directory")
		fmt.Println()
		fmt.Println("Freeze Updates:")
		fmt.Println("  Create a 'FREEZE' file in the external directory to prevent")
		fmt.Println("  automatic upgrades when new versions are released.")
	}

	fetchCmd.Parse(args)

	// Set directory if specified
	if *dir != "external" {
		os.Setenv("ENIGMATIC_GOFETCH_DIRECTORY", *dir)
	}

	var err error
	if *version != "" {
		fmt.Printf("Installing shader compilation toolchain version %s...\n", *version)
		err = gpu.EnsureLibrariesVersion(*version)
	} else {
		fmt.Println("Installing latest shader compilation toolchain...")
		err = gpu.EnsureLibraries()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("✓ Shader compilation toolchain installed successfully")
}
