CC = go build
CFLAGS =

all: ran

ran:
	mkdir -p target
	$(CC) $(CFLAGS) -o target/ran cmd/ran/ran.go

test:
	python3 test/test_all.py

clean:
	rm -rf target/