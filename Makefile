build:
	gcloud --project gorvel builds submit --tag asia.gcr.io/gorvel/server

deploy:
	gcloud run deploy gorvel-server --image asia.gcr.io/gorvel/server --project gorvel --region asia-northeast1