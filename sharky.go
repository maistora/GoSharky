// This file was auto-generated from the
// Grooveshark API Extractor
// package sharky
package main

import (
	"./struc"
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Get your own KEY and SECTER here http://developers.grooveshark.com/api
const KEY = ""
const SECRET = ""
const API_HOST = "https://api.grooveshark.com"
const API_ENDPOIT = "/ws3.php"
const SIG_GET_KEY = "?sig="

type RequestData struct {
	Method     string            `json:"method"`
	Parameters map[string]string `json:"parameters"`
	Header     map[string]string `json:"header"`
}

type Response struct {
	Header map[string]interface{} `json:"header"`
	Result map[string]interface{} `json:"result"`
}

func generateRequestData(method, sessionID string, params map[string]string) *RequestData {
	data := new(RequestData)
	data.Method = method
	data.Parameters = params

	header := make(map[string]string)
	header["wsKey"] = KEY
	header["sessionID"] = sessionID
	data.Header = header

	return data
}

func generateSignature(postData, secret []byte) string {
	mac := hmac.New(md5.New, secret)
	mac.Write(postData)
	signature := fmt.Sprintf("%x", mac.Sum(nil))

	return signature
}

func generateApiURL(sig string) string {
	return API_HOST + API_ENDPOIT + SIG_GET_KEY + sig
}

func main() {
	params := make(map[string]string)
	reqData := generateRequestData("startSession", "", params)
	buf, _ := json.Marshal(&reqData)
	signature := generateSignature(buf, []byte(SECRET))
	url := generateApiURL(signature)
	fmt.Println(url)
	body := bytes.NewReader(buf)
	r, _ := http.Post(url, "application/json;charset=utf-8", body)
	response, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	var resp Response
	fmt.Println(string(response))
	json.Unmarshal(response, &resp)
	fmt.Println(resp.Header["hostname"])
	fmt.Println(resp.Result["success"])
}

/*
func Request(method, url string, httpBody io.Reader) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, httpBody)
	req.Close = true
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		// whatever
	}
	defer resp.Body.Close()

	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		// Whatever
	}
	fmt.Println(response)
}
*/

type Sharky struct {
	// TODO impl
	// session field is needed
}

// Use addUserLibrarySongsEx instead. Add songs to a user's library.
// Song metadata should be spread across all 3 params. albumIDs[0] should
// be the respective albumID for songIDs[0] and same with artistIDs.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	// TODO impelemnt
}

// Get user library songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserLibrarySongs(limit, page int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Add songs to a user's library. Songs should be an array of objects
// representing each song with keys: songID, albumID, artistID, trackNum.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserLibrarySongsEx(songs string) {
	// TODO impelemnt
}

// Remove songs from a user's library.
// Returns true if everything is OK.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RemoveUserLibrarySongs(songIDs, albumIDs, artistIDs string) bool {
	// TODO impelemnt
	return false
}

// Get subscribed playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserPlaylistsSubscribed() []struc.Playlist {
	// TODO impelemnt
	return nil
}

// Get playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserPlaylists(limit int) []struc.Playlist {
	// TODO impelemnt
	return nil
}

// Get user favorite songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserFavoriteSongs(limit int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Remove a set of favorite songs for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RemoveUserFavoriteSongs(songIDs string) {
	// TODO impelemnt
}

// Logout a user using an established session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) Logout() {
	// TODO impelemnt
}

// Authenticate a user using a token from http://grooveshark.com/auth/.
// See Overview for documentation.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AuthenticateToken(token string) {
	// TODO impelemnt
}

// Get logged-in user info from sessionID
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserInfo() *struc.UserInfo {
	// TODO impelemnt
	return nil
}

// Get logged-in user subscription info. Returns type of subscription
// and either dateEnd or recurring.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetUserSubscriptionDetails() *struc.UserSubscriptionInfo {
	// TODO impelemnt
	return nil
}

// Add a favorite song for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AddUserFavoriteSong(songID int) {
	// TODO impelemnt
}

// Subscribe to a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) SubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Unsubscribe from a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) UnsubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get country from IP. If an IP is omitted, it will use the request's IP.
func (sharky *Sharky) GetCountry(ip string) *struc.Country {
	// TODO impelemnt
	return nil
}

// Get playlist information. To get songs as well, call getPlaylist.
func (sharky *Sharky) GetPlaylistInfo(playlistID string) *struc.PlaylistInfo {
	// TODO impelemnt
	return nil
}

// Get a subset of today's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) GetPopularSongsToday(limit int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func (sharky *Sharky) GetPopularSongsMonth(limit int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Useful for testing if the service is up. Returns "Hello, World" in various languages.
func (sharky *Sharky) PingService() string {
	// TODO impelemnt
	// http.Get()
	return ""
}

// Describe service methods
func (sharky *Sharky) GetServiceDescription() *struc.ServiceDescription {
	// TODO impelemnt
	return nil
}

// Undeletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) UndeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Deletes a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) DeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get songs on a playlist. Use getPlaylist instead.
func (sharky *Sharky) GetPlaylistSongs(playlistID string, limit int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Get playlist info and songs.
func (sharky *Sharky) GetPlaylist(playlistID string, limit int) *struc.Playlist {
	// TODO impelemnt
	return nil
}

// Set playlist songs, overwrites any already saved
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) SetPlaylistSongs(playlistID int, songIDs string) {
	// TODO impelemnt
}

// Create a new playlist, optionally adding songs to it.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) CreatePlaylist(name, songIDs string) {
	// TODO impelemnt
}

// Renames a playlist.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RenamePlaylist(playlistID int, name string) {
	// TODO impelemnt
}

// Authenticate a user using an established session, a login and an md5 of their password.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) Authenticate(login, password string) {
	// TODO impelemnt
}

// Get userID from username
func (sharky *Sharky) GetUserIDFromUsername(username string) string {
	// TODO impelemnt
	return ""
}

// Get meta-data information about one or more albums
func (sharky *Sharky) GetAlbumsInfo(albumIDs string) []struc.AlbumInfo {
	// TODO impelemnt
	return nil
}

// Get songs on an album. Returns all songs, verified and unverified
func (sharky *Sharky) GetAlbumSongs(albumID, limit int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Get meta-data information about one or more artists
func (sharky *Sharky) GetArtistsInfo(artistIDs string) []struc.ArtistInfo {
	// TODO impelemnt
	return nil
}

// Get information about a song or multiple songs.
// The songID(s) should always be passed in as an array.
func (sharky *Sharky) GetSongsInfo(songIDs string) []struc.SongInfo {
	// TODO impelemnt
	return nil
}

// Check if an album exists
func (sharky *Sharky) GetDoesAlbumExist(albumID int) bool {
	// TODO impelemnt
	return false
}

// Check if a song exists
func (sharky *Sharky) GetDoesSongExist(songID int) bool {
	// TODO impelemnt
	return false
}

// Check if an artist exists
func (sharky *Sharky) GetDoesArtistExist(artistID int) bool {
	// TODO impelemnt
	return false
}

// Authenticate a user (login) using an established session.
// Please use the authenticate method instead.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) AuthenticateUser(username, token string) {
	// TODO impelemnt
}

// Get an artist's verified albums
func (sharky *Sharky) GetArtistVerifiedAlbums(artistID int) []struc.Album {
	// TODO impelemnt
	return nil
}

// Get an artist's albums, verified and unverified
func (sharky *Sharky) GetArtistAlbums(artistID int) []struc.Album {
	// TODO impelemnt
	return nil
}

// Get 100 popular songs for an artist
func (sharky *Sharky) GetArtistPopularSongs(artistID int) []struc.Song {
	// TODO impelemnt
	return nil
}

// ================= Search =================

// Perform a playlist search.
func (sharky *Sharky) GetPlaylistSearchResults(query string, limit int) []struc.Playlist {
	// TODO impelemnt
	return nil
}

// Perform an album search.
func (sharky *Sharky) GetAlbumSearchResults(query string, limit int) []struc.Album {
	// TODO impelemnt
	return nil
}

// Perform a song search.
func (sharky *Sharky) GetSongSearchResults(query string, country struc.Country, limit, offset int) []struc.Song {
	// TODO impelemnt
	return nil
}

// Perform an artist search.
func (sharky *Sharky) GetArtistSearchResults(query string, limit int) []struc.Artist {
	// TODO impelemnt
	return nil
}

// ================= Streams =================

// Get stream key, ID, etc. from songID. Requires country object obtained from getCountry
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetStreamKeyStreamServer(songID int, country struc.Country, lowBitrate bool) *struc.StreamDetails {
	// TODO impelemnt
	return nil
}

// ================= URLS =================

// Get Grooveshark URL for tinysong base 62.
func (sharky *Sharky) GetSongURLFromTinysongBase62(base62 string) *struc.SongUrl {
	// TODO impelemnt
	return nil
}

// Get playable song URL from songID
func (sharky *Sharky) GetSongURLFromSongID(songID int) *struc.SongUrl {
	// TODO impelemnt
	return nil
}

// Get playlist URL from playlistID
func (sharky *Sharky) GetPlaylistURLFromPlaylistID(playlistID int) *struc.PlaylistUrl {
	// TODO impelemnt
	return nil
}

// Get a song's Tinysong.com url.
func (sharky *Sharky) GetTinysongURLFromSongID(songID int) *struc.TinysongUrl {
	// TODO impelemnt
	return nil
}

// ================= Users (no auth) =================

// Get playlists created by a userID. Does not require an authenticated session.
func (sharky *Sharky) GetUserPlaylistsByUserID(userID, limit int) []struc.Playlist {
	// TODO impelemnt
	return nil
}

// Get user info from userID
func (sharky *Sharky) GetUserInfoFromUserID(userID int) *struc.UserInfo {
	// TODO impelemnt
	return nil
}

// ================= Recs =================

// Get similar artist for a given artistID.
func (sharky *Sharky) GetSimilarArtists(artistID, limit, page int) []struc.Artist {
	// TODO impelemnt
	return nil
}

// ================= Sessions =================

// Start a session
func (sharky *Sharky) StartSession() {
	// TODO impelemnt

	// {"method":'addUserFavoriteSong",'parameters":{"songID":30547543},"header":{"wsKey":'key","sessionID":'df8fec35811a6b240808563d9f72fa2'}}

	values := make(url.Values)
	values.Set("method", "startSession")
	values.Set("parameters", "{}")
	values.Set("header", "{\"wsKey\":\"golang_nikolay\"}")
	r, err := http.PostForm("http://api.grooveshark.com/ws3.php?sig=3a27a148229e9daceb45e263646b8d8b", values)
	if err != nil {
		fmt.Printf("error posting stat to stathat: %s", err)
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("stathat post result body: %s", body)
	r.Body.Close()

}

// ================= Trials =================

// Gets a trial for an application and the provided uniqueID or logged in user.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetTrialInfo(uniqueID string) *struc.TrialInfo {
	// TODO impelemnt
	return nil
}

// Starts a trial for a user bound to your application and the provided uniqueID.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) CreateTrial(uniqueID string) {
	// TODO impelemnt
}

// ================= Autocomplete =================

// Autocomplete search. Type parameter is 'music', 'playlist', or 'user'. Returns an array of words.
func (sharky *Sharky) GetAutocompleteSearchResults(query, typeParam string, limit int) []string {
	// TODO impelemnt
	return nil
}

// ================= Subscriber streams =================

// Get stream key, ID, etc. from songID for a subscriber account.
// Requires country object obtained from getCountry and a logged-in
// sessionID from a Grooveshark Anywhere subscriber.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) GetSubscriberStreamKey(songID int, country struc.Country, lowBitrate bool, uniqueID string) *struc.StreamKey {
	// TODO impelemnt
	return nil
}

// Mark a song as having been played for greater than or equal to 30 seconds.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) MarkStreamKeyOver30Secs(streamKey string, streamServerID int, uniqueID string) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Mark a song as complete (played for greater than or equal to 30 seconds,
// and having reached the last second either through seeking or normal playback).
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) MarkSongComplete(songID int, streamKey string, streamServerID int, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// ================= Autoplay =================

// Grab a relevant song for autoplay
func (sharky *Sharky) GetAutoplaySong(autoplayState struc.AutoplayState) *struc.Song {
	// TODO impelemnt
	return nil
}

// Gets a list of tags (stations)
func (sharky *Sharky) GetAutoplayTags() []struc.Tag {
	// TODO impelemnt
	return nil
}

// Start autoplay using a tag and grab a relevant song
func (sharky *Sharky) StartAutoplayTag(tagID int) {
	// TODO impelemnt
}

// Start autoplay and grab a relevant song
// TODO check if the params are right
func (sharky *Sharky) StartAutoplay(artistIDs, songIDs []string) {
	// TODO impelemnt
}

// Remove a vote up for a song
func (sharky *Sharky) RemoveVoteUpAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Vote up a song
func (sharky *Sharky) VoteUpAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Remove a song from the autoplay state
func (sharky *Sharky) RemoveSongFromAutoplay(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Add a song to the autoplay state
func (sharky *Sharky) AddSongToAutoplay(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Vote down a song
func (sharky *Sharky) VoteDownAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO imApelemnt
}

// Remove a vote down for a song
func (sharky *Sharky) RemoveVoteDownAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// ================= Tinysong =================

// Get Grooveshark songID for tinysong base 62.
func (sharky *Sharky) GetSongIDFromTinysongBase62(base62 string) string {
	// TODO impelemnt
	return ""
}

// ================= Register =================

// Register and authenticate a user using an established session.
// The username is alpha-numeric with a period, dash or underscore allowed
// in the middle. The username can be blank or 5-32 characters.
// Passwords must be between 5 and 32 characters.
// Note: You must provide a sessionID with this method.
func (sharky *Sharky) RegisterUser(emailAddress, password, fullName, username, gender, birthDate string) {
	// TODO impelemnt
}
