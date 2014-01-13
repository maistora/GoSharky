// This file was auto-generated from the
// Grooveshark API Extractor
package main


// Use addUserLibrarySongsEx instead. Add songs to a user's library.
// Song metadata should be spread across all 3 params. albumIDs[0] should 
// be the respective albumID for songIDs[0] and same with artistIDs.
// Note: You must provide a sessionID with this method.
func addUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	// TODO impelemnt
}

// Get user library songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func getUserLibrarySongs(limit, page int) {
	// TODO impelemnt
}

// Add songs to a user's library. Songs should be an array of objects 
// representing each song with keys: songID, albumID, artistID, trackNum.
// Note: You must provide a sessionID with this method.
func addUserLibrarySongsEx(songs string) {
	// TODO impelemnt
}

// Remove songs from a user's library.
// Note: You must provide a sessionID with this method.
func removeUserLibrarySongs(songIDs, albumIDs, artistIDs string) {
	// TODO impelemnt
}

// Get subscribed playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func getUserPlaylistsSubscribed() {
	// TODO impelemnt
}

// Get playlists of the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func getUserPlaylists(limit int) {
	// TODO impelemnt
}

// Get user favorite songs. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func getUserFavoriteSongs(limit int) {
	// TODO impelemnt
}

// Remove a set of favorite songs for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func removeUserFavoriteSongs(songIDs string) {
	// TODO impelemnt
}

// Logout a user using an established session.
// Note: You must provide a sessionID with this method.
func logout() {
	// TODO impelemnt
}

// Authenticate a user using a token from http://grooveshark.com/auth/. 
// See Overview for documentation.
// Note: You must provide a sessionID with this method.
func authenticateToken(token string) {
	// TODO impelemnt
}

// Get logged-in user info from sessionID
// Note: You must provide a sessionID with this method.
func getUserInfo() {
	// TODO impelemnt
}

// Get logged-in user subscription info. Returns type of subscription
// and either dateEnd or recurring.
// Note: You must provide a sessionID with this method.
func getUserSubscriptionDetails() {
	// TODO impelemnt
}

// Add a favorite song for a user. Must provide a logged-in sessionID.
// Note: You must provide a sessionID with this method.
func addUserFavoriteSong(songID int) {
	// TODO impelemnt
}

// Subscribe to a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func subscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Unsubscribe from a playlist for the logged-in user. Requires an authenticated session.
// Note: You must provide a sessionID with this method.
func unsubscribePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get country from IP. If an IP is omitted, it will use the request's IP.
func getCountry(ip string) {
	// TODO impelemnt
}

// Get playlist information. To get songs as well, call getPlaylist.
func getPlaylistInfo(playlistID string) {
	// TODO impelemnt
}

// Get a subset of today's popular songs, from the Grooveshark popular billboard.
func getPopularSongsToday(limit int) {
	// TODO impelemnt
}

// Get a subset of this month's popular songs, from the Grooveshark popular billboard.
func getPopularSongsMonth(limit int) {
	// TODO impelemnt
}

// Useful for testing if the service is up. Returns "Hello, World" in various languages.
func pingService() {
	// TODO impelemnt
}

// Describe service methods
func getServiceDescription() {
	// TODO impelemnt
}

// Undeletes a playlist.
// Note: You must provide a sessionID with this method.
func undeletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Deletes a playlist.
// Note: You must provide a sessionID with this method.
func deletePlaylist(playlistID int) {
	// TODO impelemnt
}

// Get songs on a playlist. Use getPlaylist instead.
func getPlaylistSongs(playlistID string, limit int) {
	// TODO impelemnt
}

// Get playlist info and songs.
func getPlaylist(playlistID string, limit int) {
	// TODO impelemnt
}

// Set playlist songs, overwrites any already saved
// Note: You must provide a sessionID with this method.
func setPlaylistSongs(playlistID int, songIDs string) {
	// TODO impelemnt
}

// Create a new playlist, optionally adding songs to it.
// Note: You must provide a sessionID with this method.
func createPlaylist(name, songIDs string) {
	// TODO impelemnt
}

// Renames a playlist.
// Note: You must provide a sessionID with this method.
func renamePlaylist(playlistID int, name string) {
	// TODO impelemnt
}

// Authenticate a user using an established session, a login and an md5 of their password.
// Note: You must provide a sessionID with this method.
func authenticate(login, password string) {
	// TODO impelemnt
}

// Get userID from username
func getUserIDFromUsername(username string) {
	// TODO impelemnt
}

// Get meta-data information about one or more albums
func getAlbumsInfo(albumIDs string) {
	// TODO impelemnt
}

// Get songs on an album. Returns all songs, verified and unverified
func getAlbumSongs(albumID, limit int) {
	// TODO impelemnt
}

// Get meta-data information about one or more artists
func getArtistsInfo(artistIDs string) {
	// TODO impelemnt
}

// Get information about a song or multiple songs.
// The songID(s) should always be passed in as an array.
func getSongsInfo(songIDs string) {
	// TODO impelemnt
}

// Check if an album exists
func getDoesAlbumExist(albumID int) {
	// TODO impelemnt
}

// Check if a song exists
func getDoesSongExist(songID int) {
	// TODO impelemnt
}

// Check if an artist exists
func getDoesArtistExist(artistID int) {
	// TODO impelemnt
}

// Authenticate a user (login) using an established session.
// Please use the authenticate method instead.
// Note: You must provide a sessionID with this method.
func authenticateUser(username, token string) {
	// TODO impelemnt
}

// Get an artist's verified albums
func getArtistVerifiedAlbums(artistID int) {
	// TODO impelemnt
}

// Get an artist's albums, verified and unverified
func getArtistAlbums(artistID int) {
	// TODO impelemnt
}

// Get 100 popular songs for an artist
func getArtistPopularSongs(artistID int) {
	// TODO impelemnt
}

// ================= Search =================

// Perform a playlist search.
func getPlaylistSearchResults(query string, limit int) {
	// TODO impelemnt
}

// Perform an album search.
func getAlbumSearchResults(query string, limit int) {
	// TODO impelemnt
}

// Perform a song search.
func getSongSearchResults(query string, country object, limit, offset int) {
	// TODO impelemnt
}

// Perform an artist search.
func getArtistSearchResults(query string, limit int) {
	// TODO impelemnt
}

// ================= Streams =================

// Get stream key, ID, etc. from songID. Requires country object obtained from getCountry
// Note: You must provide a sessionID with this method.
func getStreamKeyStreamServer(songID int, country object, lowBitrate bool) {
	// TODO impelemnt
}

// ================= URLS =================

// Get Grooveshark URL for tinysong base 62.
func getSongURLFromTinysongBase62(base62 string) {
	// TODO impelemnt
}

// Get playable song URL from songID
func getSongURLFromSongID(songID int) {
	// TODO impelemnt
}

// Get playlist URL from playlistID
func getPlaylistURLFromPlaylistID(playlistID int) {
	// TODO impelemnt
}

// Get a song's Tinysong.com url.
func getTinysongURLFromSongID(songID int) {
	// TODO impelemnt
}

// ================= Users (no auth) =================

// Get playlists created by a userID. Does not require an authenticated session.
func getUserPlaylistsByUserID(userID, limit int) {
	// TODO impelemnt
}

// Get user info from userID
func getUserInfoFromUserID(userID int) {
	// TODO impelemnt
}

// ================= Recs =================

// Get similar artist for a given artistID.
func getSimilarArtists(artistID, limit, page int) {
	// TODO impelemnt
}

// ================= Sessions =================

// Start a session
func startSession() {
	// TODO impelemnt
}

// ================= Trials =================

// Gets a trial for an application and the provided uniqueID or logged in user.
// Note: You must provide a sessionID with this method.
func getTrialInfo(uniqueID string) {
	// TODO impelemnt
}

// Starts a trial for a user bound to your application and the provided uniqueID.
// Note: You must provide a sessionID with this method.
func createTrial(uniqueID string) {
	// TODO impelemnt
}

// ================= Autocomplete =================

// Autocomplete search. Type parameter is 'music', 'playlist', or 'user'. Returns an array of words.
func getAutocompleteSearchResults(query, type string, limit int) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Get stream key, ID, etc. from songID for a subscriber account.
// Requires country object obtained from getCountry and a logged-in
// sessionID from a Grooveshark Anywhere subscriber.
// Note: You must provide a sessionID with this method.
func getSubscriberStreamKey(songID int, country object, lowBitrate bool, uniqueID string) {
	// TODO impelemnt
}

// Mark a song as having been played for greater than or equal to 30 seconds.
// Note: You must provide a sessionID with this method.
func markStreamKeyOver30Secs(streamKey string, streamServerID int, uniqueID string) {
	// TODO impelemnt
}

// ================= Subscriber streams =================

// Mark a song as complete (played for greater than or equal to 30 seconds,
// and having reached the last second either through seeking or normal playback).
// Note: You must provide a sessionID with this method.
func markSongComplete(songID int, streamKey string, streamServerID int, autoplayState object) {
	// TODO impelemnt
}

// ================= Autoplay =================

// Grab a relevant song for autoplay
func getAutoplaySong(autoplayState object) {
	// TODO impelemnt
}

// Gets a list of tags (stations)
func getAutoplayTags() {
	// TODO impelemnt
}

// Start autoplay using a tag and grab a relevant song
func startAutoplayTag(tagID int) {
	// TODO impelemnt
}

// Start autoplay and grab a relevant song
func startAutoplay(artistIDs, songIDs object) {
	// TODO impelemnt
}

// Remove a vote up for a song
func removeVoteUpAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// Vote up a song
func voteUpAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// Remove a song from the autoplay state
func removeSongFromAutoplay(song, autoplayState object) {
	// TODO impelemnt
}

// Add a song to the autoplay state
func addSongToAutoplay(song, autoplayState object) {
	// TODO impelemnt
}

// Vote down a song
func voteDownAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// Remove a vote down for a song
func removeVoteDownAutoplaySong(song, autoplayState object) {
	// TODO impelemnt
}

// ================= Tinysong =================

// Get Grooveshark songID for tinysong base 62.
func getSongIDFromTinysongBase62(base62 string) {
	// TODO impelemnt
}

// ================= Register =================

// Register and authenticate a user using an established session.
// The username is alpha-numeric with a period, dash or underscore allowed
// in the middle. The username can be blank or 5-32 characters.
// Passwords must be between 5 and 32 characters.
// Note: You must provide a sessionID with this method.
func registerUser(emailAddress, password, fullName, username, gender, birthDate string) {
	// TODO impelemnt
}
