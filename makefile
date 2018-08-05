all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o meal_team.linux main.go

clean:
	rm -f meal_team.*
