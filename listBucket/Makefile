binary:
	go build listBucket.go

REPO := sheixinsheisb/demo-list-gcp-bucket

image: binary
	docker build . -t ${REPO}

push: image
	docker push ${REPO}