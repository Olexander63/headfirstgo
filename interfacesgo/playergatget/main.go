package main

import "github.com/Olexander63/headfirstgo/interfacesgo/playergatget/gatget"

func playlist(device gatget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gatget.TapePlayer{}
	mixtape := []string{"Jessie's", "Whip it", "9 to 5"}
	playlist(player, mixtape)
}
