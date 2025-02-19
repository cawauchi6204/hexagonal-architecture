# GCPプロジェクト設定
PROJECT_ID := your-project-id
REGION := asia-northeast1
DB_INSTANCE := todo-db

.PHONY: setup-gcp
setup-gcp:
	# Artifact Registryの作成
	gcloud artifacts repositories create todo-app \
		--repository-format=docker \
		--location=$(REGION) \
		--description="Todo application container registry"

	# Cloud SQLインスタンスの作成
	gcloud sql instances create $(DB_INSTANCE) \
		--database-version=MYSQL_8_0 \
		--tier=db-f1-micro \
		--region=$(REGION) \
		--root-password=$$DB_ROOT_PASSWORD

	# データベースの作成
	gcloud sql databases create todo \
		--instance=$(DB_INSTANCE)

	# Cloud SQLユーザーの作成
	gcloud sql users create todo-user \
		--instance=$(DB_INSTANCE) \
		--password=$$DB_USER_PASSWORD

.PHONY: setup-secrets
setup-secrets:
	# Secret Managerへの機密情報の登録
	echo -n "$$DB_USER" | gcloud secrets create db-user \
		--data-file=- \
		--replication-policy="automatic"
	
	echo -n "$$DB_PASSWORD" | gcloud secrets create db-password \
		--data-file=- \
		--replication-policy="automatic"
	
	echo -n "$$DB_DATABASE" | gcloud secrets create db-name \
		--data-file=- \
		--replication-policy="automatic"

.PHONY: deploy
deploy:
	# Cloud Buildを使用したデプロイ
	gcloud builds submit \
		--config=deployments/cloudbuild.yaml \
		--substitutions=_REGION=$(REGION),_PROJECT_ID=$(PROJECT_ID)

.PHONY: migrate
migrate:
	# マイグレーションの実行
	cd ../migration && \
	INSTANCE_CONNECTION_NAME=$(PROJECT_ID):$(REGION):$(DB_INSTANCE) \
	DB_USER=$$DB_USER \
	DB_PASSWORD=$$DB_PASSWORD \
	DB_DATABASE=todo \
	./run_migration.sh

.PHONY: setup-all
setup-all: setup-gcp setup-secrets migrate deploy

# ローカル開発環境でCloud SQL Proxyを使用してテストする場合
.PHONY: run-proxy
run-proxy:
	cloud-sql-proxy $(PROJECT_ID):$(REGION):$(DB_INSTANCE)
