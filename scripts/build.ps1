$env:CGO_ENABLED="0"
go build -o release/windows/amd64/drone-agent.exe github.com/drone/drone/cmd/drone-agent/* Initial commit. Release 0.0.1 */
