# go-graphql-template


## エンジニア向けドキュメント
## 前提
- backendで変更するファイルはすべて[app/](./app/)にあります。
- Makefileも変更可能性があるため同様に[app/](./app/)にあるため、ドキュメントに書かれているmakeコマンドはすべて[app/](./app/)で実行してください。

## 利用技術
| 技術 | 利用しているライブラリやツール等 |
| ---- | ---- |
| 言語 | go1.22.5 |
| ローカル環境構築 | Docker |
| ORM | gorm |
| migration | goose |
| seed | 自前実装 |
| スキーマ | GraphQL |
| GraphQLのcode生成 | gqlgen |

## インデックス
