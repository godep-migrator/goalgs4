REPO_BASE := github.com/meatballhat/goalgs4
LIBS := \
	$(REPO_BASE)
TARGETS := \
	$(LIBS) \
	$(REPO_BASE)/algs4-binarysearch \
	$(REPO_BASE)/algs4-flips \
	$(REPO_BASE)/algs4-gcd \
	$(REPO_BASE)/algs4-rolls \
	$(REPO_BASE)/algs4-randomseq \
	$(REPO_BASE)/algs4-average \
	$(REPO_BASE)/algs4-interval2d \
	$(REPO_BASE)/algs4-date \
	$(REPO_BASE)/algs4-test-accumulator \
	$(REPO_BASE)/algs4-stats \
	$(REPO_BASE)/algs4-queue-client \
	$(REPO_BASE)/algs4-stack-client \
	$(REPO_BASE)/algs4-dijkstra-two-stack-eval \
	$(REPO_BASE)/algs4-stack-client2 \
	$(REPO_BASE)/algs4-queue-client2 \
	$(REPO_BASE)/algs4-threesum \
	$(REPO_BASE)/algs4-threesum-stopwatch


test: build
	go test -x $(GOFLAGS) $(LIBS)

build: deps
	go install -x $(GOFLAGS) $(TARGETS)

deps:
	go get -x -n $(GOFLAGS) $(TARGETS)

clean:
	go clean -x -x $(TARGETS)

.PHONY: test build deps clean
