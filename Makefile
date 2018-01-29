OBJ:=github

build:
	go build -o ${OBJ} main.go

clean:
	-rm ${OBJ}

install: build
ifdef GOBIN
	mv ${OBJ} ${GOBIN}
	@exit 0
endif
ifdef GOPATH
	mkdir -p ${GOPATH}/bin
	mv ${OBJ} ${GOPATH}/bin
	@exit 0
endif


.PHONY:
	build clean
