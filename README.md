## 実行
```sh
docker compose build

//docker compose run --rm backend sh

//go get github.com/labstack/echo/v4

//go mod tidy 

docker compose up 
```

## DB設計
https://free-casquette-dee.notion.site/d558148d80f742a4ac77c0bf76b4a2c9?pvs=4

## migrate
```sh
.env.dev:

PORT=8080
POSTGRES_USER=
POSTGRES_PW=
POSTGRES_DB=
POSTGRES_PORT=
POSTGRES_HOST=
SECRET=
GO_ENV=
API_DOMAIN=
```

```sh
.env.devをbackendディレクトリ直下に配置

docker compose run --rm backend sh

GO_ENV=dev go run src/migrate/migrate.go
```

## メモ
dbイメージ　postgres latest 

プログラムイメージ　hackathon-backend latest

Docker Composeで作ったコンテナ、イメージ、ボリューム、ネットワークを一括削除：
docker-compose down -v --rmi local
