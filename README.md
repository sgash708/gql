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

### `graph/model`について
- 自作定義のgoファイルを配置可能。
- `gqlgen`を実行した際に、`schema.graphqls`から「自作定義のgoファイルの不足フィールド」が`schema.resolvers.go`に生成される。
- graphqlの型定義に対応する構造体を管理する。

### `graph/schema.resolvers.go`
- `schema.graphqls`から自動生成される
  - ロジックを実装するところ
- エンドポイントの管理をする
  - handler的な扱いをする

### `graph/resolvers.go`
- 状態管理をする場所
- 依存関係の宣言する場所
- `server.go`で作成する時に初期化する

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
- `graph/model/models_gen.go`や`graph/schema.resolvers.go`を再生成する。
  - `graph/resolver.go`と`server.go`はコマンドで再生成されない。

```:bash
gqlgen
```

### GraphQLクエリの実行

```
mutation createTodo {
  createTodo(input: { text: "todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}

query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}
```
