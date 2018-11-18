
freq : freq.go
	go build -ldflags="-s -w" freq.go

clean:
	rm freq
