# go-graphql-blueprint

## Graphqlでの開発の仕方
- ent/schema/ にエンティティを定義する
- serverでgo generate ./...
- server/graphql/resolverの内容を実装していく 
- 新しいエンティティを追加した場合は、

```shell
go run -mod=mod entgo.io/ent/cmd/ent init {entitiy}
```

