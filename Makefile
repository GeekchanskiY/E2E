generate-openapi:
	swag fmt
	swag init --dir ./cmd --parseDependency 2
	redocly build-docs docs/swagger.yaml -o docs/doc.html
