package main

import (
	"github.com/elliotmoore/TimeInc-Radeks-LeavingCard/sarahConnor/clippy"
	"time"
)

func main() {
	clippy.SayGoodbye("It looks like you're trying to write a goodbye letter. Would you like help with that?")
	time.Sleep(3 * time.Second)
	clippy.SayGoodbye("You're not trying to leave are you?")
	time.Sleep(3 * time.Second)
	clippy.SayGoodbye("Wait! You can't leave us! Everything will break!")
	time.Sleep(3 * time.Second)
	clippy.SayGoodbye("*Sniff* Okay, well I guess this is goodbye...")
	time.Sleep(3 * time.Second)
	clippy.SayGoodbye("You'll do amazing things without us then. :'(")
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		clippy.DanceAway()
	}
}
