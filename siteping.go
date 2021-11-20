package siteping

import (
	"io"
	"net/http"
	"strings"
	"time"
)

type Site struct {
	Url     string
	Content string
}

// PingSite tries to reach a URL; pass an optional string
// to check for specific content in the response body.
func PingSite(url string, content string) (bool, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	resp, err := client.Get(url)
	if err != nil {
		return false, err
	}

	if content != "" {
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}

		check := strings.Contains(string(b), content)
		return check, nil
	}

	return true, nil
}
