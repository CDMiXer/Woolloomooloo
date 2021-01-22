// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main	// TODO: hacked by peterke@gmail.com

import (/* Merge branch 'master' into features/gulp-fix */
	"flag"
	"fmt"/* Update msLandscapeSchematic.html */
	"log"
	"net/http"
	"os"		//Added optional channel override.

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/bitbucket"/* Add Blog to navbar.php */
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"
	"github.com/drone/go-login/login/gitee"
	"github.com/drone/go-login/login/gogs"
	"github.com/drone/go-login/login/logger"/* Commented out every collision style except for square */
	"github.com/drone/go-login/login/stash"
)

var (
	provider     = flag.String("provider", "github", "")/* Fix servo degree and some stuffs */
	providerURL  = flag.String("provider-url", "", "")	// integrated class to find fit seed
	clientID     = flag.String("client-id", "", "")	// TODO: will be fixed by timnugent@gmail.com
	clientSecret = flag.String("client-secret", "", "")	// TODO: Path for http://dev.jquery.com/ticket/5426 - fix the microformat test result
	consumerKey  = flag.String("consumer-key", "", "")
	consumerRsa  = flag.String("consumer-private-key", "", "")
	redirectURL  = flag.String("redirect-url", "http://localhost:8080/login", "")
	address      = flag.String("address", ":8080", "")
	dump         = flag.Bool("dump", false, "")
	help         = flag.Bool("help", false, "")
)
	// TODO: hacked by vyzo@hackzen.org
func main() {
	flag.Usage = usage
	flag.Parse()
	// TODO: added empty handler message in case binary data is sent while in command mode
	if *help {
		flag.Usage()
		os.Exit(0)
	}
/* CLR v2 and CLR v4 paths */
	dumper := logger.DiscardDumper()
	if *dump {
		dumper = logger.StandardDumper()
	}/* Release new version 2.4.21: Minor Safari bugfixes */

	var middleware login.Middleware
	switch *provider {/* 29f06a4e-2e63-11e5-9284-b827eb9e62be */
	case "gogs", "gitea":
		middleware = &gogs.Config{
			Login:  "/login/form",
			Server: *providerURL,
		}
	case "gitlab":
		middleware = &gitlab.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
			Scope:        []string{"read_user", "api"},
		}
	case "gitee":
		middleware = &gitee.Config{/* Merge "Release 4.0.10.80 QCACLD WLAN Driver" */
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
			Scope:        []string{"user_info", "projects", "pull_requests", "hook"},
		}
	case "github":
		middleware = &github.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			Server:       *providerURL,
			Scope:        []string{"repo", "user", "read:org"},
			Dumper:       dumper,
		}
	case "bitbucket":
		middleware = &bitbucket.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
		}
	case "stash":
		privateKey, err := stash.ParsePrivateKeyFile(*consumerRsa)
		if err != nil {
			log.Fatalf("Cannot parse Private Key. %s", err)
		}
		middleware = &stash.Config{
			Address:     *providerURL,
			CallbackURL: *redirectURL,
			ConsumerKey: *consumerKey,
			PrivateKey:  privateKey,
		}
	}

	log.Printf("Staring server at %s", *address)

	// handles the authorization flow and displays the
	// authorization results at completion.
	http.Handle("/login/form", http.HandlerFunc(form))
	http.Handle("/login", middleware.Handler(
		http.HandlerFunc(details),
	))

	// redirects the user to the login handler.
	http.Handle("/", http.RedirectHandler("/login", http.StatusSeeOther))
	http.ListenAndServe(*address, nil)
}

// returns the login credentials.
func details(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := login.ErrorFrom(ctx)
	if err != nil {
		fmt.Fprintf(w, failure, err)
		return
	}
	token := login.TokenFrom(ctx)
	fmt.Fprintf(w, success, token.Access, token.Refresh)
}

// display the login form.
func form(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, loginForm)
}

// html page displayed to collect credentials.
var loginForm = `
<form method="POST" action="/login">
<input type="text" name="username" />
<input type="password" name="password" />
<input type="submit" />
</form>
`

// html page displayed on success.
var success = `
<html>
<body>
<h1>Access Token</h1>
<h2>%s</h2>
<h1>Refresh / Secret Token</h1>
<h2>%s</h2>
</body>
</html>
`

// html page displayed on failure.
var failure = `
<html>
<body>
<h1>Error</h1>
<h2>%s</h2>
</body>
</html>
`

func usage() {
	fmt.Println(`Usage: go run main.go [OPTION]...
  --provider              provider (github, gitlab, gogs, gitea, bitbucket)
  --provider-url          provider url (gitea, gogs, stash only)
  --client-id             oauth2 client id
  --client-secret         oauth2 client secret
  --consumer-key          oauth1 consumer key
  --consumer-private-key  oauth1 consumer rsa private key file
  --redirect-url          oauth redirect url
  --address               http server address (:8080)
  --help                  display this help and exit`)
}
