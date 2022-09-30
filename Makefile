export DOCKER_BUILDKIT=1

PROJ_NAME=golang-leetcode

.PHONY: test
test:
	docker build . \
	-t $(PROJ_NAME)-test \
	--progress plain \
	--target test \

.PHONY: lint
lint:
	docker build . \
	-t $(PROJ_NAME)-lint \
	--progress plain \
	--target lint

.PHONY: cover
cover:
	docker build . \
	-t $(PROJ_NAME)-cover \
	--progress plain \
	--target cover

.PHONY: clean
clean:
	bash -c "docker image rm --force $(PROJ_NAME)-{test,lint,cover}"