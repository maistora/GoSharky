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

	return streamDetails.Url
}
