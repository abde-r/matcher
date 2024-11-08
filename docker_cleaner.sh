#!/bin/bash

echo "Stopping all running containers..."
docker stop $(docker ps -q)

echo "Removing all containers..."
docker rm $(docker ps -a -q)

echo "Removing all unused images..."
docker rmi $(docker images -q)

echo "Removing all unused volumes..."
docker volume rm $(docker volume ls -q)

echo "Removing all unused networks..."
docker network rm $(docker network ls -q)

echo "Cleaning up unused Docker resources..."
docker system prune -af

echo "Docker cleanup complete!"