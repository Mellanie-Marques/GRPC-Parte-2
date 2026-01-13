#!/usr/bin/env bash
set -e
REPO=/mnt/c/Users/mella/Downloads/Pr-tica---Microsservi-os--com--gRPC---Parte-1-master/Pr-tica---Microsservi-os--com--gRPC---Parte-1-master
cd "$REPO"
if docker ps -a --format '{{.Names}}' | grep -q '^ms-mysql$'; then
docker rm -f ms-mysql || true
fi

docker run -d --name ms-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=minhasenha mysql:8.0

# wait for mysql
for i in $(seq 1 20); do
  if docker exec ms-mysql mysql -uroot -pminhasenha -e "SELECT 1" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

# create databases
printf "CREATE DATABASE IF NOT EXISTS \`order\`; CREATE DATABASE IF NOT EXISTS \`payment\`;\n" | docker exec -i ms-mysql mysql -uroot -pminhasenha

echo "MYSQL_OK"
