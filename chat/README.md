`chat`ディレクトリの直下に`secret.go`を作成し、認証用のIDやキーを次の通り記述すること。

```go:chat.go
package main

const (
	SecurityKey          = "..."
	GoogleClientId       = "..."
	GoogleClientSecret   = "..."
	FacebookClientId     = "..."
	FacebookClientSecret = "..."
	GithubClientId       = "..."
	GithubClientSecret   = "..."
)
```