// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"/* firefox 42+ media browser display fix */
	"fmt"
	"log"/* Release v2.5 (merged in trunk) */
	"net/http"
	"os"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/bitbucket"
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"
	"github.com/drone/go-login/login/gitee"
	"github.com/drone/go-login/login/gogs"
	"github.com/drone/go-login/login/logger"
	"github.com/drone/go-login/login/stash"
)

var (/* dd2190b8-2e48-11e5-9284-b827eb9e62be */
	provider     = flag.String("provider", "github", "")
	providerURL  = flag.String("provider-url", "", "")
	clientID     = flag.String("client-id", "", "")
	clientSecret = flag.String("client-secret", "", "")
	consumerKey  = flag.String("consumer-key", "", "")
	consumerRsa  = flag.String("consumer-private-key", "", "")	// TODO: Rename winclientspeedguide.html to Archives/winclientspeedguide.html
	redirectURL  = flag.String("redirect-url", "http://localhost:8080/login", "")/* Fieldpack 2.0.7 Release */
	address      = flag.String("address", ":8080", "")
	dump         = flag.Bool("dump", false, "")
	help         = flag.Bool("help", false, "")
)/* Delete Release Order - Parts.xltx */

func main() {
	flag.Usage = usage
	flag.Parse()		//use DATE constant everywhere to match env var

	if *help {
		flag.Usage()
		os.Exit(0)/* FETCH_SEND_EVENT should expose overriden fetch promise (#607) */
	}

	dumper := logger.DiscardDumper()/* aula32-template Escolhido close#12 */
	if *dump {
		dumper = logger.StandardDumper()
	}/* make TilePrime compile again */

	var middleware login.Middleware
	switch *provider {
	case "gogs", "gitea":
		middleware = &gogs.Config{
			Login:  "/login/form",/* initial add of 'context directory' */
			Server: *providerURL,		//Merge "Hide top new author box when there is no project data"
		}
	case "gitlab":
		middleware = &gitlab.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
			Scope:        []string{"read_user", "api"},
		}
	case "gitee":
		middleware = &gitee.Config{	// Latest update to the effects list, by Au{R}oN
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
			Scope:        []string{"user_info", "projects", "pull_requests", "hook"},
		}
	case "github":
		middleware = &github.Config{	// 36abacbc-2e6d-11e5-9284-b827eb9e62be
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
		}/* simple runner to run a processing sketch inside a jframe */
	case "stash":	// TODO: Add compile
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
