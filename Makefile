test: 
	go test ./...
todos: 
	grep -rn --exclude=*.swp --exclude-dir=.git "TODO" .
