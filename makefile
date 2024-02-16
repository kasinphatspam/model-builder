dc-nuclear:
	- docker stop $$(docker ps -a -q)
	- docker kill $$(docker ps -q)
	- docker rm $$(docker ps -a -q)
	- docker rmi $$(docker images -q)
	- docker system prune --all --force --volumes

dcup-dev:
	- docker-compose up -d

dc-down:
	- docker-compose down