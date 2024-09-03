# Makefile

# Build the Docker image for the ML API
docker-build:
	docker build -t file-registry .

docker-run:
	docker run -p 9000:8000 file-registry
