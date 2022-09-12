set OkamiConfig=D:\okami-project\backend\src\okami.auth.backend\config\
set OAUTH_HOST=0.0.0.0
set OAUTH_PORT=7000
set OAUTH_RESOURCE_ID=auth
set OAUTH_DB_CONNECTION=user=postgres password=bg1603 dbname=okami sslmode=disable host=localhost port=5432
set OAUTH_DB_SCHEMA=okami.auth
set OAUTH_DB_VIEW_CONNECTION=user=postgres password=bg1603 dbname=okami sslmode=disable host=localhost port=5432
set OAUTH_DB_VIEW_SCHEMA=okami.auth
set OAUTH_REDIS_HOST=localhost
set OAUTH_REDIS_PORT=6379
set OAUTH_REDIS_DB=7
@REM set OAUTH_REDIS_PASSWORD=
@REM set OAUTH_CLIENT_SECRET=4a99014348424f9ca38fe6a543887164
@REM set OAUTH_SIGNATURE_KEY=43d0dac2ec2e435eaa036987abbdbe9c
set OAUTH_AUTH_USERID=1
set OAUTH_CLIENTID=f23992cfdef34e1f9fcdd441d27d5cb7
set OAUTH_JWT_KEY=60651914d7154c12a7d345a5ca2e8722
set OAUTH_INTERNAL_KEY=77b15103809b4005a96cbec02f73bd56
@REM set OAUTH_KAFKA=localhost:9092,localhost:9091
@REM set OAUTH_KAFKA=10.10.0.194:9092

go run main.go development