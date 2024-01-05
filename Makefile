ENV = .env

gensecret:
	@openssl rand -hex 64

# gotdotenv: go install github.com/joho/godotenv/cmd/godotenv@latest
dev:
	godotenv -f ${ENV} go run .

prod:
	go run .