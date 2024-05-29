package source_xsktcomvn

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/daominah/crawler_xsmb/pkg/core"
)

//go:embed parse_xsktcomvn_test.html
var testDataXsktcomvn string // result of TestDataSourceXsktcomvn_DownloadHTML

func TestParseXsktcomvn(t *testing.T) {
	if len(strings.TrimSpace(testDataXsktcomvn)) == 0 {
		t.Fatalf("empty testDataXsktcomvn")
	}
	got, err := (DataSourceXsktcomvn{}).ParseToResultXSMB(testDataXsktcomvn)
	if err != nil {
		t.Fatalf("error Parsexsktcomvn: %v", err)
	}
	want := core.ResultXSMB{
		DateISO: "2024-05-28",
		GiaiDB:  "47490",
		Giai1:   "72043",
		Giai2:   []string{"09830", "29003"},
		Giai3:   []string{"30879", "15157", "62025", "16755", "23357", "05880"},
		Giai4:   []string{"8291", "6953", "9981", "1132"},
		Giai5:   []string{"0811", "0341", "2380", "4935", "4914", "8694"},
		Giai6:   []string{"232", "937", "880"},
		Giai7:   []string{"63", "29", "75", "12"},
	}
	if got.DateISO != want.DateISO {
		t.Errorf("error DateISO, got %v, want %v", got.DateISO, want.DateISO)
	}
	if got.GiaiDB != want.GiaiDB {
		t.Errorf("error GiaiDB, got %v, want %v", got.GiaiDB, want.GiaiDB)
	}
	if len(got.Giai3) != len(want.Giai3) {
		t.Errorf("error length Giai7, got %v, want %v", len(got.Giai3), len(want.Giai3))
	} else {
		if got.Giai3[5] != want.Giai3[5] {
			t.Errorf("error Giai7, got %v, want %v", got.Giai3[0], want.Giai3[0])
		}
	}
	if len(got.Giai7) != len(want.Giai7) {
		t.Errorf("error length Giai7, got %v, want %v", len(got.Giai7), len(want.Giai7))
	} else {
		if got.Giai7[0] != want.Giai7[0] {
			t.Errorf("error Giai7, got %v, want %v", got.Giai7[0], want.Giai7[0])
		}
	}
	t.Logf("got: %+v", got)
}
