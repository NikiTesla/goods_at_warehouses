PROJECTNAME=$(shell basename "$(PWD)")
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

all:
	@ echo "  >  building all for $(PROJECTNAME)..."
	@ go build -o $(PROJECTNAME) cmd/main.go
	@ sudo docker start goods_at_warehouses_pg
	@ sleep 0.1
	@ ./$(PROJECTNAME)

run:
	@ echo "  >  running cmd/main.go file..."
	@ sudo docker start goods_at_warehouses_pg
	@ sleep 0.1
	@ go run cmd/main.go

docker: migration-up
	@ sudo docker rmi -f $(PROJECTNAME)
	@ echo "  >  making docker container $(PROJECTNAME)..."
	@ sudo docker build -t $(PROJECTNAME) .
	@ sudo docker compose up

# usage make migration-up ARGS="[version]" 
migration-up:
	@ echo "  >  making migrations"
	@ sudo docker start goods_at_warehouses_pg
	@ sleep 0.1
	@ cat schemas/0001_init.up.sql | sudo docker exec -i goods_at_warehouses_pg  psql -U postgres -d postgres

# usage make migration-down ARGS="[version]" 
migration-down:
	@ echo "  >  making migrations"
	@ sudo docker start goods_at_warehouses_pg
	@ sleep 0.1
	@ cat schemas/0001_init.down.sql | sudo docker exec -i goods_at_warehouses_pg  psql -U postgres -d postgres