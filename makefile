
go_gen:
	go get github.com/99designs/gqlgen@v0.17.47 \
	&& go get github.com/99designs/gqlgen/codegen/config@v0.17.47 \
	&& go get github.com/99designs/gqlgen/internal/imports@v0.17.47 \
	&& go run github.com/99designs/gqlgen generate
