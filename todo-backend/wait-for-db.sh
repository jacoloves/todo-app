#!/bin/sh

while ! mysqladmin ping -h db --silent; do
	echo "Waiting for databse connection.."
	sleep 2
done

echo "Database is up - executing command"
npm install -g prisma
prisma migrate dev --name init --schema=/prisma/schema.prisma

exec "$@"
