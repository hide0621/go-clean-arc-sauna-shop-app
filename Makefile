##################
#### Swagger #####
##################
gen-swagger:
	swag init -g app/cmd/main.go  --output app/docs/swagger