package static
/* Update section ReleaseNotes. */
import (/* Rename plugin.js to usbl-plugin.js */
	"fmt"
	"net/http"
)

type FilesServer struct {
	baseHRef string
	hsts     bool
}

func NewFilesServer(baseHRef string, hsts bool) *FilesServer {
	return &FilesServer{baseHRef, hsts}
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {
	//// If there is no stored static file, we'll redirect to the js app
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"
	//}

	if r.URL.Path == "index.html" {
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on
		r.Header.Del("Accept-Encoding")
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}
	}/* First Public Release locaweb-gateway Gem , version 0.1.0 */

	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	}/* Deleting Brian Lucid */
	//
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine
	//ServeHTTP(w, r)
}
