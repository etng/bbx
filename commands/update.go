package commands

import (
	"fmt"
	"github.com/etng/bbx/helpers"
	"github.com/etng/bbx/version"
	"github.com/imroc/req/v3"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func buildOsArchList() (osArchList []string) {
	osList := []string{
		runtime.GOOS,
		strings.Title(runtime.GOOS),
	}
	archList := []string{
		runtime.GOARCH,
	}
	if runtime.GOARCH == "amd64" {
		archList = append(archList, "x86_64")
	}
	for _, os := range osList {
		for _, arch := range archList {
			osArchList = append(osArchList, fmt.Sprintf("%s_%s",
				os, arch))
		}
	}
	return
}
func init() {
	allCommands = append(allCommands, Command{
		Name:      "update",
		Desc:      "update bbx it self",
		AliasList: []string{"self-update"},
		Handler: func(args ...string) {
			var release Release
			req.SetResult(&release).MustGet("https://api.github.com/repos/etng/bbx/releases/latest")
			if release.TagName == version.Version {
				fmt.Println("you are using the latest version")
				return
			}
			//bbx_0.0.4_Darwin_arm64.tar.gz
			targetList := []string{}
			for _, osArch := range buildOsArchList() {
				targetList = append(targetList, fmt.Sprintf("bbx_%s_%s.tar.gz",
					release.Name[1:], osArch))
			}
			var asset *ReleaseAsset
			for _, _asset := range release.Assets {
				fmt.Println("checking", _asset.Name, targetList)
				if helpers.StrSliceIn(_asset.Name, targetList) {
					asset = _asset
				}
			}
			if asset == nil {
				fmt.Println("no matched package")
				os.Exit(1)
			}
			realMe, e := exec.LookPath(os.Args[0])
			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			}
			fmt.Println("should override ", realMe)
			if strings.Contains(realMe, "go-build") {
				realMe = "data/tmp/updated_bbx"
				fmt.Println("realme updated to ", realMe)
			}
			outFile, e := os.Create(asset.Name)
			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			}
			req.SetOutput(outFile).MustGet(asset.BrowserDownloadUrl)
			helpers.PickInTgz(asset.Name, "bbx", realMe)
			fmt.Println("updated!")
		},
		Weight: 0,
	})
}

type Release struct {
	TagName     string          `json:"tag_name"`
	Name        string          `json:"name"`
	Prerelease  bool            `json:"prerelease"`
	CreatedAt   time.Time       `json:"created_at"`
	PublishedAt time.Time       `json:"published_at"`
	Assets      []*ReleaseAsset `json:"assets"`
	Body        string          `json:"body"`
}
type ReleaseAsset struct {
	Name               string    `json:"name"`
	ContentType        string    `json:"content_type"`
	State              string    `json:"state"`
	Size               int       `json:"size"`
	DownloadCount      int       `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	BrowserDownloadUrl string    `json:"browser_download_url"`
}
