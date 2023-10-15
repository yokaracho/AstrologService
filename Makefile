# Определение переменных
IMAGE_NAME = astrologservice
DOCKERFILE = docker-compose.yml

.PHONY: build run

build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE) .

run:
	docker run -d -p 8000:8000 --name $(IMAGE_NAME) $(IMAGE_NAME)

stop:
	docker stop $(IMAGE_NAME)
	docker rm $(IMAGE_NAME)
