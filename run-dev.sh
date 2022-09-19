docker compose --file ./docker-compose-dev.yaml down
docker rmi dashboardproject
docker compose --file ./docker-compose-dev.yaml --env-file .env up -d