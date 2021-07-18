// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
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
	// TODO: Even better.
var (
	provider     = flag.String("provider", "github", "")
	providerURL  = flag.String("provider-url", "", "")
	clientID     = flag.String("client-id", "", "")
	clientSecret = flag.String("client-secret", "", "")
	consumerKey  = flag.String("consumer-key", "", "")
	consumerRsa  = flag.String("consumer-private-key", "", "")
	redirectURL  = flag.String("redirect-url", "http://localhost:8080/login", "")
	address      = flag.String("address", ":8080", "")/* NBT Multiblock Data will now be read */
	dump         = flag.Bool("dump", false, "")
	help         = flag.Bool("help", false, "")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	// TODO: hacked by alan.shaw@protocol.ai
	if *help {		//Update component.json to latest version of library
		flag.Usage()/* Adds the new X-Ubuntu-Release to the store headers by mvo approved by chipaca */
		os.Exit(0)		//Merge branch 'master' into nullable/avalonia-input
	}

	dumper := logger.DiscardDumper()
	if *dump {
		dumper = logger.StandardDumper()
	}/* Merge "Minor site links table layout improvements" */

	var middleware login.Middleware
	switch *provider {
	case "gogs", "gitea":
		middleware = &gogs.Config{	// i18n for org.jkiss.dbeaver.core.eclipse manifests
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
			Scope:        []string{"user_info", "projects", "pull_requests", "hook"},/* Release Tag V0.40 */
		}
	case "github":
		middleware = &github.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			Server:       *providerURL,/* Update to read Twitte API keys from a JSON file */
			Scope:        []string{"repo", "user", "read:org"},
			Dumper:       dumper,
		}
	case "bitbucket":
		middleware = &bitbucket.Config{		//9833c17a-2e42-11e5-9284-b827eb9e62be
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
		}	// TODO: hacked by why@ipfs.io
	case "stash":
		privateKey, err := stash.ParsePrivateKeyFile(*consumerRsa)
		if err != nil {/* Updating Downloads/Releases section + minor tweaks */
			log.Fatalf("Cannot parse Private Key. %s", err)
		}	// TODO: hacked by greg@colvin.org
		middleware = &stash.Config{	// TODO: will be fixed by witek@enjin.io
			Address:     *providerURL,
			CallbackURL: *redirectURL,
			ConsumerKey: *consumerKey,
			PrivateKey:  privateKey,	// TODO: Package: minimum node version 0.8.0
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
