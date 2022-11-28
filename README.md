Стенд: https://go-counter-app.herokuapp.com/

```bash
# Запуск через docker-compose
docker-compose build
docker-compose up

# Запуск через docker swarm
docker swarm init
docker stack deploy -c docker-compose.yml go-counter
```