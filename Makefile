build:
	CGO_ENABLED=0 GOOS=linux go build \
				-a -installsuffix cgo -o sglfs

run:
	./sglfs

docker-build:
	docker build --tag sglfs:latest \
			--target prod -f Dockerfile .

docker-build-no-cache:
	docker build --no-cache \
			--tag sglfs:latest \
			--target prod -f Dockerfile .

docker-run:
	docker run -d --rm -it sglfs:latest \
			-p 3000:3000 --volume ${PWD}/shared/:/shared/
