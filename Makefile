SHELL := /bin/bash

test-get:
	curl "http://127.0.0.1:8080/film/Ali%20Forever";
	@echo ' ';
	curl "http://127.0.0.1:8080/film/Beauty%20Grease";
	@echo ' ';
	curl "http://127.0.0.1:8080/film/Duck%20Racer";
	@echo ' ';
	curl "http://127.0.0.1:8080/film/Grail%20Frankenstein";
	@echo ' ';
	curl "http://127.0.0.1:8080/film/Hamlet%20Wisdom";
	@echo ' ';
	curl "http://127.0.0.1:8080/film/Hoosiers%20Birdcage";
	@echo ' ';