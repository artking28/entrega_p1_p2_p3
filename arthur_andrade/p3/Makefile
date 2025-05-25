
all:
	clear; cd cmd-bfe && go build -o ../bfe && cd .. && go build -o bfc cmd-bfc/main.go;

test:
	echo "CREDITO=2*5+10" | ./bfc | ./bfe;
	
clean:
	clear; rm ./bfe ./bfc;

