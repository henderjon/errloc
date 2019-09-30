export CC_TEST_REPORTER_ID = e4fb27f2de54facb74a8ea8467871b361047875335db8cc5b77089b3a54ec627
COVERAGEOUTFILE=c.out

all: test race

.PHONY: dep
dep:
	go mod vendor

.PHONY: test-vendor
test-vendor:
	go test -mod=vendor -coverprofile=coverage.out -covermode=count

.PHONY: test
test: dep
	go test -coverprofile=coverage.out -covermode=count

.PHONY: race
race: dep
	go test -race

.PHONY: test-report
test-report: test
	go tool cover -html=coverage.out

.PHONY: travis
travis:
	# install deps pre-1.13
	go get -u github.com/google/go-cmp/cmp
	go test -coverprofile $(COVERAGEOUTFILE) ./...

.PHONY: cclimate-linux
cclimate-linux:
	curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
	# curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-darwin-amd64 > ./cc-test-reporter
	chmod +x ./cc-test-reporter
	./cc-test-reporter before-build
	# install deps pre-1.13
	go get -u github.com/google/go-cmp/cmp
	go test -coverprofile $(COVERAGEOUTFILE) ./...
	./cc-test-reporter after-build --exit-code $(TRAVIS_TEST_RESULT)
