<?php
header('Content-Type: application/json; charset=utf-8');

echo json_encode([
  "message" => "Hello from PHP",
  "php_version" => PHP_VERSION,
  "app_env" => getenv("APP_ENV") ?: null,
  "database_url" => getenv("DATABASE_URL") ?: null,
], JSON_PRETTY_PRINT);
