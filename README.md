install cli gomigrate
command
migrate create -ext sql -dir resources/sql  FILE_NAME

migrate -source file://resources/sql -database postgres://root:@localhost:5432/car_pool?sslmode=disable up


install statik file in go binary
statik -f -m -src=resources -dest=lib/ -p files



generate jwt key
ssh-keygen -t rsa -P "" -b 4096 -m PEM -f resources/secret/jwtRS256.key
ssh-keygen -e -m PEM -f resources/secret/jwtRS256.key > resources/secret/jwtRS256.key.pub