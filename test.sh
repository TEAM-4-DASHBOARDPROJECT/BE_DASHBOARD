source ./.env
go test ./features/users/... -coverprofile=cover.out && go tool cover -html=cover.out