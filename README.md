# gql

## 構造

```:bash
.
├── README.md
├── go.mod
├── go.sum
├── gqlgen.yml               // 設定ファイル
├── graph
│   ├── generated            // さわらない
│   │   └── generated.go
│   ├── model                // graphqlのモデル
│   │   └── models_gen.go
│   ├── resolver.go          // 型を定義する
│   ├── schema.graphqls      // graphqlのスキーマ定義
│   └── schema.resolvers.go  // schema.graphqlsで生成されたresoloverの実装
└── server.go                // エントリポイント

3 directories, 10 files
```

## 基本的な使い方

```:bash
go mod init github.com/sgash708/gql
go get -u github.com/99designs/gqlgen
go mod tidy
```

### プロジェクト作成

```:bash
go run github.com/99designs/gqlgen init
```

### 定義ファイル変更後対応
`graph/model/models_gen.go`や`graph/schema.resolvers.go`を再生成する。

```:bash
gqlgen
```
