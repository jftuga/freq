
PROG = freq

$(PROG) : $(PROG).go
	go build -ldflags="-s -w"

clean:
	rm -f $(PROG) *~ .??*~ go.mod
