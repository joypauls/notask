EXECUTABLE_NAME := notask
TAG_COMMIT := $(shell git rev-list --abbrev-commit --tags --max-count=1)
TAG := $(shell git describe --abbrev=0 --tags ${TAG_COMMIT} 2>/dev/null || true)

build:
	go build -ldflags="-s -w -X cmd.root.version=$(TAG)" -o ./bin/${EXECUTABLE_NAME}-darwin-arm64

run:
	go build -o ${EXECUTABLE_NAME}
	# ./${EXECUTABLE_NAME} version
	# # ./${EXECUTABLE_NAME}
	# # ./${EXECUTABLE_NAME} add "This is a new task."
	# echo ""
	./${EXECUTABLE_NAME} board

dev-setup:
	sh ./dev-setup.sh

clean:
	go mod tidy

serve-docs:
	cd docs
	bundle exec jekyll serve

test:
	go test ./... -v

test-coverage:
	go test ./... -coverprofile coverage.out
	go tool cover -func coverage.out | grep total:

# tag:
# 	git tag -a v0.0.0 -m "test tag"
# 	git push --tags
