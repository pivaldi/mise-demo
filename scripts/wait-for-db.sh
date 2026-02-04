#!/usr/bin/env bash
set -euo pipefail

host="${POSTGRES_HOST:-127.0.0.1}"
port="${POSTGRES_PORT:-5438}"

echo "[wait-for-db] waiting for postgres at $host:$port ..."
for i in {1..60}; do
    if docker compose exec -T db pg_isready -U "${POSTGRES_USER:-app}" -d "${POSTGRES_DB:-app}" >/dev/null 2>&1; then
        echo "[wait-for-db] postgres is ready"
        exit 0
    fi
    sleep 1
done

echo "[wait-for-db] timeout waiting for postgres"
exit 1
