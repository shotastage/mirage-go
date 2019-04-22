build:
	rm -rf Testing/
	mkdir Testing
	go build
	mv mirage ./Testing/
