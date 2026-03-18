# The Enigmaneering Guild's CLI Tool - `e`

Currently, this only distributes redistributables from Other.  You may use it directly through the release
binaries as `e` - or, more colloquially, by importing `git.enigmaneering.org/enigmatic` into your Go project
for continuous integration of the latest redistributables while you work.

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
Create an extensionless and empty `FREEZE` file in the external directory to prevent automatic upgrades when new versions are released.

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
