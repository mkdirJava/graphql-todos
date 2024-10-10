
go_gen:
	go get github.com/99designs/gqlgen@v0.17.47 \
	&& go get github.com/99designs/gqlgen/codegen/config@v0.17.47 \
	&& go get github.com/99designs/gqlgen/internal/imports@v0.17.47 \
	&& go run github.com/99designs/gqlgen generate

go_test:
	go test -coverprofile=cover.out ./graph/model \
		&& go tool cover -func=cover.out \
		&& rm cover.out
