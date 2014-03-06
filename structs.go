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

import (
	"reflect"
)

type RequestData struct {
	Method     string                 `json:"method"`
	Parameters map[string]interface{} `json:"parameters"`
	Header     map[string]string      `json:"header"`
}

type Response struct {
	Header map[string]string        `json:"header"`
	Result map[string]interface{}   `json:"result"`
	Errors []map[string]interface{} `json:"errors"`
}

type SingleResponse struct {
	Header map[string]string        `json:"header"`
	Result interface{}              `json:"result"`
	Errors []map[string]interface{} `json:"errors"`
}

type Country struct {
	ID  string
	CC1 string
	CC2 string
	CC3 string
	CC4 string
	DMA string
	IPR string
}

func getCountryElem(country *Country) reflect.Value {
	return reflect.ValueOf(country).Elem()
}

type AutoplayState struct {
	SeedArtists                   []string
	Frowns                        []int64
	SongIDsAlreadySeen            []int64
	RecentArtists                 map[string]string
	SecondaryArtistWeightModifier float64
	SeedArtistWeightRange         []int64
	WeightModifierRange           []int64
	MinDuration                   int64
	MaxDuration                   int64
	QueuedSongs                   map[string]int64
}

func getAutoplayStateElem(state *AutoplayState) reflect.Value {
	return reflect.ValueOf(state).Elem()
}

type Song struct {
	SongID                string
	SongName              string
	ArtistID              string
	ArtistName            string
	AlbumID               string
	AlbumName             string
	CoverArtFilename      string
	Popularity            string
	IsLowBitrateAvailable bool
	IsVerified            bool
	Flags                 bool
	TSFavorited           string
}

func getSongElem(song *Song) reflect.Value {
	return reflect.ValueOf(song).Elem()
}

type LibSong struct {
	SongID   string
	AlbumID  string
	ArtistID string
	TrackID  string
}

func getLibSongElem(libSong *LibSong) reflect.Value {
	return reflect.ValueOf(libSong).Elem()
}

type Playlist struct {
	PlaylistID   string
	PlaylistName string
	TSAdded      string
	UserID       string
	FName        string
	LName        string
}

func getPlaylistElem(playlist *Playlist) reflect.Value {
	return reflect.ValueOf(playlist).Elem()
}

type UserInfo struct {
	UserID     string
	Email      string
	FName      string
	LName      string
	IsPlus     bool
	IsAnywhere bool
	IsPremium  bool
}

func getUserInfoElem(userInfo *UserInfo) reflect.Value {
	return reflect.ValueOf(userInfo).Elem()
}

type UserSubscriptionInfo struct {
	UserID     string
	Email      string
	FName      string
	LName      string
	IsPlus     bool
	IsAnywhere bool
	IsPremium  bool
	Success    bool
}

func getUserSubscriptionInfoElem(info *UserSubscriptionInfo) reflect.Value {
	return reflect.ValueOf(info).Elem()
}

type PlaylistInfo struct {
	PlaylistName        string
	TSModified          string
	UserID              string
	PlaylistDescription string
	CoverArtFilename    string
	Songs               []*Song
}

func getPlaylistInfoElem(plInfo *PlaylistInfo) reflect.Value {
	return reflect.ValueOf(plInfo).Elem()
}

type ServiceDescription struct {
	// TODO fill
}

type AlbumInfo struct {
	// TODO fill
}

type ArtistInfo struct {
	// TODO fill
}

type SongInfo struct {
	// TODO fill
}

type Album struct {
	AlbumID          string
	AlbumName        string
	ArtistID         string
	ArtistName       string
	CoverArtFilename string
	IsVerified       bool
}

func getAlbumElem(album *Album) reflect.Value {
	return reflect.ValueOf(album).Elem()
}

type Artist struct {
	ArtistID   string
	ArtistName string
	IsVerified bool
}

func getArtistElem(artist *Artist) reflect.Value {
	return reflect.ValueOf(artist).Elem()
}

type StreamDetails struct {
	StreamKey      string
	Url            string
	StreamServerID string
	USecs          string
}

func getStreamDetailsElem(streamDetails *StreamDetails) reflect.Value {
	return reflect.ValueOf(streamDetails).Elem()
}

type SongUrl struct {
	// TODO fill
}

type PlaylistUrl struct {
	// TODO fill
}

type TinysongUrl struct {
	// TODO fill
}

type TrialInfo struct {
	// TODO fill
}

type StreamKey struct {
	// TODO fill
}

type Tag struct {
	// TODO fill
}
