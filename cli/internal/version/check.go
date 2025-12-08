package version

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-version"
)

// CheckForUpdate checks if there is a newer version available on GitHub.
// It returns a warning message if an update is available, or an empty string otherwise.
func CheckForUpdate(currentVersion string) string {
	// Skip if env var is set
	if os.Getenv("SPT_FLOW_SKIP_UPDATE_CHECK") == "true" {
		return ""
	}

	// Skip if current version is a dev build or empty
	if currentVersion == "dev" || currentVersion == "" {
		return ""
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Fetch latest release
	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.github.com/repos/SmartPointsTech/turbo-flow-claude/releases/latest", nil)
	if err != nil {
		return "" // Fail silently
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "" // Fail silently
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "" // Fail silently
	}

	var release struct {
		TagName string `json:"tag_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "" // Fail silently
	}

	// Parse versions
	vCurrent, err := version.NewVersion(currentVersion)
	if err != nil {
		return ""
	}

	vLatest, err := version.NewVersion(release.TagName)
	if err != nil {
		return ""
	}

	// Compare
	if vLatest.GreaterThan(vCurrent) {
		return fmt.Sprintf("\nUpdate available: %s -> %s\nRun 'go install github.com/SmartPointsTech/turbo-flow-claude/cli/cmd/spt-flow@latest' to update.\n", vCurrent, vLatest)
	}

	return ""
}
