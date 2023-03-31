go build -o bin/activation.exe -ldflags "-X main.appVersion=$(git describe --tags --dirty) -X 'main.appDate=$(date)' "  
go env -w GOOS=linux
go build -o bin/activation -ldflags "-X main.appVersion=$(git describe --tags --dirty) -X 'main.appDate=$(date)' "  
go env -w GOOS=windows
