// Copyright (c) 2013, Nikolay Georgiev
// All rights reserved.

// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:

// * Redistributions of source code must retain the above copyright notice, this
//   list of conditions and the following disclaimer.

// * Redistributions in binary form must reproduce the above copyright notice, this
//   list of conditions and the following disclaimer in the documentation and/or
//   other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
// ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
package sharky_test

import (
	"fmt"
	"github.com/maistora/sharky"
	"net/http"
)

// Start the server and Go to your browser
// localhost:8000/sharky will redirect you to
// the music streamed by Grooveshark
func exampleServerStart() {
	http.HandleFunc("/sharky/1", popularSong)
	http.HandleFunc("/sharky/2", customSongSearch)
	http.ListenAndServe(":8000", nil)
}

func popularSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", getSongStream())
	w.WriteHeader(303)
}

func customSongSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", findSongAndGetStream())
	w.WriteHeader(303)
}

// This is sample sessionStart, authentication, song retrieve,
// country retrieve (needed for the stream) and
// obtain the stream URL
func setUp() *sharky.Sharky {
	sharky := sharky.New("key", "secret")
	sharky.StartSession()
	sharky.Authenticate("username", "password")

	return sharky
}

func getSongStream() string {
	sharky := setUp()
	country := sharky.GetCountry("")
	songs := sharky.GetPopularSongsMonth(2)
	streamDetails := sharky.GetStreamKeyStreamServer(songs[1].SongID, country, false)

	return streamDetails.Url
}

func showPlaylist() {
	sharky := setUp()
	playlists := sharky.GetUserPlaylists(5)
	fmt.Println(playlists[0])
	// Output:
	// &{95320696 test2 2014-02-15 00:35:47}
}

func showPlaylistSongs() {
	sharky := setUp()
	plSongs := sharky.GetPlaylistSongs("95320696", 100)
	fmt.Println(plSongs[0])
	// Output:
	// &{29880235 Numb 671 Linkin Park 113811 Meteora 113811.jpg 1404613534 true false 0 }
}

func showPlaylistInfo() {
	sharky := setUp()
	plInfo := sharky.GetPlaylist("95320696", 100)

	fmt.Println(plInfo)
	fmt.Println(plInfo.Songs[0])
	// Output:
	// &{test2  24922693 testing golang sharky 113811.jpg [0xc210270090 0xc210270120 0xc2102701b0]}
	// &{29880235 Numb 671 Linkin Park 113811 Meteora 113811.jpg 1404613534 true false 0 }
}

func pingService() {
	sharky := setUp()
	fmt.Println(sharky.PingService())
	// Output: Hello world in different languages. Ex.
	// Olá!, mundo.
}

func showAlbumSearchResults() {
	sharky := setUp()
	fmt.Println(sharky.GetAlbumSearchResults("meteora", 10)[0])
	// Output:
	// &{113811 Meteora 671 Linkin Park 113811.jpg true}
}

func findSongAndGetStream() string {
	sharky := setUp()
	country := sharky.GetCountry("") // I know it is already invoked in setUp but...
	song := sharky.GetSongSearchResults("counting stars", country, 10, 0)[0]
	fmt.Println(song)
	// Output:
	// &{38377063 Counting Stars 401901 OneRepublic 8545065 Native 8545065-20140206135006.jpg  true false 0 }

	streamDetails := sharky.GetStreamKeyStreamServer(song.SongID, country, false)

	return streamDetails.Url
}
