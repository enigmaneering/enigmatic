# e - The Enigmaneering Guild CLI Tool

The official command-line interface for The Enigmaneering Guild's tools and utilities.

## Installation

Download the latest release for your platform:

### macOS
```bash
# ARM64 (Apple Silicon)
curl -L -o e https://git.enigmaneering.org/enigmatic/releases/latest/download/e-darwin-arm64
chmod +x e
sudo mv e /usr/local/bin/

# x86_64 (Intel)
curl -L -o e https://git.enigmaneering.org/enigmatic/releases/latest/download/e-darwin-amd64
chmod +x e
sudo mv e /usr/local/bin/
```

### Linux
```bash
# x86_64
curl -L -o e https://git.enigmaneering.org/enigmatic/releases/latest/download/e-linux-amd64
chmod +x e
sudo mv e /usr/local/bin/

# ARM64
curl -L -o e https://git.enigmaneering.org/enigmatic/releases/latest/download/e-linux-arm64
chmod +x e
sudo mv e /usr/local/bin/
```

### Windows
```powershell
# x86_64
Invoke-WebRequest -Uri "https://git.enigmaneering.org/enigmatic/releases/latest/download/e-windows-amd64.exe" -OutFile "e.exe"

# ARM64
Invoke-WebRequest -Uri "https://git.enigmaneering.org/enigmatic/releases/latest/download/e-windows-arm64.exe" -OutFile "e.exe"
```

## Usage

```bash
e <command> [flags]
```

### Commands

#### `fetch`
Download shader compilation toolchain binaries from the redistributables repository.

```bash
# Install latest version
e fetch

# Install specific version
e fetch -version v0.0.42

# Install to custom directory
e fetch -dir /opt/shaders

# Show help
e fetch -help
```

**Flags:**
- `-version` - Specific version to download (e.g., v0.0.42). Defaults to latest.
- `-dir` - Directory to install libraries (default: ./external)

**Environment Variables:**
- `ENIGMATIC_GOFETCH_DIRECTORY` - Override installation directory

**Freeze Updates:**
Create a `FREEZE` file in the external directory to prevent automatic upgrades when new versions are released.

## Go Module Usage

Use `e` functionality programmatically in your Go projects:

```go
import "git.enigmaneering.org/enigmatic/gpu"

func main() {
    // Downloads and extracts latest toolchain to ./external/
    if err := gpu.EnsureLibraries(); err != nil {
        log.Fatal(err)
    }

    // Specify a version and custom directory
    os.Setenv("ENIGMATIC_GOFETCH_DIRECTORY", "./my-shaders")
    if err := gpu.EnsureLibrariesVersion("v0.0.45"); err != nil {
        log.Fatal(err)
    }
}
```

**Install:**
```bash
go get git.enigmaneering.org/enigmatic@latest
```

## Building from Source

```bash
git clone https://github.com/enigmaneering/enigmatic.git
cd enigmatic
go build -o e
```

## License

Part of The Enigmaneering Guild's toolchain.
