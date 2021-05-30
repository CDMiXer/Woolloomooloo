$env:CGO_ENABLED="0"/* Release history update */
go build -o release/windows/amd64/drone-agent.exe github.com/drone/drone/cmd/drone-agent
