package source_xsktcomvn

import (
	"fmt"
	"strings"
	"time"

	"github.com/daominah/crawler_xsmb/pkg/core"
	"github.com/mywrap/textproc"
	"golang.org/x/net/html"
)

func (s DataSourceXsktcomvn) ParseToResultXSMB(rawHTML string) (ret core.ResultXSMB, err error) {
	rootNode := textproc.HTMLParseToNode(rawHTML)
	resultBoxs, err := textproc.HTMLXPath(rootNode, `//*[@class="box-ketqua"]`)
	if err != nil || len(resultBoxs) == 0 {
		return ret, fmt.Errorf("HTMLXPath box-ketqua, err %v", err)
	}
	resultBox := resultBoxs[0]

	dates, err := textproc.HTMLXPath(resultBox, `//*[contains(@id,"MBngay")]`)
	if err != nil || len(dates) == 0 {
		return ret, fmt.Errorf("HTMLXPath MBNgay, err %v", err)
	}
	// theirDate = "MBngay28-05"
	theirDate := getHTMLAttr(dates[0], "id")
	words := strings.Split(strings.TrimPrefix(theirDate, "MBngay"), "-")
	if len(words) < 2 {
		return ret, fmt.Errorf("unexpected date data format")
	}
	ret.DateISO = fmt.Sprintf("%v-%v-%v", time.Now().Year(), words[1], words[0])

	tableRows, err := textproc.HTMLXPath(resultBox, `//tr`)
	if err != nil {
		return ret, fmt.Errorf("HTMLXPath box-ketqua tr, err %v", err)
	}

	for _, tr := range tableRows {
		cells, err := textproc.HTMLXPath(tr, `//td`)
		if err != nil || len(cells) < 2 {
			continue
		}
		resultNumbersStr := textproc.HTMLGetText(cells[1])
		resultNumbers := strings.Fields(resultNumbersStr)
		switch getHTMLAttr(cells[0], "title") {
		case "Giải ĐB":
			if len(resultNumbers) > 0 {
				ret.GiaiDB = resultNumbers[0]
			}
		case "Giải nhất":
			if len(resultNumbers) > 0 {
				ret.Giai1 = resultNumbers[0]
			}
		case "Giải nhì":
			ret.Giai2 = resultNumbers
		case "Giải ba":
			ret.Giai3 = resultNumbers
		case "Giải tư":
			ret.Giai4 = resultNumbers
		case "Giải năm":
			ret.Giai5 = resultNumbers
		case "Giải sáu":
			ret.Giai6 = resultNumbers
		case "Giải bảy":
			ret.Giai7 = resultNumbers
		default:
			// do nothing
		}
	}

	return ret, nil
}

func getHTMLAttr(node *html.Node, attrKey string) string {
	for _, attr := range node.Attr {
		if attr.Key != attrKey {
			continue
		}
		return attr.Val
	}
	return ""
}
