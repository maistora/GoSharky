package main

import (
	"goSharky"
	"GoSharky/struc"
)

func main() {
	sharky := new(goSharky.GoSharky)
	songs := sharky.GetPopularSongMonth(10)
	country := sharky.GetCountry("17.235.72.103")
	streamDetails := GetStreamKeyStreamServer(songs[0].id, country, false)

	mp3player := new(3rdPartyMp3Player)
	mp3player.play(streamDetails.stream)
}