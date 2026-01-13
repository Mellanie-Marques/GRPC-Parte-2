#!/usr/bin/env bash
set -euo pipefail
REPO="/mnt/c/Users/mella/Downloads/Pr-tica---Microsservi-os--com--gRPC---Parte-1-master/Pr-tica---Microsservi-os--com--gRPC---Parte-1-master"
cd "$REPO"

# cleanup previous runs: stop lingering go run processes and docker container
if command -v pkill >/dev/null 2>&1; then
  pkill -f '/microservices/payment' || true
  pkill -f '/microservices/order' || true
fi
if docker ps -a --format '{{.Names}}' | grep -q '^ms-mysql$'; then
  docker rm -f ms-mysql >/dev/null || true
fi

# helper to find free port from candidates
find_free_port() {
  local candidates=("$@")
  for p in "${candidates[@]}"; do
    if command -v ss >/dev/null 2>&1; then
      if ss -ltn | awk '{print $4}' | grep -q ":${p}$"; then
        continue
      else
        echo "$p"
        return 0
      fi
    else
      if netstat -ltn 2>/dev/null | awk '{print $4}' | grep -q ":${p}$"; then
        continue
      else
        echo "$p"
        return 0
      fi
    fi
  done
  return 1
}

MYSQL_PORT=$(find_free_port 3308 3309 3310 3311 3312 3313 3314 3315) || { echo "No free mysql port" >&2; exit 1; }
echo "Using host port $MYSQL_PORT for MySQL"

# remove existing container
if docker ps -a --format '{{.Names}}' | grep -q '^ms-mysql$'; then
  docker rm -f ms-mysql >/dev/null || true
fi

# start mysql container
docker run -d --name ms-mysql -p ${MYSQL_PORT}:3306 -e MYSQL_ROOT_PASSWORD=minhasenha mysql:8.0 >/dev/null

# wait for mysql
for i in {1..30}; do
  if docker exec ms-mysql mysql -uroot -pminhasenha -e "SELECT 1" >/dev/null 2>&1; then
    break
  fi
  sleep 1
done

# create databases
printf "CREATE DATABASE IF NOT EXISTS \`order\`; CREATE DATABASE IF NOT EXISTS \`payment\`;\n" | docker exec -i ms-mysql mysql -uroot -pminhasenha

echo "MySQL ready on host port ${MYSQL_PORT}"

# find ports for services
PAYMENT_PORT=$(find_free_port 3001 3002 3003 3004 3005) || { echo "No free payment port" >&2; exit 1; }
ORDER_PORT=$(find_free_port 3000 3006 3007 3008 3009) || { echo "No free order port" >&2; exit 1; }

echo "Using PAYMENT_PORT=${PAYMENT_PORT}, ORDER_PORT=${ORDER_PORT}"

# start payment
cd "$REPO/microservices/payment"
export DB_DRIVER=mysql
export DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:${MYSQL_PORT})/payment"
export APPLICATION_PORT=${PAYMENT_PORT}
export ENV=development
nohup go run cmd/main.go > /tmp/payment.log 2>&1 &
PAYMENT_PID=$!
echo "Started payment (pid=${PAYMENT_PID}), waiting for port ${PAYMENT_PORT}..."
for i in {1..30}; do
  if ss -ltn | awk '{print $4}' | grep -q ":${PAYMENT_PORT}$"; then break; fi
  sleep 1
done

# start order
cd "$REPO/microservices/order"
export PAYMENT_SERVICE_URL=localhost:${PAYMENT_PORT}
export DB_DRIVER=mysql
export DATA_SOURCE_URL="root:minhasenha@tcp(127.0.0.1:${MYSQL_PORT})/order"
export APPLICATION_PORT=${ORDER_PORT}
export ENV=development
nohup go run cmd/main.go > /tmp/order.log 2>&1 &
ORDER_PID=$!
echo "Started order (pid=${ORDER_PID}), waiting for port ${ORDER_PORT}..."
for i in {1..30}; do
  if ss -ltn | awk '{print $4}' | grep -q ":${ORDER_PORT}$"; then break; fi
  sleep 1
done

# run client test
cd "$REPO/microservices/order/client"
set +e
CLIENT_OUTPUT=$(go run main.go 2>&1)
CLIENT_EXIT=$?
set -e

# summary
echo "--- SUMMARY ---"
echo "MYSQL_HOST_PORT=${MYSQL_PORT}"
echo "PAYMENT_PID=${PAYMENT_PID} (logs: /tmp/payment.log)"
echo "ORDER_PID=${ORDER_PID} (logs: /tmp/order.log)"
echo "CLIENT_EXIT=${CLIENT_EXIT}"
echo "CLIENT_OUTPUT:"
echo "$CLIENT_OUTPUT"

echo "Tail payment log (last 30 lines):"
tail -n 30 /tmp/payment.log || true

echo "Tail order log (last 30 lines):"
tail -n 30 /tmp/order.log || true
#!/usr/bin/env bash
set -euo pipefail
REPO="/mnt/c/Users/mella/Downloads/Pr-tica---Microsservi-os--com--gRPC---Parte-1-master/Pr-tica---Microsservi-os--com--gRPC---Parte-1-master"
cd "$REPO"

# helper to find free port from candidates
find_free_port() {
