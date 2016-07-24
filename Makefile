.PHONY: docker

docker:
	cp -r ../../../golang.org/x/net/context ./context
	docker build -t hkdnet/context-example .
	rm -rf ./context
