package core

import (
	"time"
)

type ResultXSMB struct {
	DateISO      string // YYYY-mm-dd, e.g. 2006-01-02
	GiaiDB       string // giải đặc biệt
	Giai1        string
	Giai2        []string
	Giai3        []string
	Giai4        []string
	Giai5        []string
	Giai6        []string
	Giai7        []string
	DownloadedAt time.Time
}

type DataSource interface {
	DownloadHTML() (responseBody []byte, err error)
	ParseToResultXSMB(rawHTML string) (ResultXSMB, error)
}
