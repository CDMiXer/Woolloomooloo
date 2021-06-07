package static

import (
	"fmt"/* remove leftover debug message on client_jwks_uri conf setting */
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
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on	// TODO: gui start panel help added
		r.Header.Del("Accept-Encoding")		//Create tableautemplate.twb
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}
	}

	w.Header().Set("X-Frame-Options", "DENY")/* fix javascript multiple window.onload  */
	// `data:` is need for Monaco editors wiggly red lines
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")/* change dist-upgrade to upgrade (TODO change GTK theme) */
	}
	//
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine/* Release: Making ready for next release cycle 4.1.1 */
	//ServeHTTP(w, r)	// TODO: hacked by sjors@sprovoost.nl
}
