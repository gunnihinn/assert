src=assert.go\
	assert_test.go
report=coverage.out

.PHONY: cover clean

$(report): $(src)
	go test -coverprofile=$(report)

cover: $(report)
	go tool cover -html=$(report)

clean:
	rm -f $(report)
