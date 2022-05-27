package tgyaimg

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"main/imgbb"
	"main/rndstck"
	"main/yaimg"
	"mvdan.cc/xurls/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

const DataDirPath = "data"

func Start(pref tele.Settings, imgbbToken string, whitelist ...string) {
	_, err := os.Stat(yaimg.PyScriptPath)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = os.Stat(DataDirPath)
	if err != os.ErrNotExist {
		err = os.Mkdir(DataDirPath, 0644)
		if err != nil && err != os.ErrExist {
			log.Fatalln(err)
		}
	} else if err != nil {
		log.Fatalln(err)
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
	}

	adminList := make([]*tele.Chat, 0)

	whiteMap := map[string]bool{}
	for _, whiteName := range whitelist {
		if adminID, err := strconv.Atoi(whiteName); err == nil && adminID != 0 {
			adminChat, err := bot.ChatByID(int64(adminID))
			if err != nil {
				log.Fatalf("When getting admin list with %d, an error occured:\n %s\n", adminID, err)
			}
			adminList = append(adminList, adminChat)
		} else {
			whiteMap[whiteName] = true
		}
	}

	if len(adminList) > 0 {
		bot.OnError = func(err error, ctx tele.Context) {
			log.Println(err)
			for _, chat := range adminList {
				_, err = bot.Send(chat, fmt.Sprintf("An error. Пук.\n%s\n@%s", err.Error(), ctx.Chat().Username))
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

	urlRegex := xurls.Relaxed()

	bot.Handle(tele.OnText, func(ctx tele.Context) error {
		if containsOrEmpty(whiteMap, ctx.Chat().Username) {
			for _, url := range urlRegex.FindAllString(ctx.Text(), -1) {
				stickerMsg, err := ctx.Bot().Reply(ctx.Message(), rndstck.Get())
				if err != nil {
					return err
				}

				err = FindAndReportSources(ctx, url)
				if err != nil {
					return err
				}

				err = ctx.Bot().Delete(stickerMsg)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	bot.Handle(tele.OnPhoto, func(ctx tele.Context) error {
		if containsOrEmpty(whiteMap, ctx.Chat().Username) {
			stickerMsg, err := ctx.Bot().Reply(ctx.Message(), rndstck.Get())
			if err != nil {
				return err
			}

			path, err := DownloadTelegramFile(bot, &ctx.Message().Photo.File, DataDirPath)
			if err != nil {
				return err
			}

			url, err := imgbb.Post(imgbbToken, path)
			if err != nil {
				return err
			}

			err = FindAndReportSources(ctx, url)
			if err != nil {
				return err
			}

			err = os.Remove(path)
			if err != nil {
				return err
			}

			err = ctx.Bot().Delete(stickerMsg)
			if err != nil {
				return err
			}
		}

		return nil
	})

	bot.Start()
}

func DownloadTelegramFile(bot *tele.Bot, file *tele.File, directory ...string) (filePath string, err error) {
	fileInfo, err := bot.FileByID(file.FileID)
	if err != nil {
		return
	}
	localFilename := fmt.Sprintf("%s_%s",
		time.Now().Format("mehfabric_2006-01-02_15-04-05"),
		fileInfo.FilePath[strings.LastIndex(fileInfo.FilePath, "file_")+5:])

	var dir string
	if len(directory) > 0 {
		dir = directory[0]
	} else {
		dir = "."
	}

	filePath = fmt.Sprintf("%s/%s", dir, localFilename)
	err = bot.Download(file, filePath)
	if err != nil {
		return
	}

	return
}

func FindAndReportSources(ctx tele.Context, url string) (err error) {
	sources, orig, err := yaimg.YandexGetSources(url)
	if err != nil {
		return err
	}
	report := yaimg.ReportToSliceOfStrings(sources, orig)

	for _, msg := range report {
		err = ctx.Reply(msg)
		if err != nil {
			return err
		}
	}

	return
}

func containsOrEmpty(set map[string]bool, key string) bool {
	_, isContained := set[key]
	return isContained || len(set) == 0
}
