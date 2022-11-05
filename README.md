# go-graphql-blueprint

# Server
## ent
- ent/schema/ にエンティティを定義する
- serverでgo generate ./...
## protoでの開発の仕方
- 下記の理由があるため、IDでxidを使うのは、grpcにするならやめたほうがいい
- ent/schemaにgrpcのannotationをつけていく形のみ
- annotationを見て勝手に.protoを作成してくるし、server, messageも作成してくれる
```shell
go run -mod=mod entgo.io/ent/cmd/ent init {entitiy}
```
- listのpageTokenがstringになってしまっているので、packageToken := string(bytes)になっている部分を
```
pageToken, err := xid.FromBytes(bytes)
if err != nil {
    return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
}
```

- に置換しないといけない
- forkして修正することも考えられるがあまりやりたくない
- ent/protocmd/main.goで置換している。

## Graphqlでの開発の仕方
- server/graphql/resolverの内容を実装していく


# Front
- WIP