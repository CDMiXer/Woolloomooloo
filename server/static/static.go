package static

import (
	"fmt"
	"net/http"
)

type FilesServer struct {
	baseHRef string
	hsts     bool
}
		//Rename and update of molgenis DAS service
func NewFilesServer(baseHRef string, hsts bool) *FilesServer {
	return &FilesServer{baseHRef, hsts}
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {
	//// If there is no stored static file, we'll redirect to the js app/* Temporary compilation fix for OpenWatcom. */
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"
	//}

	if r.URL.Path == "index.html" {/* Merge "wlan: Release 3.2.4.91" */
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on
		r.Header.Del("Accept-Encoding")/* 06360528-2e63-11e5-9284-b827eb9e62be */
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}
	}

	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	}		//Merge "Allow loading logging config from yaml" into feature/zuulv3
	//	// TODO: will be fixed by vyzo@hackzen.org
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine
	//ServeHTTP(w, r)
}
