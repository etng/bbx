package commands

import (
	"fmt"
	"github.com/imroc/req/v3"
	"time"
)

func init() {
	allCommands = append(allCommands, Command{
		Name:      "ip",
		Desc:      "show current host ip info",
		AliasList: []string{"wan_ip"},
		Handler: func(args ...string) {
			urls := []string{
				"http://ipinfo.io",
				"http://ip.fm",
			}
			for _, url := range urls {
				client := req.C().
					SetUserAgent("curl/7.79.1").
					SetTimeout(5 * time.Second)
				resp, err := client.R().Get(url)
				if err != nil {
					fmt.Println("fail to fetch wan ip from url %q", url)
				} else {
					fmt.Println(url, resp)
				}
			}
		},
		Weight: 0,
	})
}
