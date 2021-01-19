package main

import (
	"log"

	"os"

	"math/rand"
	"time"

	"github.com/0xAX/notificator"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	resetTicker := func() *time.Ticker {
		duration := r.Intn(40-20) + 20 + 1
		return time.NewTicker(time.Duration(duration) * time.Minute)
	}

	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			log.Println("tick at", t)
			f, err := os.Open("./heartbeat-sound-effect/Heartbeat-sound-effect.mp3")
			if err != nil {
				log.Fatal(err)
			}

			streamer, format, err := mp3.Decode(f)
			if err != nil {
				log.Fatal(err)
			}

			sr := format.SampleRate * 2
			speaker.Init(sr, format.SampleRate.N(time.Second/10))

			notify := notificator.New(notificator.Options{
				DefaultIcon: "./stretch.jpg",
				AppName:     "Stretch",
			})

			notify.Push("MOVEMENT!", "If you're sitting, stand.\nIf you're standing, stretch.", "", notificator.UR_CRITICAL)
			speaker.Play(streamer)
			ticker.Stop()
			ticker = resetTicker()
		}
	}
}
