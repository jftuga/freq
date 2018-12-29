
PROG = freq

$(PROG) : $(PROG).go
	VERS='$(shell awk -F/ '/refs\/tags\// {print $$NF}' .git/packed-refs | tail -1)' ; \
	go build -ldflags "-s -w -X main.version=$$VERS" $(PROG).go

clean:
	rm -f $(PROG) *~ .??*~
