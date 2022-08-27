package main

import "github.com/Olexander63/headfirstgo/interfacesgo/playergatget/gatget"

type Player interface {
	Play(string)
	Stop()
}

func playlist(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's", "Whip it", "9 to 5"}
	var player Player = gatget.TapePlayer{}
	playlist(player, mixtape)
	player = gatget.TapeRecorder{}
	playlist(player, mixtape)

}
