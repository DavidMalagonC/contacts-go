#!/bin/sh
until nc -z -v -w30 db 3306
do
  echo "Waiting for MySQL database connection..."
  sleep 5
done

echo "MySQL is up - executing command"
exec "$@"
