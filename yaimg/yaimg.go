package yaimg

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

const PyScriptPath = "yaimg/ya_img.py"

type PicInfo struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func YandexGetSources(url string) (sources []PicInfo, originalImg PicInfo, err error) {
	type yaImgResponse struct {
		Src   []PicInfo `json:"src"`
		Warn  []string  `json:"warn"`
		Sizes []int     `json:"sizes"`
	}

	outputStr, err := exec.Command("python3",
		PyScriptPath,
		"--url", url).Output()
	if err != nil {
		return
	}

	var resp yaImgResponse
	err = json.Unmarshal(outputStr, &resp)
	if err != nil {
		return
	}

	sources = resp.Src
	originalImg = PicInfo{
		Url:    url,
		Width:  resp.Sizes[0],
		Height: resp.Sizes[1],
	}

	return
}

func ReportToSliceOfStrings(sources []PicInfo, origImg PicInfo) (report []string) {
	lesser := make([]string, 0)
	for _, el := range sources {
		msg := fmt.Sprintf("%d×%d (%+d, %+d) | %s\n", el.Width, el.Height,
			el.Width-origImg.Width, el.Height-origImg.Width, el.Url)
		if el.Width < origImg.Width && el.Height < origImg.Height {
			lesser = append(lesser, msg)
		} else {
			report = append(report, msg)
		}
	}

	if len(lesser) > 0 {
		report = append(report, strings.Join(lesser, "\n"))
	}
	if len(report) == 0 {
		report = append(report, fmt.Sprintf("Nothing was found :c"))
	}

	report = append(report, fmt.Sprintf("%d×%d | %s", origImg.Width, origImg.Height, origImg.Url))

	return
}
