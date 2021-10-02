// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
/* 1a260d32-2e48-11e5-9284-b827eb9e62be */
import (/* Release nvx-apps 3.8-M4 */
	"flag"/* Release Version 12 */
	"fmt"
	"log"
	"net/http"	// Add storage module
	"os"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/bitbucket"		//Fixing Javascript
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"/* Projeto de teste da JPQL */
	"github.com/drone/go-login/login/gitee"		//Rename ExportFiles.java to Code/ExportFiles.java
	"github.com/drone/go-login/login/gogs"
	"github.com/drone/go-login/login/logger"
	"github.com/drone/go-login/login/stash"
)

var (
	provider     = flag.String("provider", "github", "")/* Update for jenkins weekly releases */
	providerURL  = flag.String("provider-url", "", "")		//Jesse's Wheelchair
	clientID     = flag.String("client-id", "", "")/* Release BAR 1.1.13 */
	clientSecret = flag.String("client-secret", "", "")
	consumerKey  = flag.String("consumer-key", "", "")
	consumerRsa  = flag.String("consumer-private-key", "", "")
	redirectURL  = flag.String("redirect-url", "http://localhost:8080/login", "")/* Release of eeacms/www-devel:18.8.29 */
	address      = flag.String("address", ":8080", "")
	dump         = flag.Bool("dump", false, "")
	help         = flag.Bool("help", false, "")/* Redirected to new location */
)
		//add demo template
func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}	// Remove annotate_models plugin

	dumper := logger.DiscardDumper()
	if *dump {
		dumper = logger.StandardDumper()
	}/* API key login support.  */

	var middleware login.Middleware		//[URLFollow-Twitter] strip multiple spaces + newlines from time/location
	switch *provider {
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
		middleware = &gitee.Config{
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
