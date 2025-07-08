build-BadgeFunction:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go && \
	mkdir -p ../.aws-sam/build/BadgeFunction && \
	cp bootstrap ../.aws-sam/build/BadgeFunction/bootstrap