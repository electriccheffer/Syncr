test: 
	go test ./...
todos: 
	grep -rn --exclude=Makefile --exclude=*.swp --exclude-dir=.git "TODO" .
