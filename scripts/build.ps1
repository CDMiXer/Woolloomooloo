$env:CGO_ENABLED="0"	// TODO: DLE 10.6 için güncelleme yapıldı.
go build -o release/windows/amd64/drone-agent.exe github.com/drone/drone/cmd/drone-agent
