APP=pagerender

build:
	go build -o $(APP) 
	docker build -t gcr.io/ronoaldoconsulting/$(APP):latest --build-arg GIT_HASH=$$(git rev-parse --short HEAD) .

run: build
	docker run --rm --env DEV=true -p 8080:8080 -it gcr.io/ronoaldoconsulting/$(APP):latest

deploy: build
	gcloud docker -- push gcr.io/ronoaldoconsulting/$(APP):latest
