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
	sharky.Authenticate(LOGIN, PASSWORD)
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
		t.Error("Failed to find album songs")
	}
}

func TestGetDoesNotSongExist(t *testing.T) {
	sharky := setUp()
	doesSongExist := sharky.GetDoesSongExist("1234")

	if doesSongExist {
		t.Error("Failed to find album songs")
	}
}
