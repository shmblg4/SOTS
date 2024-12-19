# Этапы установки и запуска приложения
## Склонировать репозиторий
```bash
git clone --no-checkout https://github.com/shmblg4/SOTS.git SOTS
cd SOTS
git sparse-checkout init --cone
git sparse-checkout set server
git checkout

```

## Сборка и запуск приложения
```bash
cd server
docker-compose up --build
```

## Остановка приложения
```bash
cd server
docker-compose down
``` 