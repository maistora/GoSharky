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
	"fmt"
	"log"
)

// Get your own KEY and SECTER here http://developers.grooveshark.com/api
const API_HOST = "api.grooveshark.com"
const API_ENDPOIT = "/ws3.php"
const SIG_GET_KEY = "?sig="
const HTTPS = "https://"
const HTTP = "http://"
const CONTENT_TYPE = "application/json;charset=utf-8"

const NO_ACCESS_ERR = "Our web services key does not have access to the invoked method."
const METHOD_NOT_IMPL_ERR = "This method is not yet implemented due to lack of knowledge what is the IDs string format"

func New(key, secret string) *Sharky {
	return new(Sharky).Init(key, secret)
}

// ######################  Sharky's methods  ######################
type Sharky struct {
	SessionID string
	Key       string
	Secret    string
	UserInfo  *UserInfo
}

// Initializes Sharky with key and secret needed for communication with
// GS API and username and password, needed for some specific functionality
func (sharky *Sharky) Init(key, secret string) *Sharky {
	sharky.Key = key
	sharky.Secret = secret

	return sharky
}

func (sharky *Sharky) CallWithHttp(method string, params map[string]interface{}) map[string]interface{} {
	return makeCall(method, params, sharky.SessionID, HTTP, sharky.Key, sharky.Secret)
}

func (sharky *Sharky) CallWithHttps(method string, params map[string]interface{}) map[string]interface{} {
	return makeCall(method, params, sharky.SessionID, HTTPS, sharky.Key, sharky.Secret)
}

func (sharky *Sharky) SingleCallHttp(method string, params map[string]interface{}) interface{} {
	return makeSingleResultCall(method, params, sharky.SessionID, HTTP, sharky.Key, sharky.Secret)
}

// Use addUserLibrarySongsEx instead. Add songs to a user's library.
// Song metadata should be spread across all 3 params. albumIDs[0] should
// be the respective albumID for songIDs[0] and same with artistIDs.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	log.Panic(METHOD_NOT_IMPL_ERR)
}

// Get user library songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserLibrarySongs(limit, page int) []*Song {
	if page <= 0 {
		return nil
	}
	return sharky.getSongs(limit, "getUserLibrarySongs")
}

// Add songs to a user's library. Songs should be an array of objects
// representing each song with keys: songID, albumID, artistID, trackNum.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserLibrarySongsEx(songs []LibSong) {
	// TODO check it
	params := make(map[string]interface{})
	params["songs"] = songs

	sharky.CallWithHttp("addUserLibrarySongsEx", params)
}

// Remove songs from a user's library.
// Returns true if everything is OK.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RemoveUserLibrarySongs(songs []LibSong) bool {
	// TODO check it
	params := make(map[string]interface{})
	params["songs"] = songs

	result := sharky.CallWithHttp("removeUserLibrarySongs", params)

	if suc, ok := result["success"].(bool); ok {
		return suc
	}

	return false
}

// Get subscribed playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserPlaylistsSubscribed() []*Playlist {
	result := sharky.CallWithHttp("getUserPlaylistsSubscribed", nil)
	return sharky.getPlaylists(result)
}

// Get playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserPlaylists(limit int) []*Playlist {
	params := make(map[string]interface{})
	params["limit"] = limit

	result := sharky.CallWithHttp("getUserPlaylists", params)
	return sharky.getPlaylists(result)
}

func (sharky *Sharky) getPlaylists(result map[string]interface{}) []*Playlist {
	if playlists, ok := result["playlists"].([]interface{}); ok {
		plArr := make([]*Playlist, 0)
		for _, plParams := range playlists {
			if plParam, ok := plParams.(map[string]interface{}); ok {
				playlist := new(Playlist)
				elem := getPlaylistElem(playlist)
				mapToStruct(plParam, &elem)
				plArr = append(plArr, playlist)
			}
		}

		return plArr
	}
	return nil
}

// Get user favorite songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserFavoriteSongs(limit int) []*Song {
	return sharky.getSongs(limit, "getUserFavoriteSongs")
}

// Remove a set of favorite songs for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RemoveUserFavoriteSongs(songIDs string) {
	log.Panic(METHOD_NOT_IMPL_ERR)
}

// Logout a user using an established session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) Logout() {
	sharky.CallWithHttp("logout", nil)
	sharky.UserInfo = nil
	sharky.SessionID = ""
	sharky.Key = ""
	sharky.Secret = ""
}

// Authenticate a user using a token from http://grooveshark.com/auth/.
// See Overview for documentation.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AuthenticateToken(token string) {
	params := make(map[string]interface{})
	params["token"] = token

	result := sharky.CallWithHttps("authenticateToken", params)

	sharky.auth(result)
}

// Get logged-in user info from sessionID
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserInfo() *UserInfo {
	if sharky.UserInfo == nil {
		result := sharky.CallWithHttp("getUserInfo", nil)
		userInfo := new(UserInfo)
		elem := getUserInfoElem(userInfo)
		mapToStruct(result, &elem)
		sharky.UserInfo = userInfo
	}
	return sharky.UserInfo
}

// Get logged-in user subscription info. Returns type of subscription
// and either dateEnd or recurring.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserSubscriptionDetails() *UserSubscriptionInfo {
	panic(NO_ACCESS_ERR)
	return nil
}

// Add a favorite song for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserFavoriteSong(songID string) {
	params := make(map[string]interface{})
	params["songID"] = songID

	result := sharky.CallWithHttp("addUserFavoriteSong", params)

	logMsg(result, "Song added to Favorites successfully.", "Error adding song to Favorites.")
}

// Subscribe to a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) SubscribePlaylist(playlistID string) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.CallWithHttp("subscribePlaylist", params)

	sucMsg := "Subscribtion to playlist finished successfully"
	errMsg := "Cannot subscribe to playlist"
	logMsg(result, sucMsg, errMsg)
}

// Unsubscribe from a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) UnsubscribePlaylist(playlistID string) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.CallWithHttp("unsubscribePlaylist", params)

	sucMsg := "Unsubscribed from playlist."
	errMsg := "Cannot upsubscribe from playlist."
	logMsg(result, sucMsg, errMsg)
	// TODO impelemnt
}

// Get country from IP. If an IP is omitted, it will use the request's IP.
func (sharky *Sharky) GetCountry(ip string) *Country {
	params := make(map[string]interface{})
	params["ip"] = ip
	result := sharky.CallWithHttp("getCountry", params)

	country := new(Country)
	elem := getCountryElem(country)
	mapToStruct(result, &elem)

	return country
}

// Get playlist information. To get songs as well, call getPlaylist.
func (sharky *Sharky) GetPlaylistInfo(playlistID string) *PlaylistInfo {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.CallWithHttp("getPlaylistInfo", params)

	playlistInfo := new(PlaylistInfo)
	elem := getPlaylistInfoElem(playlistInfo)
	mapToStruct(result, &elem)

	return playlistInfo
}

// Get a subset of today's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) GetPopularSongsToday(limit int) []*Song {
	return sharky.getSongs(limit, "getPopularSongsToday")
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) GetPopularSongsMonth(limit int) []*Song {
	return sharky.getSongs(limit, "getPopularSongsMonth")
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) getSongs(limit int, method string) []*Song {
	if limit <= 0 {
		return nil
	}
	params := make(map[string]interface{})
	params["limit"] = limit
	result := sharky.CallWithHttp(method, params)

	return sharky.processSongs(result)
}

func (sharky *Sharky) processSongs(result map[string]interface{}) []*Song {
	if result["songs"] == nil {
		result["songs"] = result["Songs"]
	}
	if songs, ok := result["songs"].([]interface{}); ok {
		songArr := make([]*Song, 0)
		for _, songParams := range songs {
			if sParams, ok := songParams.(map[string]interface{}); ok {
				song := new(Song)
				elem := getSongElem(song)
				mapToStruct(sParams, &elem)
				songArr = append(songArr, song)
			}
		}
		return songArr
	}
	return nil
}

// Useful for testing if the service is up. Returns "Hello, World" in various languages.
func (sharky *Sharky) PingService() string {
	result := sharky.SingleCallHttp("pingService", nil)
	if val, ok := result.(string); ok {
		return val
	}
	return ""
}

// Describe service methods
func (sharky *Sharky) GetServiceDescription() *ServiceDescription {
	// TODO impl
	return nil
}

// Undeletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) UndeletePlaylist(playlistID string) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.CallWithHttp("undeletePlaylist", params)
	logMsg(result, "Playlist undelited.", "Cannot undelite playlist.")
}

// Deletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) DeletePlaylist(playlistID string) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	result := sharky.CallWithHttp("deletePlaylist", params)
	logMsg(result, "Playlist deleted.", "Cannot delite playlist.")
}

// Get songs on a playlist. Use getPlaylist instead.
func (sharky *Sharky) GetPlaylistSongs(playlistID string, limit int) []*Song {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	params["limit"] = limit
	result := sharky.CallWithHttp("getPlaylistSongs", params)
	return sharky.processSongs(result)
}

// Get playlist info and songs.
func (sharky *Sharky) GetPlaylist(playlistID string, limit int) *PlaylistInfo {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	params["limit"] = limit
	result := sharky.CallWithHttp("getPlaylist", params)

	playlistInfo := new(PlaylistInfo)
	elem := getPlaylistInfoElem(playlistInfo)
	mapToStruct(result, &elem)
	playlistInfo.Songs = sharky.processSongs(result)
	return playlistInfo
}

// Set playlist songs, overwrites any already saved
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) SetPlaylistSongs(playlistID string, songIDs string) {
	log.Panic(METHOD_NOT_IMPL_ERR)
}

// Create a new playlist, optionally adding songs to it.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) CreatePlaylist(name, songIDs string) {
	log.Panic(METHOD_NOT_IMPL_ERR)
}

// Renames a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RenamePlaylist(playlistID string, name string) {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID
	params["name"] = name

	sharky.CallWithHttp("renamePlaylist", params)
}

// Authenticate a user using an established session, a login and an md5 of their password.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) Authenticate(login, password string) {
	params := make(map[string]interface{})
	params["login"] = login
	params["password"] = md5sum(password)

	result := sharky.CallWithHttps("authenticate", params)

	sharky.auth(result)
}

func (sharky *Sharky) auth(result map[string]interface{}) {
	if suc, ok := result["success"].(bool); ok {
		if suc {
			log.Println("Authentication successful.")
		} else {
			log.Fatal("Invalid credentials. Check your username and password then try again.")
		}
	}
	userInfo := new(UserInfo)
	elem := getUserInfoElem(userInfo)
	mapToStruct(result, &elem)

	sharky.UserInfo = userInfo
}

// Get userID from username
func (sharky *Sharky) GetUserIDFromUsername(username string) string {
	params := make(map[string]interface{})
	params["username"] = username

	result := sharky.CallWithHttp("getUserIDFromUsername", params)
	if val, ok := result["UserID"].(string); ok {
		return val
	}

	if val, ok := result["UserID"].(int64); ok {
		return fmt.Sprintf("%v", val)
	}

	return ""
}

// Get meta-data information about one or more albums
func (sharky *Sharky) GetAlbumsInfo(albumIDs string) []*AlbumInfo {
	log.Panic(METHOD_NOT_IMPL_ERR)
	return nil
}

// Get songs on an album. Returns all songs, verified and unverified
func (sharky *Sharky) GetAlbumSongs(albumID string, limit int) []*Song {
	params := make(map[string]interface{})
	params["albumID"] = albumID
	params["limit"] = limit

	result := sharky.CallWithHttp("getAlbumSongs", params)
	songs := sharky.processSongs(result)
	return songs
}

// Get meta-data information about one or more artists
func (sharky *Sharky) GetArtistsInfo(artistIDs string) []*ArtistInfo {
	log.Panic(METHOD_NOT_IMPL_ERR)
	return nil
}

// Get information about a song or multiple songs.
// The songID(s) should always be passed in as an array.
func (sharky *Sharky) GetSongsInfo(songIDs string) []SongInfo {
	log.Panic(METHOD_NOT_IMPL_ERR)
	return nil
}

// Check if an album exists
func (sharky *Sharky) GetDoesAlbumExist(albumID string) bool {
	params := make(map[string]interface{})
	params["albumID"] = albumID

	return sharky.getSingleBoolResult(params, "getDoesAlbumExist")
}

// Check if a song exists
func (sharky *Sharky) GetDoesSongExist(songID string) bool {
	params := make(map[string]interface{})
	params["songID"] = songID

	return sharky.getSingleBoolResult(params, "getDoesSongExist")
}

// Check if an artist exists
func (sharky *Sharky) GetDoesArtistExist(artistID string) bool {
	params := make(map[string]interface{})
	params["artistID"] = artistID

	return sharky.getSingleBoolResult(params, "getDoesArtistExist")
}

func (sharky *Sharky) getSingleBoolResult(params map[string]interface{}, method string) bool {
	result := sharky.SingleCallHttp(method, params)
	if val, ok := result.(bool); ok {
		return val
	}

	return false
}

// Authenticate a user (login) using an established session.
// Please use the authenticate method instead.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AuthenticateUser(username, token string) {
	panic("Not implemented - better use authenticate method instead.")
}

// Get an artist's verified albums
func (sharky *Sharky) GetArtistVerifiedAlbums(artistID string) []*Album {
	params := make(map[string]interface{})
	params["artistID"] = artistID

	result := sharky.CallWithHttp("getArtistVerifiedAlbums", params)

	return sharky.processAlbums(result)
}

// Get an artist's albums, verified and unverified
func (sharky *Sharky) GetArtistAlbums(artistID string) []*Album {
	params := make(map[string]interface{})
	params["artistID"] = artistID

	result := sharky.CallWithHttp("getArtistAlbums", params)

	return sharky.processAlbums(result)
}

// Get 100 popular songs for an artist
func (sharky *Sharky) GetArtistPopularSongs(artistID string) []*Song {
	params := make(map[string]interface{})
	params["artistID"] = artistID
	result := sharky.CallWithHttp("getArtistPopularSongs", params)
	return sharky.processSongs(result)
}

// ================= Search =================

// Perform a playlist search.
func (sharky *Sharky) GetPlaylistSearchResults(query string, limit int) []*Playlist {
	params := make(map[string]interface{})
	params["query"] = query
	params["limit"] = limit
	result := sharky.CallWithHttp("getPlaylistSearchResults", params)

	return sharky.processPlaylists(result)
}

func (sharky *Sharky) processPlaylists(result map[string]interface{}) []*Playlist {
	if result["playlists"] == nil {
		result["playlists"] = result["Playlists"]
	}
	if playlist, ok := result["playlists"].([]interface{}); ok {
		playlistArr := make([]*Playlist, 0)
		for _, playlistParam := range playlist {
			if plParam, ok := playlistParam.(map[string]interface{}); ok {
				newPlaylist := new(Playlist)
				elem := getPlaylistElem(newPlaylist)
				mapToStruct(plParam, &elem)
				playlistArr = append(playlistArr, newPlaylist)
			}
		}

		return playlistArr
	}

	return nil
}

// Perform an album search.
// This method could also return Paging stats but it is not impelemented
func (sharky *Sharky) GetAlbumSearchResults(query string, limit int) []*Album {
	params := make(map[string]interface{})
	params["query"] = query
	params["limit"] = limit
	result := sharky.CallWithHttp("getAlbumSearchResults", params)

	return sharky.processAlbums(result)
}

func (sharky *Sharky) processAlbums(result map[string]interface{}) []*Album {
	if result["albums"] == nil {
		result["albums"] = result["Albums"]
	}
	if album, ok := result["albums"].([]interface{}); ok {
		albumArr := make([]*Album, 0)
		for _, albumParam := range album {
			if aParams, ok := albumParam.(map[string]interface{}); ok {
				newAlbum := new(Album)
				elem := getAlbumElem(newAlbum)
				mapToStruct(aParams, &elem)
				albumArr = append(albumArr, newAlbum)
			}
		}

		return albumArr
	}

	return nil
}

// Perform a song search.
func (sharky *Sharky) GetSongSearchResults(query string, country *Country, limit, offset int) []*Song {
	params := make(map[string]interface{})
	params["query"] = query
	params["country"] = country
	params["limit"] = limit
	params["offset"] = offset

	result := sharky.CallWithHttp("getSongSearchResults", params)

	return sharky.processSongs(result)
}

// Perform an artist search.
func (sharky *Sharky) GetArtistSearchResults(query string, limit int) []*Artist {
	params := make(map[string]interface{})
	params["query"] = query
	params["limit"] = limit

	result := sharky.CallWithHttp("getArtistSearchResults", params)

	return sharky.processArtists(result)
}

func (sharky *Sharky) processArtists(result map[string]interface{}) []*Artist {
	if result["artists"] == nil {
		result["artists"] = result["Artists"]
	}
	if artist, ok := result["artists"].([]interface{}); ok {
		artistArr := make([]*Artist, 0)
		for _, artistParam := range artist {
			if aParams, ok := artistParam.(map[string]interface{}); ok {
				newArtist := new(Artist)
				elem := getArtistElem(newArtist)
				mapToStruct(aParams, &elem)
				artistArr = append(artistArr, newArtist)
			}
		}
		return artistArr
	}
	return nil
}

// ================= Streams =================

// Get stream key, ID, etc. from songID. Requires country object obtained from getCountry
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetStreamKeyStreamServer(songID string, country *Country, lowBitrate bool) *StreamDetails {
	params := make(map[string]interface{})
	params["songID"] = songID
	params["country"] = country
	params["lowBitrate"] = lowBitrate

	result := sharky.CallWithHttp("getStreamKeyStreamServer", params)

	streamDetails := new(StreamDetails)
	elem := getStreamDetailsElem(streamDetails)
	mapToStruct(result, &elem)

	return streamDetails
}

// ================= URLS =================

// Get Grooveshark URL for tinysong base 62.
func (sharky *Sharky) GetSongURLFromTinysongBase62(base62 string) *SongUrl {
	// TODO impelemnt
	panic("Not implemented")
	return nil
}

// Get playable song URL from songID
func (sharky *Sharky) GetSongURLFromSongID(songID string) string {
	params := make(map[string]interface{})
	params["songID"] = songID

	return sharky.getUrl("getSongURLFromSongID", params)
}

// Get playlist URL from playlistID
func (sharky *Sharky) GetPlaylistURLFromPlaylistID(playlistID string) string {
	params := make(map[string]interface{})
	params["playlistID"] = playlistID

	return sharky.getUrl("getPlaylistURLFromPlaylistID", params)
}

// Get a song's Tinysong.com url.
func (sharky *Sharky) GetTinysongURLFromSongID(songID string) string {
	params := make(map[string]interface{})
	params["songID"] = songID

	return sharky.getUrl("getTinysongURLFromSongID", params)
}

func (sharky *Sharky) getUrl(method string, params map[string]interface{}) string {
	result := sharky.CallWithHttp(method, params)
	if val, ok := result["url"].(string); ok {
		return val
	}
	return ""
}

// ================= Users (no auth) =================

// Get playlists created by a userID. Does not require an authenticated session.
func (sharky *Sharky) GetUserPlaylistsByUserID(userID string, limit int) []*Playlist {
	panic(NO_ACCESS_ERR)
	return nil
}

// Get user info from userID
func (sharky *Sharky) GetUserInfoFromUserID(userID string) *UserInfo {
	panic(NO_ACCESS_ERR)
	return nil
}

// ================= Recs =================

// Get similar artist for a given artistID.
func (sharky *Sharky) GetSimilarArtists(artistID string, limit, page int) []*Artist {
	params := make(map[string]interface{})
	params["artistID"] = artistID
	params["limit"] = limit
	params["page"] = page

	result := sharky.CallWithHttp("GetSimilarArtists", params)

	artistArr := result["artists"]
	if artistArr != nil {
		if val, ok := artistArr.(map[string]interface{}); ok {
			return sharky.processArtists(val)
		}
	}

	return nil
}

// ================= Sessions =================

// Start a session
func (sharky *Sharky) StartSession() {
	result := sharky.CallWithHttps("startSession", nil)
	if val, ok := result["sessionID"].(string); ok {
		sharky.SessionID = val
	}
	if isEmpty(sharky.SessionID) {
		log.Fatalln("Cannot obtain session ID.")
	}
}

// ================= Trials =================

// Gets a trial for an application and the provided uniqueID or logged in user.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetTrialInfo(uniqueID string) *TrialInfo {
	panic(NO_ACCESS_ERR)
	return nil
}

// Starts a trial for a user bound to your application and the provided uniqueID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) CreateTrial(uniqueID string) {
	panic(NO_ACCESS_ERR)
}

// ================= Autocomplete =================

// Autocomplete search. Type parameter is 'music', 'playlist', or 'user'. Returns an array of words.
func (sharky *Sharky) GetAutocompleteSearchResults(query, typeParam string, limit int) []string {
	params := make(map[string]interface{})
	params["query"] = query
	params["type"] = typeParam // music, playlist or user
	params["limit"] = limit

	result := sharky.CallWithHttp("getAutocompleteSearchResults", params)

	values := result["words"]
	if words, ok := values.([]interface{}); ok {
		return toStringArray(words)
	}
	return nil
}

// ================= Subscriber streams =================

// Get stream key, ID, etc. from songID for a subscriber account.
// Requires country object obtained from getCountry and a logged-in
// sessionID from a Grooveshark Anywhere subscriber.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetSubscriberStreamKey(songID string, country Country, lowBitrate bool, uniqueID string) *StreamKey {
	panic(NO_ACCESS_ERR)
	return nil
}

// Mark a song as having been played for greater than or equal to 30 seconds.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) MarkStreamKeyOver30Secs(streamKey string, streamServerID string, uniqueID string) {
	panic(NO_ACCESS_ERR)
}

// ================= Subscriber streams =================

// Mark a song as complete (played for greater than or equal to 30 seconds,
// and having reached the last second either through seeking or normal playback).
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) MarkSongComplete(songID, streamKey string, streamServerID int, autoplayState AutoplayState) {
	// TODO impl when the AutoplayState is known
	panic(METHOD_NOT_IMPL_ERR)
}

// ================= Autoplay =================

// Grab a relevant song for autoplay
func (sharky *Sharky) GetAutoplaySong(autoplayState AutoplayState) *Song {
	// TODO impl when the AutoplayState is known
	panic(METHOD_NOT_IMPL_ERR)
}

// Gets a list of tags (stations)
func (sharky *Sharky) GetAutoplayTags() []*Tag {
	panic(NO_ACCESS_ERR)
}

// Start autoplay using a tag and grab a relevant song
func (sharky *Sharky) StartAutoplayTag(tagID string) {
	panic(NO_ACCESS_ERR)
}

// Start autoplay and grab a relevant song
// TODO check if the params are right
func (sharky *Sharky) StartAutoplay(artistIDs, songIDs []string) {
	panic(NO_ACCESS_ERR)
}

// Remove a vote up for a song
func (sharky *Sharky) RemoveVoteUpAutoplaySong(song Song, autoplayState AutoplayState) {
	panic(NO_ACCESS_ERR)
}

// Vote up a song
func (sharky *Sharky) VoteUpAutoplaySong(song Song, autoplayState AutoplayState) {
	panic(NO_ACCESS_ERR)
}

// Remove a song from the autoplay state
func (sharky *Sharky) RemoveSongFromAutoplay(song Song, autoplayState AutoplayState) {
	panic(NO_ACCESS_ERR)
}

// Add a song to the autoplay state
func (sharky *Sharky) AddSongToAutoplay(song Song, autoplayState AutoplayState) {
	panic(NO_ACCESS_ERR)
}

// Vote down a song
func (sharky *Sharky) VoteDownAutoplaySong(song Song, autoplayState AutoplayState) {
	panic(NO_ACCESS_ERR)
}

// Remove a vote down for a song
func (sharky *Sharky) RemoveVoteDownAutoplaySong(song Song, autoplayState AutoplayState) {
	panic(NO_ACCESS_ERR)
}

// ================= Tinysong =================

// Get Grooveshark songID for tinysong base 62.
func (sharky *Sharky) GetSongIDFromTinysongBase62(base62 string) string {
	panic(NO_ACCESS_ERR)
	return ""
}

// ================= Register =================

// Register and authenticate a user using an established session.
// The username is alpha-numeric with a period, dash or underscore allowed
// in the middle. The username can be blank or 5-32 characters.
// Passwords must be between 5 and 32 characters.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RegisterUser(emailAddress, password, fullName, username, gender, birthDate string) {
	panic(NO_ACCESS_ERR)
}
