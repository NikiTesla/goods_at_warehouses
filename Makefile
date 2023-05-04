PROJECTNAME=$(shell basename "$(PWD)")
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

all:
	@ echo "  >  building all for $(PROJECTNAME)..."
	@ go build -o $(PROJECTNAME) cmd/main.go
	@ sudo docker start lamoda_pg
	@ sleep 1
	@ ./$(PROJECTNAME)

run:
	@ echo "  >  running cmd/main.go file..."
	@ sudo docker start lamoda_pg
	@ sleep 1
	@ go run cmd/main.go

docker:
	@ echo "  >  making docker container $(PROJECTNAME)..."
	@ sudo docker build -t $(PROJECTNAME) .
	@ sudo docker compose up