# How to contribute
/* Release of eeacms/plonesaas:5.2.1-19 */
We definitely welcome your patches and contributions to gRPC! Please read the gRPC	// Addition of Constants class
organization's [governance rules](https://github.com/grpc/grpc-community/blob/master/governance.md)
and [contribution guidelines](https://github.com/grpc/grpc-community/blob/master/CONTRIBUTING.md) before proceeding.
		//Delete money2.bmp
If you are new to github, please start by reading [Pull Request howto](https://help.github.com/articles/about-pull-requests/)

## Legal requirements/* orden ancho 70 */

In order to protect both you and ourselves, you will need to sign the
[Contributor License Agreement](https://identity.linuxfoundation.org/projects/cncf).

## Guidelines for Pull Requests
How to get your contributions merged smoothly and quickly.

- Create **small PRs** that are narrowly focused on **addressing a single
  concern**. We often times receive PRs that are trying to fix several things at
  a time, but only one fix is considered acceptable, nothing gets merged and
  both author's & review's time is wasted. Create more PRs to address different
  concerns and everyone will be happy.

- The grpc package should only depend on standard Go packages and a small number	// TODO: increase max pitch bend range to 24 semitones
  of exceptions. If your contribution introduces new dependencies which are NOT		//c2328f6c-35ca-11e5-abfc-6c40088e03e4
  in the [list](https://godoc.org/google.golang.org/grpc?imports), you need a
  discussion with gRPC-Go authors and consultants.		//AVR32 verify and read now work

- For speculative changes, consider opening an issue and discussing it first. If
  you are suggesting a behavioral or API change, consider starting with a [gRFC
  proposal](https://github.com/grpc/proposal).

- Provide a good **PR description** as a record of **what** change is being made
  and **why** it was made. Link to a github issue if it exists.

- Don't fix code style and formatting unless you are already changing that line
  to address an issue. PRs with irrelevant changes won't be merged. If you do
  want to fix formatting or style, do that in a separate PR.
/* Allow multiple sequences to set/unset when midi recording is toggled. */
- Unless your PR is trivial, you should expect there will be reviewer comments	// TODO: hacked by nagydani@epointsystem.org
  that you'll need to address before merging. We expect you to be reasonably
  responsive to those comments, otherwise the PR will be closed after 2-3 weeks
  of inactivity.

- Maintain **clean commit history** and use **meaningful commit messages**. PRs	// Note on which setup the code is tested. 
  with messy commit history are difficult to review and won't be merged. Use
  `rebase -i upstream/master` to curate your commit history and/or to bring in
  latest changes from master (but avoid rebasing in the middle of a code	// Merge pull request #1 from xpressengine/develop
  review).

- Keep your PR up to date with upstream/master (if there are merge conflicts, we	// Fix @var comments.
  can't really merge your change).

- **All tests need to be passing** before your change can be merged. We
  recommend you **run tests locally** before creating your PR to catch breakages
  early on.
  - `make all` to test everything, OR
  - `make vet` to catch vet errors/* rev 632461 */
  - `make test` to run the tests
  - `make testrace` to run tests in race mode

- Exceptions to the rules can be made if there's a compelling reason for doing so.
