package services

import (
	"os"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

func HandleCron() {
	count, _ := strconv.Atoi(os.Getenv("CONTENT_COUNT"))
	location, _ := time.LoadLocation("America/Toronto")

	c := cron.New(
		cron.WithLocation(location),
	)

	c.AddFunc("@every 8h", func() {
		contents := getRandomContents(int(count))

		if len(contents) == 0 {
			refillTipsCollection()

			return
		}

		for _, tip := range contents {
			deleteTip(tip["_id"])

			filePath, _ := tip["path"].(string)
			res := getContent(filePath)

			if res != "" {
				kafkaMsg := KafkaMessage{Data: res}
				emitToKafka(kafkaMsg)
			}
		}
	})

	c.Start()

	select {}
}
