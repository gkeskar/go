package main

import (
	"fmt"

	"github.com/headfirstgo/gadget"
)

type Player interface {
	Play(s string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, s := range songs {
		device.Play(s)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {

	var player Player
	player = gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})

}
