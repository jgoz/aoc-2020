days = $(wildcard day*)
bins = $(addprefix bin/,$(notdir $(days)))

# Search in day* folders for .go files
vpath %.go $(days)

.PHONY: all
all: $(bins)

.PHONY: clean
clean:
	go clean
	rm bin/*

# Actual binary targets (bin/day1, bin/day2, etc)
$(bins): bin/%: %.go
	go build -o $@ $<

# Aliases (`make day1` -> `make bin/day1`)
.PHONY: $(days)
$(days): %: bin/%
