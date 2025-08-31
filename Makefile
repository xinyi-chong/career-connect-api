ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "template-single"
DOCKER_NAME = "template-single"

include ./hack/hack.mk

DOCKER_COMPOSE = docker compose -f ./manifest/docker/docker-compose.yaml

dc-up:
	$(DOCKER_COMPOSE) up --build

dc-down:
	$(DOCKER_COMPOSE) down

dc-logs:
	$(DOCKER_COMPOSE) logs -f

dc-ps:
	$(DOCKER_COMPOSE) ps
