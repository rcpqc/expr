go test -v -cover -coverprofile=cover -coverpkg=../builtin,../errs,../eval,../types,../
go tool cover -html=cover -o cover.html