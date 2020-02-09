# 前提

`meander`ディレクトリの直下に`secret.go`を作成し、以下の通りに記述すること。

```go:secret.go
package meander

var APIKey = "Google Places APIのAPI KEY"
```

# テスト

以下のURLを利用することでテストが可能:

```

東京: http://localhost:8080/recommendations?lat=35.681236&lng=139.767125&radius=5000&journey=cafe|bar|casino&cost=$...$$$
大阪: http://localhost:8080/recommendations?lat=34.662778&lng=135.572867&radius=5000&journey=cafe|bar|casino&cost=$...$$$
京都: http://localhost:8080/recommendations?lat=34.932416&lng=135.771056&radius=5000&journey=cafe|bar|casino&cost=$...$$$
```