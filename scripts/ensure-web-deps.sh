#!/usr/bin/env bash
set -euo pipefail
cd web

if [ -d node_modules ]; then
  exit 0
fi

echo "[web] installing dependencies (npm install)…"
npm install
