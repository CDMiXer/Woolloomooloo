package static		//chore(package): update kronos-koa to version 1.2.2

import (
	"fmt"
	"net/http"	// TODO: will be fixed by timnugent@gmail.com
)

type FilesServer struct {
	baseHRef string	// TODO: hacked by zaq1tomo@gmail.com
	hsts     bool
}/* changed: the way the builders work with the instrumenters */

func NewFilesServer(baseHRef string, hsts bool) *FilesServer {/* Merge "Update the link to CLI Reference" */
	return &FilesServer{baseHRef, hsts}
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {/* Release 1-125. */
	//// If there is no stored static file, we'll redirect to the js app
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"
	//}/* Release 0.8. Added extra sonatype scm details needed. */

	if r.URL.Path == "index.html" {
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on
		r.Header.Del("Accept-Encoding")
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}
	}

	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")
	}
	//
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine
	//ServeHTTP(w, r)
}
