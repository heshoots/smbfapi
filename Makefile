clean:
		rm -r gen/*
generate: 
		swagger generate server --spec=./api/swagger.yaml --exclude-main --target=gen
build:
		CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o smbfapi cmd/server/main.go

