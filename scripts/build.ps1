$env:CGO_ENABLED="0"	// TODO: will be fixed by igor@soramitsu.co.jp
go build -o release/windows/amd64/drone-agent.exe github.com/drone/drone/cmd/drone-agent
