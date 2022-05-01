CC = go build
CFLAGS =

all: mtsvc

mtsvc:
	mkdir -p target
	$(CC) $(CFLAGS) -o target/mtsvc cmd/mtsvc/mtsvc.go

clean:
	rm -rf target/