// This file was auto-generated from the
// Grooveshark API Extractor
package main

import (
	"./struc"
)

// Use addUserLibrarySongsEx instead. Add songs to a user's library.
// Song metadata should be spread across all 3 params. albumIDs[0] should 
// be the respective albumID for songIDs[0] and same with artistIDs.
// Note: You must provide a sessionID with this method.
func AddUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	// TODO impelemnt
}

// Get user library songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func GetUserLibrarySongs(limit, page int) []struc.Song {
	// TODO impelemnt
}

// Add songs to a user's library. Songs should be an array of objects 
// representing each song with keys: songID, albumID, artistID, trackNum.
// Note: You must provide a sessionID with this method.
func AddUserLibrarySongsEx(songs string) {
	// TODO impelemnt
}

// Remove songs from a user's library.
// Returns true if everything is OK.
// Note: You must provide a sessionID with this method.
func RemoveUserLibrarySongs(songIDs, albumIDs, artistIDs string) bool {
	// TODO impelemnt
}

// Get subscribed playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func GetUserPlaylistsSubscribed() []struc.Playlist {
	// TODO impelemnt
}

// Get playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func GetUserPlaylists(limit int) []struc.Playlist {
	// TODO impelemnt
}

// Get user favorite songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func GetUserFavoriteSongs(limit int) []struc.Song {
	// TODO impelemnt
}

// Remove a set of favorite songs for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func RemoveUserFavoriteSongs(songIDs string) {
	// TODO impelemnt
}

// Logout a user using an established session.
// Note: You must provide a sessionID with this method.
func Logout() {
	// TODO impelemnt
}

// Authenticate a user using a token from http://grooveshark.com/auth/. 
// See Overview for documentation.
// Note: You must provide a sessionID with this method.
func AuthenticateToken(token string) {
	// TODO impelemnt
}

// Get logged-in user info from sessionID
// Note: You must provide a sessionID with this method.
func GetUserInfo() struc.UserInfo {
	// TODO impelemnt
}

// Get logged-in user subscription info. Returns type of subscription
// and either dateEnd or recurring.
// Note: You must provide a sessionID with this method.
func GetUserSubscriptionDetails() struc.UserSubscriptionInfo {
	// TODO impelemnt
}

// Add a favorite song for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func AddUserFavoriteSong(songID int) {
	// TODO impelemnt
}

// Subscribe to a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func SubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Unsubscribe from a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func UnsubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get country from IP. If an IP is omitted, it will use the request's IP.
func GetCountry(ip string) struc.Country {
	// TODO impelemnt
}

// Get playlist information. To get songs as well, call getPlaylist.
func GetPlaylistInfo(playlistID string) struc.PlaylistInfo {
	// TODO impelemnt
}

// Get a subset of today's popular songs, from the Grooveshark popular billboard.
func GetPopularSongsToday(limit int) struc.Song {
	// TODO impelemnt
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func GetPopularSongsMonth(limit int) struc.Song {
	// TODO impelemnt
}

// Useful for testing if the service is up. Returns "Hello, World" in various languages.
func PingService() string {
	// TODO impelemnt
}

// Describe service methods
func GetServiceDescription() struc.ServiceDescription {
	// TODO impelemnt
}

// Undeletes a playlist.
// Note: You must provide a sessionID with this method.
func UndeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Deletes a playlist.
// Note: You must provide a sessionID with this method.
func DeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get songs on a playlist. Use getPlaylist instead.
func GetPlaylistSongs(playlistID string, limit int) []struc.Song {
	// TODO impelemnt
}

// Get playlist info and songs.
func GetPlaylist(playlistID string, limit int) struc.Playlist {
	// TODO impelemnt
}

// Set playlist songs, overwrites any already saved
// Note: You must provide a sessionID with this method.
func SetPlaylistSongs(playlistID int, songIDs string) {
	// TODO impelemnt
}

// Create a new playlist, optionally adding songs to it.
// Note: You must provide a sessionID with this method.
func CreatePlaylist(name, songIDs string) {
	// TODO impelemnt
}

// Renames a playlist.
// Note: You must provide a sessionID with this method.
func RenamePlaylist(playlistID int, name string) {
	// TODO impelemnt
}

// Authenticate a user using an established session, a login and an md5 of their password.
// Note: You must provide a sessionID with this method.
func Authenticate(login, password string) {
	// TODO impelemnt
}

// Get userID from username
func GetUserIDFromUsername(username string) string {
	// TODO impelemnt
}

// Get meta-data information about one or more albums
func GetAlbumsInfo(albumIDs string) []struc.AlbumInfo {
	// TODO impelemnt
}

// Get songs on an album. Returns all songs, verified and unverified
func GetAlbumSongs(albumID, limit int) []struc.Song {
	// TODO impelemnt
}

// Get meta-data information about one or more artists
func GetArtistsInfo(artistIDs string) []struc.ArtistInfo {
	// TODO impelemnt
}

// Get information about a song or multiple songs.
// The songID(s) should always be passed in as an array.
func GetSongsInfo(songIDs string) []struc.SongInfo {
	// TODO impelemnt
}

// Check if an album exists
func GetDoesAlbumExist(albumID int) bool {
	// TODO impelemnt
}

// Check if a song exists
func GetDoesSongExist(songID int) bool {
	// TODO impelemnt
}

// Check if an artist exists
func GetDoesArtistExist(artistID int) bool {
	// TODO impelemnt
}

// Authenticate a user (login) using an established session.
// Please use the authenticate method instead.
// Note: You must provide a sessionID with this method.
func AuthenticateUser(username, token string) {
	// TODO impelemnt
}

// Get an artist's verified albums
func GetArtistVerifiedAlbums(artistID int) []struc.Album {
	// TODO impelemnt
}

// Get an artist's albums, verified and unverified
func GetArtistAlbums(artistID int) []struc.Album {
	// TODO impelemnt
}

// Get 100 popular songs for an artist
func GetArtistPopularSongs(artistID int) []struc.Song {
	// TODO impelemnt
}

// ================= Search =================

// Perform a playlist search.
func GetPlaylistSearchResults(query string, limit int) []struc.Playlist {
	// TODO impelemnt
}

// Perform an album search.
func GetAlbumSearchResults(query string, limit int) []struc.Album {
	// TODO impelemnt
}

// Perform a song search.
func GetSongSearchResults(query string, country struc.Country, limit, offset int) []struc.Song {
	// TODO impelemnt
}

// Perform an artist search.
func GetArtistSearchResults(query string, limit int) []struc.Artist {
	// TODO impelemnt
}

// ================= Streams =================

// Get stream key, ID, etc. from songID. Requires country object obtained from getCountry
// Note: You must provide a sessionID with this method.
func GetStreamKeyStreamServer(songID int, country struc.Country, lowBitrate bool) struc.StreamDetails {
	// TODO impelemnt
}

// ================= URLS =================

// Get Grooveshark URL for tinysong base 62.
func GetSongURLFromTinysongBase62(base62 string) struc.SongUrl {
	// TODO impelemnt
}

// Get playable song URL from songID
func GetSongURLFromSongID(songID int) struc.SongUrl {
	// TODO impelemnt
}

// Get playlist URL from playlistID
func GetPlaylistURLFromPlaylistID(playlistID int) struc.PlaylistUrl {
	// TODO impelemnt
}

// Get a song's Tinysong.com url.
func GetTinysongURLFromSongID(songID int) struc.TinysongUrl {
	// TODO impelemnt
}

// ================= Users (no auth) =================

// Get playlists created by a userID. Does not require an authenticated session.
func GetUserPlaylistsByUserID(userID, limit int) []struc.Playlist {
	// TODO impelemnt
}

// Get user info from userID
func GetUserInfoFromUserID(userID int) struc.UserInfo {
	// TODO impelemnt
}

// ================= Recs =================

// Get similar artist for a given artistID.
func GetSimilarArtists(artistID, limit, page int) []struc.Artist {
	// TODO impelemnt
}

// ================= Sessions =================

// Start a session
func StartSession() {
	// TODO impelemnt
}

// ================= Trials =================

// Gets a trial for an application and the provided uniqueID or logged in user.
// Note: You must provide a sessionID with this method.
func GetTrialInfo(uniqueID string) struc.TrialInfo {
	// TODO impelemnt
}

// Starts a trial for a user bound to your application and the provided uniqueID.
// Note: You must provide a sessionID with this method.
func CreateTrial(uniqueID string) {
	// TODO impelemnt
}

// ================= Autocomplete =================

// Autocomplete search. Type parameter is 'music', 'playlist', or 'user'. Returns an array of words.
func GetAutocompleteSearchResults(query, type string, limit int) []string {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Get stream key, ID, etc. from songID for a subscriber account.
// Requires country object obtained from getCountry and a logged-in
// sessionID from a Grooveshark Anywhere subscriber.
// Note: You must provide a sessionID with this method.
func GetSubscriberStreamKey(songID int, country struc.Country, lowBitrate bool, uniqueID string) struc.StreamKey {
	// TODO impelemnt
}

// Mark a song as having been played for greater than or equal to 30 seconds.
// Note: You must provide a sessionID with this method.
func MarkStreamKeyOver30Secs(streamKey string, streamServerID int, uniqueID string) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Mark a song as complete (played for greater than or equal to 30 seconds,
// and having reached the last second either through seeking or normal playback).
// Note: You must provide a sessionID with this method.
func MarkSongComplete(songID int, streamKey string, streamServerID int, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// ================= Autoplay =================

// Grab a relevant song for autoplay
func GetAutoplaySong(autoplayState struc.AutoplayState) struc.Song {
	// TODO impelemnt
}

// Gets a list of tags (stations)
func GetAutoplayTags() []struc.Tag {
	// TODO impelemnt
}

// Start autoplay using a tag and grab a relevant song
func StartAutoplayTag(tagID int) {
	// TODO impelemnt
}

// Start autoplay and grab a relevant song
// TODO check if the params are right
func StartAutoplay(artistIDs, songIDs []string) {
	// TODO impelemnt
}

// Remove a vote up for a song
func RemoveVoteUpAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Vote up a song
func VoteUpAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Remove a song from the autoplay state
func RemoveSongFromAutoplay(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Add a song to the autoplay state
func AddSongToAutoplay(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// Vote down a song
func VoteDownAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO imApelemnt
}

// Remove a vote down for a song
func RemoveVoteDownAutoplaySong(song struc.Song, autoplayState struc.AutoplayState) {
	// TODO impelemnt
}

// ================= Tinysong =================

// Get Grooveshark songID for tinysong base 62.
func GetSongIDFromTinysongBase62(base62 string) string {
	// TODO impelemnt
}

// ================= Register =================

// Register and authenticate a user using an established session.
// The username is alpha-numeric with a period, dash or underscore allowed
// in the middle. The username can be blank or 5-32 characters.
// Passwords must be between 5 and 32 characters.
// Note: You must provide a sessionID with this method.
func RegisterUser(emailAddress, password, fullName, username, gender, birthDate string) {
	// TODO impelemnt
}
