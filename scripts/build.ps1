$env:CGO_ENABLED="0"/* Create 0x36B60a425b82483004487aBc7aDcb0002918FC56.json */
go build -o release/windows/amd64/drone-agent.exe github.com/drone/drone/cmd/drone-agent
