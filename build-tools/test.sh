#!/usr/bin/env bash

current_directory="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"

project_name="ns-authentication"
docker_network="ns-authentication_default"
docker_min_version="18.06.0"
docker_compose_min_version="1.23.1"
docker_runner_image="microlayers/golang-with-extras:1.11.2-alpine3.8"

function checkBinary {
	local search=$1
	local version=$2

	binary=$(which $search 2>/dev/null)
	if [ "$binary" == "" ]; then
		echo "Missing $search, please install it (minimum version: $version)"

		exit 1
	fi
}

function checkDocker {
	checkBinary docker $docker_min_version
	checkBinary docker-compose $docker_compose_min_version

	docker_client_version=$(docker version | grep -A5 Client | grep Version | awk '{print $2}')
	docker_engine_version=$(docker version | grep -A5 Engine | grep Version | awk '{print $2}')
	docker_compose_version=$(docker-compose version | grep "docker-compose version" | awk '{print $3}' | sed 's/,//')

	echo "docker (client): $docker_client_version"
	echo "docker (engine): $docker_engine_version"
	echo "docker-compose: $docker_compose_version"
}

function startServices {
	echo "Starting up services"
	cd $current_directory
	docker-compose \
		--project-name $project_name \
		--file docker-compose.yml \
		up --detach
}

function stopServices {
	echo "Shutting down services"
	cd $current_directory
	docker-compose \
		--project-name $project_name \
		--file docker-compose.yml \
		down --volumes --remove-orphans
}

function runTests {
	docker pull $docker_runner_image
	docker run \
		--network "${project_name}_default" \
		--tty \
		--volume "${current_directory}/..":"/app" \
		"${docker_runner_image}" \
		ls -l /app/scripts
}

checkDocker
startServices
runTests
stopServices
