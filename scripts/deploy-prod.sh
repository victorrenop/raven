export $(grep -v '^#' creds.env | xargs -d '\n')
docker stop raven-prod
docker rm raven-prod
docker rmi $(docker images -a -q raven:prod)
docker pull victorrenop/raven:prod
docker run -d -e TOKEN=$TOKEN_DEV --name last_gas-dev victorrenop/raven:prod
