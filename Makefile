
PROG = freq

$(PROG) : $(PROG).go
	go build -ldflags="-s -w"

test:
	go test

clean:
	rm -f $(PROG) *~ .??*~ go.mod
