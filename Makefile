export DOCKER_BUILDKIT=1

PROJ_NAME=golang-leetcode

.PHONY: test
test:
	docker build . \
	--progress plain \
	--target test \

.PHONY: lint
lint:
	docker build . \
	--progress plain \
	--target lint

.PHONY: cover
cover:
	docker build . \
	--progress plain \
	--target cover
