SHELL := /bin/bash


NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m



run:
	@echo ''
	@printf '$(OK_COLOR)快糙猛，跑一哈 .. 🚀$(NO_COLOR)\n'
	@go run ./main.go
	@echo '🎯'
	@echo ''
	



container_open:
	@echo ''
	@printf '$(OK_COLOR)打开容器服务 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml up -d 
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''


container_close:
	@echo ''
	@printf '$(OK_COLOR)关闭容器服务 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml down 
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''




login_mysql:
	@echo ''
	@printf '$(OK_COLOR)登录 MySQL 容器 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml exec db sh -c 'mysql -uroot -p${MYSQL_ROOT_PASSWORD}'
	@printf '$(OK_COLOR)退出 .. 🎯$(NO_COLOR)\n'
	@echo ''



login_redis:
	@echo ''
	@printf '$(OK_COLOR)登录 Redis 容器 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml exec cache sh -c 'redis-cli'
	@printf '$(OK_COLOR)退出 .. 🎯$(NO_COLOR)\n'
	@echo ''

 
container_detail:
	@echo ''
	@printf '$(OK_COLOR)查看容器配置 .. 🚀$(NO_COLOR)\n'
	@docker-compose -f ./docker-compose.yml config
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''


container_net:
	@echo ''
	@printf '$(OK_COLOR)查看 MySQL 容器 IP 地址 .. 🚀$(NO_COLOR)\n'
	@docker inspect db | grep IPAddress
	@echo ''
	@printf '$(OK_COLOR)查看 Redis 容器 IP 地址 .. 🚀$(NO_COLOR)\n'
	@docker inspect cache | grep IPAddress
	@echo ''
	@printf '$(OK_COLOR) .. 🎯$(NO_COLOR)\n'
	@echo ''





