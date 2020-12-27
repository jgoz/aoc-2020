tests = $(wildcard *_test.go)
days = $(patsubst %.go,%,$(filter-out $(tests),$(wildcard *.go)))

$(days): %: %.go
	go build -o bin/$@ $<

debug:
	$(info $$days is [${days}])

clean:
	go clean
	rm bin/*
