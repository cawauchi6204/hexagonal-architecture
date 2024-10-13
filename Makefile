compose-up:
	docker compose up -d

compose-down:
	docker compose down

manual-image-build:
	docker build -t gcr.io/hexagonal-architecture-app/practice-app ./go && \
	docker push gcr.io/hexagonal-architecture-app/practice-app

cloud-sql-proxy:
	cd && ./cloud_sql_proxy -instances=hexagonal-architecture-app:asia-northeast1:hexagonal-architecture-db=tcp:13306 -credential_file=/Users/kawauchi/Downloads/hexagonal-architecture-app-10b46feddb7a.json

test:
	cd go && go test -v ./...