debug:
	go run cmd/gocker/main.go | buildctl debug dump-llb --dot | dot -T png -o out.png
