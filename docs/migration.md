# データベース更新方法(goose を利用)

## migration ファイルの生成

```bash
$ make b
root$ cd ./databases/migrate
root$ goose create ./add_table_worker sql
```

> [!TIP]
> migration の名前は[add or edit _ delete]

## migration の中身を記載

- 20240810085323_add_table_user.sql ファイルが生成されるため、そのファイルに内容を記載。　記載方法は、以下のサイトを参考にする。

## テーブルを更新する

```bash
goose mysql "${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up
```
