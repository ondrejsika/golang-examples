test:
	go build

	mkdir -p bin
	ln -sf $(shell pwd)/binary_aliases_example $(shell pwd)/bin/hello
	ln -sf $(shell pwd)/binary_aliases_example $(shell pwd)/bin/ahoj
	ln -sf $(shell pwd)/binary_aliases_example $(shell pwd)/bin/dela
	ln -sf $(shell pwd)/binary_aliases_example $(shell pwd)/bin/nela

	./bin/hello
	./bin/ahoj
	./bin/dela
	./bin/nela

	./binary_aliases_example

clean:
	rm -f bin/
