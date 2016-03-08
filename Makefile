bin/logboom: *.go
	go build -o bin/logboom .

docker::
	docker build -t ejholmes/logboom .
