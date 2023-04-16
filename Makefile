get-balance:
	go run main.go getbalance -address wallace
print:
	go run main.go printchain
create-blockchain-wallace:
	go run main.go createblockchain -address wallace
send:
	go run main.go send -from wallace -to lucia -amount 50
