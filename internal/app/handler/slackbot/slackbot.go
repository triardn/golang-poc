package slackbot

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/kitabisa/golang-poc/internal/app/handler"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

type SlackBotHandler struct {
	handler.HandlerOption
	http.Handler
}

func (h SlackBotHandler) Request(w http.ResponseWriter, r *http.Request) (data interface{}, pageToken *string, err error) {
	secret := h.Options.Config.Get("SLACK_SIGNING_SECRET").(string)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}

	sv, err := slack.NewSecretsVerifier(r.Header, secret)
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err = sv.Write(body); err != nil {
		return
	}

	if err = sv.Ensure(); err != nil {
		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())
	if err != nil {
		return
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err = json.Unmarshal([]byte(body), &r)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
	}

	// make sure testing nya yg bener, since ini event yg di handle itu message event, semua message yg masuk di channel itu bakal di handle di sini
	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.MessageEvent:
			fmt.Println("ev.User: ", ev.User)
			if ev.Files != nil {
				// do check file type. only accept csv
				if ev.Files[0].Filetype == "csv" {
					req, _ := http.NewRequest("GET", ev.Files[0].URLPrivateDownload, nil)
					req.Header.Set("Authorization", "Bearer "+h.Options.Config.GetString("SLACK_TOKEN"))
					client := new(http.Client)
					res, err := client.Do(req)
					if err != nil {
						log.Println(err)
					}
					defer res.Body.Close()

					body, err := ioutil.ReadAll(res.Body)
					if err != nil {
						log.Println(err)
					}

					file, err := os.OpenFile("./sample.csv", os.O_CREATE|os.O_WRONLY, 0666)
					if err != nil {
						log.Println(err)
					}
					defer file.Close()

					_, err = file.Write(body)
					if err != nil {
						log.Println(err)
					}
					log.Printf("Download finished: %s\n", ev.Files[0].URLPrivateDownload)

					f, err := os.Open("./sample.csv")
					if err != nil {
						log.Fatal("Unable to read input file ./sample.csv", err)
					}
					defer f.Close()

					csvReader := csv.NewReader(f)
					records, err := csvReader.ReadAll()
					if err != nil {
						log.Fatal("Unable to parse file as CSV for ./sample.csv", err)
					}

					for i, record := range records {
						if i == 0 {
							continue
						}

						// replace bagian bawah ini dengan bisnis proses (save db, call external service, etc...)
						fmt.Println("Ayo kita kirim wa ke ", record[0])
					}

					// react postingan
					api := slack.New(h.Options.Config.GetString("SLACK_BOT_TOKEN"))

					itemRef := slack.ItemRef{
						Channel:   ev.Channel,
						Timestamp: ev.TimeStamp,
						File:      ev.Files[0].ID,
					}

					_ = api.AddReaction("okay", itemRef)
					_ = api.AddReaction("thumbsup", itemRef)
					_ = api.AddReaction("pinguin_dance", itemRef)

					// remove reaction doesn't works, need to research more
					err = api.RemoveReaction("okay", itemRef)
					fmt.Println(err)
					err = api.RemoveReaction("pinguin_dance", itemRef)
					fmt.Println(err)

					// send message & mention users that upload csv
					_, _, _ = api.PostMessage(ev.Channel, slack.MsgOptionText(fmt.Sprintf("Hi <@%s>, Done!", ev.User), false))
				}
			}
		}
	}

	return
}
