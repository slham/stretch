package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/0xAX/notificator"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func main() {
	min := flag.Int("min", 20, "minimum number of minutes for ticker")
	max := flag.Int("max", 40, "maximum number of minutes for ticker")

	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Seed(time.Now().UnixNano())

	resetTicker := func() *time.Ticker {
		duration := r.Intn(*max-*min) + *min
		log.Println(fmt.Sprintf("ticker set for %d minutes", duration))
		return time.NewTicker(time.Duration(duration) * time.Minute)
	}

	ticker := resetTicker()
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
				DefaultIcon: "./golang.png",
				AppName:     "Stretch",
			})

			notify.Push("MOVEMENT!", "If you're sitting, stand.\nIf you're standing, stretch!", "", notificator.UR_CRITICAL)

			speaker.Play(streamer)
			ticker.Stop()
			ticker = resetTicker()
		}
	}
}
