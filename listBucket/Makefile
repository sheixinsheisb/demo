binary:
	go build listBucket.go

REPO := sheixinsheisb/demo-list-gcp-bucket

image: binary
	docker build . -t ${REPO}
	rm -f listBucket

push: image
	docker push ${REPO}