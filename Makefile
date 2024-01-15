Push:
	go build -o push-swap pushSwap/main.go
Checek:
	go build -o checker checker/main.go
Both:
	go build -o push-swap pushSwap/main.go
	go build -o checker checker/main.go

RemPush:
	go rm -o push-swap 
RemChecek:
	go rm -o checker
Rem:
	go rm -o push-swap 
	go rm -o checker 
