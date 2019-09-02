# dates

## Run application

docker-compose up

before run docker-compose please run 'GO111MODULE=on go mod vendor'

## Use

request with query params:
GET http://localhost:3333/dates?month=3&day=4

or with json (described in task):
GET http://localhost:3333/dates
    Body:
        {
            "month": 3,
            "day": 4
        }