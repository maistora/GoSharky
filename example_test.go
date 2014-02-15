package sharky_test

import (
	"github.com/maistora/sharky"
	"net/http"
)

// Start the server and Go to your browser
// localhost:8000/sharky will redirect you to
// the music streamed by Grooveshark
func exampleServerStart() {
	http.HandleFunc("/sharky/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", getSongPath())
	w.WriteHeader(303)
}

// This is sample sessionStart, authentication, song retrieve,
// country retrieve (needed for the stream) and
// obtain the stream URL
func getSongPath() string {
	sharky := sharky.New("key", "secret")
	sharky.StartSession()
	sharky.Authenticate("username", "password")
	songs := sharky.GetPopularSongsMonth(2)
	country := sharky.GetCountry("")
	streamDetails := sharky.GetStreamKeyStreamServer(songs[1].SongID, country, false)

	playlists := sharky.GetUserPlaylists(5)
	plSongs := sharky.GetPlaylistSongs("95320696", 100)
	plInfo := sharky.GetPlaylist("95320696", 100)

	fmt.Println(playlists[0])
	fmt.Println(plSongs[0])
	fmt.Printf("\n---- %v", plInfo)
	fmt.Println(plInfo.Songs[0])
	fmt.Println(sharky.PingService())
	fmt.Println(sharky.GetAlbumSearchResults("meteora", 10)[0])
	fmt.Println(sharky.GetSongSearchResults("counting stars", country, 10, 0)[0])

	return streamDetails.Url
}
