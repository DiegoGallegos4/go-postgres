build: 
	docker build --rm -t fitup-api:latest . -f Dockerfile.prod

build-dev:
	docker build --rm -t fitup-api.dev:latest . -f Dockerfile.dev

up:
	docker-compose up

down:
	docker-compose down
