include .env
export

d=$$(date -Iseconds) # use ISO 8601

server-gen:
	go generate ./api
	sed -i '' -e 's/\[\]/\[\]\*/g' ./api/generation.go

server-build:
	GOAMD64=v3 go build -race -o server.bin -ldflags="-X 'conf.Version=0.1.0' -X 'conf.Author=$$(users)' -X 'conf.Time=$d' -X 'conf.Contact=rafif.mulia.r@gmail.com'" ./cmd/server

server-run: server-build
	./server.bin -debug -prefork

# Current profiling still doesn't support preforking, because it will write to the same filename.
server-runtime-pprof: server-build
	./server.bin -debug -cpuprofile=./performance/server.prof -memprofile=./performance/server.mprof

# Current profiling still doesn't support preforking, because they will use same tcp address.
# We don't know what we are accessing is child 1 or child 2.
server-live-pprof: server-build
	./server.bin -debug -httplivepprof=true
