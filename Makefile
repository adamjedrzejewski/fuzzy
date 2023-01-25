fuzzy: main.o fuzzy.o
	cc main.o fuzzy.o -o fuzzy

main.o:
	cc -c main.c -o main.o

fuzzy.o:
	cc -c fuzzy.c -o fuzzy.o

clean:
	rm -f fuzzy *.o test
