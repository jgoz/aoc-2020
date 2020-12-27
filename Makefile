gos = $(wildcard */*.go)
tests = $(wildcard */*_test.go)
days = $(patsubst %.go,%,$(filter-out $(tests),$(gos)))

$(days): %: %.go
	go build -o bin/$(notdir $@) $<

all: $(days)

debug:
	$(info $$days is [${days}])

clean:
	go clean
	rm bin/*
