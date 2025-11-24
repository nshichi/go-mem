
D=$(date +%Y-%m-%d)
H="${HOSTNAME%.local}"
F="${H}_${D}.txt"
cd ../ana/
go test -bench . -benchmem -run=^$ >bench-out/$F
