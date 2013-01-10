REPO_BASE := github.com/meatballhat/box-o-sand/algs4/src/go
TARGETS := \
	$(REPO_BASE)/algs4 \
	$(REPO_BASE)/algs4/algs4-binarysearch \
	$(REPO_BASE)/algs4/algs4-flips \
	$(REPO_BASE)/algs4/algs4-gcd \
	$(REPO_BASE)/algs4/algs4-rolls \
	$(REPO_BASE)/algs4/algs4-randomseq \
	$(REPO_BASE)/algs4/algs4-average \
	$(REPO_BASE)/algs4/algs4-interval2d \
	$(REPO_BASE)/algs4/algs4-date \
	$(REPO_BASE)/algs4/algs4-test-accumulator \
	$(REPO_BASE)/algs4/algs4-stats


test: build
	go test -x $(TARGETS)

build: deps
	go install -x $(TARGETS)

fmt:
	go fmt -x $(TARGETS)

deps:
	go get -x -n $(TARGETS)

clean:
	go clean -x -x $(TARGETS)

.PHONY: test build deps clean fmt
