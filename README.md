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
