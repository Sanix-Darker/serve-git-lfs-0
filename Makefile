build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sglfs

run:
	./sglfs

docker-build:
	docker build -t sanix-darker/sglfs:latest --target prod -f Dockerfile .

docker-build-no-cache:
	docker build --no-cache -t sanix-darker/sglfs:latest --target prod -f Dockerfile .

docker-run:
	docker run -d -it --rm -p 3000:3000 -v ${PWD}/shared:/shared sanix-darker/sglfs:latest
