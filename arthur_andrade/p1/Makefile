

all:
	cd srcParserLPN && go build -o ../bin/parserLPN cmd/main.go && cd .. && \
	cd srcAssembler && go build -o ../bin/assembler cmd/main.go && cd .. && \
	cd srcNeanderVM && go build -o ../bin/neanderVM cmd/main.go && cd ..

test:
	./bin/parserLPN programa.lpn out/output.asm
	./bin/assembler out/output.asm out/output.bin
	./bin/neanderVM out/output.bin

clear:
	rm bin/parserLPN bin/assembler bin/neanderVM
