package commands

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

//https://core.telegram.org/bots/features#botfather

func init() {
	allCommands = append(allCommands, Command{
		Name:      "tgbot",
		Desc:      "run a telegram-bot",
		AliasList: []string{"telegram-bot", "bot"},
		Handler: func(args ...string) {
			rootDir := "data/tgbot"
			botFs := os.DirFS(rootDir)
			var (
				// Universal markup builders.
				menu     = &tele.ReplyMarkup{ResizeKeyboard: true}
				selector = &tele.ReplyMarkup{}

				// Reply buttons.
				btnHelp     = menu.Text("ℹ Help")
				btnSettings = menu.Text("⚙ Settings")

				// Inline buttons.
				//
				// Pressing it will cause the client to
				// send the bot a callback.
				//
				// Make sure Unique stays unique as per button kind
				// since it's required for callback routing to work.
				//
				btnPrev = selector.Data("⬅", "prev")
				btnNext = selector.Data("➡", "next")
			)

			menu.Reply(
				menu.Row(btnHelp),
				menu.Row(btnSettings),
			)
			selector.Inline(
				selector.Row(btnPrev, btnNext),
			)
			pref := tele.Settings{
				Token:   os.Getenv("TG_TOKEN"),
				Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
				Verbose: true,
			}

			b, err := tele.NewBot(pref)
			if err != nil {
				log.Fatal(err)
				return
			}
			b.Handle("/start", func(c tele.Context) error {
				return c.Reply("Send /hello to begin validation!", menu)
			})
			// On reply button pressed (message)
			b.Handle(&btnHelp, func(c tele.Context) error {
				return c.Edit("Here is some help: ...")
			})

			// On inline button pressed (callback)
			b.Handle(&btnPrev, func(c tele.Context) error {
				return c.Respond()
			})
			b.Handle("/hello", func(c tele.Context) error {
				return c.Send("Hello!")
			})
			b.Handle("/images", func(c tele.Context) error {
				if images, err := fs.Glob(botFs, "images/*.jpeg"); err == nil {
					img := images[rand.Intn(len(images))]
					fmt.Println(img)
					b.Send(c.Sender(), &tele.Photo{File: tele.FromDisk(filepath.Join(rootDir, img))})
				}
				return c.Send("you should received an image!")
			})
			b.Handle("/videos", func(c tele.Context) error {
				if videos, err := fs.Glob(botFs, "videos/*.mp4"); err == nil {
					video := videos[rand.Intn(len(videos))]
					b.Send(c.Sender(), &tele.Video{File: tele.FromDisk(filepath.Join(rootDir, video))})
				}
				return c.Send("you should received a video!")
			})
			b.Start()
		},
		Weight: 0,
	})
}
