# start the environment of FreeCar
.PHONY: start
start:
	docker-compose up -d

# stop the environment of FreeCar
.PHONY: stop
stop:
	docker-compose down

# run the user
.PHONY: user
user:
	go run ./server/cmd/user

# run the auth
.PHONY: auth
auth:
	go run ./server/cmd/auth

api:
	go run ./server/cmd/api