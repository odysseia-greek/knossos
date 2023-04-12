INITS = drakon periandros thrasyboulos
git_short_hash=`git rev-parse --short HEAD`

copy:
	for i in $(INITS); do \
			cp ./Dockerfile ./$$i; \
	done

delete:
	for i in $(INITS); do \
			rm -f ./$$i/Dockerfile; \
	done

dist:
	for i in $(INITS); do \
			cd ./$$i && mkdir -p dist/bin/darwin && cd .. ; \
	done

build:
	for i in $(INITS); do \
			cd ./$$i && GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build main.go && mv main ./dist/bin/darwin/$$i && cd ..; \
	done

docker-build: copy dist build delete
		for i in $(INITS); do \
    			cd ./$$i && docker build --build-arg project_name=$$i -t $$i:${git_short_hash} . && cd ..;  \
    	done

docker-push: copy
		for i in $(INITS); do \
    			docker push $$i:${git_short_hash};  \
    	done