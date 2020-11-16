deploy:
	export GO111MODULE=on && gcloud app deploy app.yaml
	gcloud app logs tail -s default