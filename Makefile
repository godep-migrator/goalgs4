REPO_BASE := github.com/meatballhat/box-o-sand/algs4/src/go
TARGETS := \
	$(REPO_BASE)/algs4 \
	$(REPO_BASE)/algs4-binarysearch \
	$(REPO_BASE)/algs4-flips \
	$(REPO_BASE)/algs4-gcd \
	$(REPO_BASE)/algs4-rolls \
	$(REPO_BASE)/algs4-randomseq \
	$(REPO_BASE)/algs4-average \
	$(REPO_BASE)/algs4-interval2d

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

.PHONY: test build clean fmt
