package source_xsktcomvn

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// DataSourceXsktcomvn downloads data from "xskt.com.vn"
type DataSourceXsktcomvn struct{}

func (s DataSourceXsktcomvn) DownloadHTML() ([]byte, error) {
	r, err := http.NewRequest("GET", "https://xskt.com.vn/xsmb", nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	httpClient := &http.Client{Timeout: 64 * time.Second}
	w, err := httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer w.Body.Close()
	b, err := io.ReadAll(w.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}
	return b, err
}
