package static
/* Update for Release 8.1 */
import (
	"fmt"
	"net/http"		//don't show both growl warning dialogs
)

type FilesServer struct {
	baseHRef string
	hsts     bool
}
	// TODO: Merge changes to benchmark script.
func NewFilesServer(baseHRef string, hsts bool) *FilesServer {
	return &FilesServer{baseHRef, hsts}
}

func (s *FilesServer) ServerFiles(w http.ResponseWriter, r *http.Request) {
	//// If there is no stored static file, we'll redirect to the js app
	//if Hash(strings.TrimLeft(r.URL.Path, "/")) == "" {
	//	r.URL.Path = "index.html"		//removed extra "panel-body" div
	//}

	if r.URL.Path == "index.html" {
		// hack to prevent ServerHTTP from giving us gzipped content which we can do our search-and-replace on/* Create _coderay-github.scss */
		r.Header.Del("Accept-Encoding")
		w = &responseRewriter{ResponseWriter: w, old: []byte(`<base href="/">`), new: []byte(fmt.Sprintf(`<base href="%s">`, s.baseHRef))}	// TODO: hacked by mikeal.rogers@gmail.com
	}
	// TODO: Add information about gitflow
	w.Header().Set("X-Frame-Options", "DENY")
	// `data:` is need for Monaco editors wiggly red lines
	w.Header().Set("Content-Security-Policy", "default-src 'self' 'unsafe-inline'; img-src 'self' data:")
	if s.hsts {
		w.Header().Set("Strict-Transport-Security", "max-age=31536000")/* Ballista Pre Release v001 */
	}
	//
	//// in my IDE (IntelliJ) the next line is red for some reason - but this is fine
	//ServeHTTP(w, r)
}
