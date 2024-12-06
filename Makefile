run:
	go mod download
	go run main.go migration
	go run main.go healthy

build:
	docker build -t go-healthy-api .

run-test:
	go test ./test -run=^TestArticlesGet$$ -v -count=1
	go test ./test -run=^TestConfigGet$$ -v -count=1
	go test ./test -run=^TestSignUp$$ -v -count=1
	go test ./test -run=^TestSignIn$$ -v -count=1
	
	go test ./test -run=^TestDiaryAdd$$ -v -count=1
	go test ./test -run=^TestDiaryGet$$ -v -count=1
	
	go test ./test -run=^TestEventAdd$$ -v -count=1

	go test ./test -run=^TestEventGet$$ -v -count=1
	go test ./test -run=^TestAchievementGet$$ -v -count=1
	go test ./test -run=^TestGraphGet$$ -v -count=1
