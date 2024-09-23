build:
	@echo
	@echo "Building snowflake in several languages"
	@echo
	@echo "C:"
	gcc -o c/recursion c/recursion.c -Wall -std=c99 -D_DEFAULT_SOURCE -Wno-missing-braces -Wunused-result -O2 -Wl,-rpath,/home/asimazbunzel/Developments/raylib/src -I. -I/home/asimazbunzel/Developments/raylib/src -I/home/asimazbunzel/Developments/raylib/src/external -I/usr/local/include -L. -L/home/asimazbunzel/Developments/raylib/src -L/home/asimazbunzel/Developments/raylib/src -L/usr/local/lib -lraylib -lGL -lm -lpthread -ldl -lrt -lX11 -lc -latomic -DPLATFORM_DESKTOP
	@echo
	@echo "Golang:"
	cd go && go mod init prj/recursion && go mod tidy && go build .

clean:
	rm -vf c/recursion go/*.mod go/*.sum go/recursion
