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
package sharky

import "testing"

const KEY = ""
const SECRET = ""
const LOGIN = ""
const PASSWORD = ""

func setUp() *Sharky {
	return New(KEY, SECRET)
}

func TestNew(t *testing.T) {
	sharky := setUp()
	if sharky.Key != KEY || sharky.Secret != SECRET {
		t.Error("Creating new Sharky has failed")
	}
}

func TestSessionIDObtain(t *testing.T) {
	sharky := setUp()
	sharky.StartSession()
	if sharky.SessionID == "" {
		t.Error("Failed to obtain SessionID")
	}
}

func TestAuthentication(t *testing.T) {
	sharky := setUp()
	sharky.StartSession()
	sharky.Authenticate(LOGIN, PASSWORD)
	if sharky.UserInfo == nil {
		t.Error("Failed to authenticate")
	}
}

func TestSongSearch(t *testing.T) {
	sharky := setUp()
	sharky.StartSession()
	country := sharky.GetCountry("")
	song := sharky.GetSongSearchResults("counting stars", country, 10, 0)[0]
	if song == nil {
		t.Error("Failed to find song")
	}
	if song.SongName != "Counting Stars" {
		t.Error("Failed to find the right song")
	}
}

func TestGetAlbumSongs(t *testing.T) {
	sharky := setUp()
	songs := sharky.GetAlbumSongs("5462", 10)

	if songs == nil || len(songs) == 0 {
		t.Error("Failed to find album songs")
	}
}

func TestGetDoesAlbumExist(t *testing.T) {
	sharky := setUp()
	doesAlbumExist := sharky.GetDoesAlbumExist("1000")

	if !doesAlbumExist {
		t.Error("Failed to find album songs")
	}
}

func TestGetDoesSongExist(t *testing.T) {
	sharky := setUp()
	doesSongExist := sharky.GetDoesSongExist("123456")

	if !doesSongExist {
		t.Error("Failed to find song but it exists")
	}
}

func TestGetDoesNotSongExist(t *testing.T) {
	sharky := setUp()
	doesSongExist := sharky.GetDoesSongExist("1234")

	if doesSongExist {
		t.Error("Found song but it does not exist")
	}
}

func TestGetDoesArtistExist(t *testing.T) {
	sharky := setUp()
	doesArtistExist := sharky.GetDoesArtistExist("1000")

	if !doesArtistExist {
		t.Error("Did not find artist but it exists.")
	}
}

func TestGetArtistVerifiedAlbums(t *testing.T) {
	sharky := setUp()
	verifiedAlbums := sharky.GetArtistVerifiedAlbums("2")

	if verifiedAlbums == nil || len(verifiedAlbums) == 0 {
		t.Error("Did not find verified albums.")
	}
}

func TestGetArtistVerifiedAlbumsWithoutAlbums(t *testing.T) {
	sharky := setUp()
	verifiedAlbums := sharky.GetArtistVerifiedAlbums("1000")

	if len(verifiedAlbums) != 0 {
		t.Error("Found unxisting verified albums for this artist ID.")
	}
}

func TestGetArtistAlbums(t *testing.T) {
	sharky := setUp()
	albums := sharky.GetArtistAlbums("2")

	if albums == nil || len(albums) == 0 {
		t.Error("Failed to find existing albums for artist.")
	}
}

func TestGetArtistAlbumsWithoutAlbums(t *testing.T) {
	sharky := setUp()
	albums := sharky.GetArtistAlbums("100000")

	if len(albums) != 0 {
		t.Error("Found unexisting albums for artist.")
	}
}

func TestGetArtistPopularSongs(t *testing.T) {
	sharky := setUp()
	songs := sharky.GetArtistPopularSongs("2")

	if songs == nil || len(songs) == 0 {
		t.Error("Failed to find popular songs for artist.")
	}
}

func TestGetArtistPopularSongsWithoutSongs(t *testing.T) {
	sharky := setUp()
	songs := sharky.GetArtistPopularSongs("100000")

	if len(songs) != 0 {
		t.Error("Found unexisting popular songs for artist.")
	}
}

func TestGetPlaylistSearchResults(t *testing.T) {
	sharky := setUp()
	playlists := sharky.GetPlaylistSearchResults("Meteora", 5)

	if playlists == nil || len(playlists) == 0 {
		t.Error("Failed to find playlists.")
	}
}

func TestGetSongURLFromSongID(t *testing.T) {
	sharky := setUp()
	url := sharky.GetSongURLFromSongID("123456")

	if url == "" || url != "http://grooveshark.com/s/Breakin+Up/F71EB?src=3" {
		t.Error("Failed to find song URL from song ID.")
	}
}

func TestGetPlaylistURLFromPlaylistID(t *testing.T) {
	sharky := setUp()
	url := sharky.GetPlaylistURLFromPlaylistID("1880")

	if url == "" || url != "http://listen.grooveshark.com/playlist/~/1880" {
		t.Error("Failed to find song URL from song ID.")
	}
}

func TestGetTinysongURLFromSongID(t *testing.T) {
	sharky := setUp()
	url := sharky.GetTinysongURLFromSongID("123456")

	if url == "" || url != "http://tinysong.com/78KY" {
		t.Error("Failed to find song Tiny URL from song ID.")
	}
}

func TestGetArtistSearchResults(t *testing.T) {
	sharky := setUp()
	artist := sharky.GetArtistSearchResults("Metallica", 10)[0]

	if artist == nil {
		t.Error("Failed to find artist.")
	}

	if artist.ArtistID != "676" {
		t.Error("Failed to find the correct artist.")
	}
}

func TestGetSimilarArtists(t *testing.T) {
	sharky := setUp()
	artists := sharky.GetSimilarArtists("12345", 10, 1)

	if artists == nil {
		t.Error("Failed to find artists")
	}

	if len(artists) == 0 {
		t.Error("Failed to find artists for existing similarities")
	}
}

func TestGetAutocompleteSearchResults(t *testing.T) {
	sharky := setUp()
	words := sharky.GetAutocompleteSearchResults("metall", "music", 10)

	if words == nil || len(words) == 0 {
		t.Error("Failed to find words")
	}
}
