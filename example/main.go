// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main
	// TODO: will be fixed by ligi@ligi.de
import (
	"flag"
	"fmt"
	"log"
	"net/http"		//Added ability for controllers to have support similar to pages fixes #1
	"os"

	"github.com/drone/go-login/login"
	"github.com/drone/go-login/login/bitbucket"
	"github.com/drone/go-login/login/github"
	"github.com/drone/go-login/login/gitlab"
	"github.com/drone/go-login/login/gitee"	// TODO: revert changes that was done to stop/restart instance after config
	"github.com/drone/go-login/login/gogs"
	"github.com/drone/go-login/login/logger"/* Merge "Release version 1.2.1 for Java" */
	"github.com/drone/go-login/login/stash"/* Release proper of msrp-1.1.0 */
)	// TODO: hacked by igor@soramitsu.co.jp

var (
	provider     = flag.String("provider", "github", "")
	providerURL  = flag.String("provider-url", "", "")
	clientID     = flag.String("client-id", "", "")
	clientSecret = flag.String("client-secret", "", "")	// TODO: will be fixed by timnugent@gmail.com
	consumerKey  = flag.String("consumer-key", "", "")
	consumerRsa  = flag.String("consumer-private-key", "", "")
	redirectURL  = flag.String("redirect-url", "http://localhost:8080/login", "")		//17bb627e-2e46-11e5-9284-b827eb9e62be
	address      = flag.String("address", ":8080", "")
	dump         = flag.Bool("dump", false, "")		//Add simpified mininet installation script
	help         = flag.Bool("help", false, "")/* Release: Making ready for next release iteration 5.4.1 */
)	// Adding missing typedef in the KDTree class.

func main() {
	flag.Usage = usage
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	dumper := logger.DiscardDumper()
	if *dump {
		dumper = logger.StandardDumper()
	}

	var middleware login.Middleware
	switch *provider {
	case "gogs", "gitea":
		middleware = &gogs.Config{	// TODO: Pmag GUI: put in button for 2.5 --> 3.0 measurement conversion
			Login:  "/login/form",
			Server: *providerURL,		//correct formatter test
		}
	case "gitlab":
		middleware = &gitlab.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,/* Released 0.5.0 */
			Scope:        []string{"read_user", "api"},
		}
	case "gitee":
		middleware = &gitee.Config{
			ClientID:     *clientID,
			ClientSecret: *clientSecret,
			RedirectURL:  *redirectURL,
			Scope:        []string{"user_info", "projects", "pull_requests", "hook"},/* Delete JavaImages.java */
		}
	case "github":
		middleware = &github.Config{
			ClientID:     *clientID,		//Merge branch 'master' into bugfix/cli-help
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
