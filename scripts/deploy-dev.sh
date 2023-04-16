export $(grep -v '^#' creds.env | xargs -d '\n')
docker stop raven-develop
docker rm raven-develop
docker rmi $(docker images -a -q raven:develop)
docker pull victorrenop/raven:develop
docker run -d -e TOKEN=$TOKEN_DEV --name last_gas-dev victorrenop/raven:develop
