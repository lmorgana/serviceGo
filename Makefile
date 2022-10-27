SRCS = ./srcs/

all: up

up:
	@ cd $(SRCS) &&  docker-compose up --build;
down:
	@ cd $(SRCS) && docker-compose down;

clean:
	@ cd $(SRCS) && docker-compose down && docker system prune ;

fclean:
	@ cd $(SRCS) && docker-compose down && docker system prune -a;

.PHONY: fclean run all