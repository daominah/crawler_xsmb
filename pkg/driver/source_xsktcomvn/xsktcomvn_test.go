package source_xsktcomvn

import (
	"testing"
)

func TestDataSourceXsktcomvn_DownloadHTML(t *testing.T) {
	var s DataSourceXsktcomvn
	data, err := s.DownloadHTML()
	if err != nil {
		t.Fatalf("error DataSourceXsktcomvn.DownloadHTML: %v", err)
	}
	_ = data
	//println(string(data))
}
