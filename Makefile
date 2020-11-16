deploy:
	export GO111MODULE=on && gcloud app deploy app.yaml -Y
	gcloud app logs tail -s default