APP=pagerender
DOCKER_REPO=ronoaldo

build:
	go build -o $(APP) 
	docker build -t $(DOCKER_REPO)/$(APP):latest --build-arg GIT_HASH=$$(git rev-parse --short HEAD) .

run: build
	docker run --rm --env DEV=true -p 8080:8080 -it --name pagerender $(DOCKER_REPO)/$(APP):latest

deploy: build
	docker push $(DOCKER_REPO)/$(APP):latest

clean:
	rm -f pagerender