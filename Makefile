worker: 
	go run ./worker/worker.go

client: 
	go run ./client/client.go

.PHONY: worker client
