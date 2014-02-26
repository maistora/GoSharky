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

const KEY = ""
const SECRET = ""
const LOGIN = ""
const PASSWORD = ""

var shrky *sharky.Sharky = nil

// This is sample sessionStart, authentication, song retrieve,
// country retrieve (needed for the stream) and
// obtain the stream URL
func ExampleSetUp() *sharky.Sharky {
	if shrky != nil {
		return shrky
	}

	shrky = sharky.New(KEY, SECRET)
	shrky.StartSession()
	shrky.Authenticate(LOGIN, PASSWORD)

	return shrky
}

func ExampleSharky_GetUserPlaylists() {
	shrky := ExampleSetUp()
	playlists := shrky.GetUserPlaylists(5)
	fmt.Println(playlists[0].PlaylistID)
	fmt.Println(playlists[0].PlaylistName)
	// Output:
	// 95320696
	// OtherName
}

func ExampleSharky_GetPlaylistSongs() {
	shrky := ExampleSetUp()
	plSongs := shrky.GetPlaylistSongs("95320696", 100)
	fmt.Println(plSongs[0].SongID)
	fmt.Println(plSongs[0].SongName)
	fmt.Println(plSongs[0].ArtistID)
	fmt.Println(plSongs[0].ArtistName)
	fmt.Println(plSongs[0].AlbumID)
	fmt.Println(plSongs[0].AlbumName)
	// Output:
	// 29880235
	// Numb
	// 671
	// Linkin Park
	// 113811
	// Meteora
}

func ExampleSharky_GetPlaylist() {
	shrky := ExampleSetUp()
	plInfo := shrky.GetPlaylist("95320696", 100)

	fmt.Println(plInfo.PlaylistName)
	fmt.Println(plInfo.PlaylistDescription)
	fmt.Println(plInfo.Songs[0].SongID)
	fmt.Println(plInfo.Songs[0].SongName)
	fmt.Println(plInfo.Songs[0].ArtistID)
	fmt.Println(plInfo.Songs[0].ArtistName)
	fmt.Println(plInfo.Songs[0].AlbumID)
	fmt.Println(plInfo.Songs[0].AlbumName)
	// Output:
	// OtherName
	// testing golang sharky
	// 29880235
	// Numb
	// 671
	// Linkin Park
	// 113811
	// Meteora
}

func ExampleSharky_GetAlbumSearchResults() {
	shrky := ExampleSetUp()
	album := shrky.GetAlbumSearchResults("meteora", 10)[0]
	fmt.Println(album.AlbumID)
	fmt.Println(album.AlbumName)
	fmt.Println(album.ArtistName)
	// Output:
	// 113811
	// Meteora
	// Linkin Park
}

func ExampleSharky_GetSongSearchResults() {
	shrky := ExampleSetUp()
	country := shrky.GetCountry("")
	song := shrky.GetSongSearchResults("counting stars", country, 10, 0)[0]
	fmt.Println(song.SongName)
	fmt.Println(song.ArtistName)
	// Output:
	// Counting Stars
	// One Republic
}

func ExampleSharky_GetArtistPopularSongs() {
	shrky := sharky.New(KEY, SECRET)
	song := shrky.GetArtistPopularSongs("2")[0]
	fmt.Println(song.SongName)
	fmt.Println(song.ArtistName)
	// Output:
	// House of Jealous Lovers
	// The Rapture
}

func ExampleSharky_GetPlaylistSearchResult() {
	shrky := ExampleSetUp()
	playlist := shrky.GetPlaylistSearchResults("Meteora", 5)[0]
	fmt.Println(playlist.PlaylistID)
	fmt.Println(playlist.PlaylistName)
	fmt.Println(playlist.TSAdded)
	fmt.Println(playlist.UserID)
	fmt.Println(playlist.FName)
	fmt.Println(playlist.LName)
	// Output:
	// 34658177
	// Meteora
	// 0
	// 5234899
	// Daron
	// Malakian
}

// Start the server and Go to your browser
// localhost:8000/sharky will redirect you to
// the music streamed by Grooveshark
func ExampleServerStart() {
	http.HandleFunc("/sharky/1", ExampleNavigateToSongStream)
	http.HandleFunc("/sharky/2", ExampleNavigateToFoundSong)
	http.ListenAndServe(":8000", nil)
}

func ExampleNavigateToSongStream(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", ExampleSharky_GetStreamKeyStreamServer())
	w.WriteHeader(303)
}

func ExampleNavigateToFoundSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", ExampleSharky_GetSongSearchResults_returnStream())
	w.WriteHeader(303)
}

func ExampleSharky_GetStreamKeyStreamServer() string {
	shrky := ExampleSetUp()
	country := shrky.GetCountry("")
	songs := shrky.GetPopularSongsMonth(2)
	streamDetails := shrky.GetStreamKeyStreamServer(songs[1].SongID, country, false)

	return streamDetails.Url
}

func ExampleSharky_GetSongSearchResults_returnStream() string {
	shrky := ExampleSetUp()
	country := shrky.GetCountry("")
	song := shrky.GetSongSearchResults("counting stars", country, 10, 0)[0]

	streamDetails := shrky.GetStreamKeyStreamServer(song.SongID, country, false)
	return streamDetails.Url
}
