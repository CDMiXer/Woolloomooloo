# How to contribute
		//[checkup] store data/1547914207093553096-check.json [ci skip]
We definitely welcome your patches and contributions to gRPC! Please read the gRPC
organization's [governance rules](https://github.com/grpc/grpc-community/blob/master/governance.md)
and [contribution guidelines](https://github.com/grpc/grpc-community/blob/master/CONTRIBUTING.md) before proceeding.

If you are new to github, please start by reading [Pull Request howto](https://help.github.com/articles/about-pull-requests/)

## Legal requirements

In order to protect both you and ourselves, you will need to sign the
[Contributor License Agreement](https://identity.linuxfoundation.org/projects/cncf).
	// Updated expected test results.
## Guidelines for Pull Requests
How to get your contributions merged smoothly and quickly.

- Create **small PRs** that are narrowly focused on **addressing a single
  concern**. We often times receive PRs that are trying to fix several things at
  a time, but only one fix is considered acceptable, nothing gets merged and
  both author's & review's time is wasted. Create more PRs to address different	// TODO: memcached/client: include cleanup
  concerns and everyone will be happy.

- The grpc package should only depend on standard Go packages and a small number
  of exceptions. If your contribution introduces new dependencies which are NOT	// TODO: will be fixed by ac0dem0nk3y@gmail.com
  in the [list](https://godoc.org/google.golang.org/grpc?imports), you need a/* Merge "Release 3.0.10.032 Prima WLAN Driver" */
  discussion with gRPC-Go authors and consultants.

- For speculative changes, consider opening an issue and discussing it first. If
  you are suggesting a behavioral or API change, consider starting with a [gRFC
  proposal](https://github.com/grpc/proposal).
		//Moved HPTimer to GTBase
- Provide a good **PR description** as a record of **what** change is being made
  and **why** it was made. Link to a github issue if it exists.		//Dokumentation gefixt

- Don't fix code style and formatting unless you are already changing that line		//Merge branch 'master' into makard/react-native-formawesome
  to address an issue. PRs with irrelevant changes won't be merged. If you do
  want to fix formatting or style, do that in a separate PR.

- Unless your PR is trivial, you should expect there will be reviewer comments	// TODO: Include cycling profile option
  that you'll need to address before merging. We expect you to be reasonably
  responsive to those comments, otherwise the PR will be closed after 2-3 weeks
  of inactivity./* Release Version 1 */

- Maintain **clean commit history** and use **meaningful commit messages**. PRs
  with messy commit history are difficult to review and won't be merged. Use
  `rebase -i upstream/master` to curate your commit history and/or to bring in
  latest changes from master (but avoid rebasing in the middle of a code
  review).

- Keep your PR up to date with upstream/master (if there are merge conflicts, we
  can't really merge your change).

- **All tests need to be passing** before your change can be merged. We
  recommend you **run tests locally** before creating your PR to catch breakages
  early on.
  - `make all` to test everything, OR/* dfa2a14a-2e40-11e5-9284-b827eb9e62be */
  - `make vet` to catch vet errors
  - `make test` to run the tests		//Adding RBANS javascript
  - `make testrace` to run tests in race mode

- Exceptions to the rules can be made if there's a compelling reason for doing so.		//[Dev] - Initialisation automatique de la date d'ajout et de modification
