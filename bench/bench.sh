TODAY=$(date +%Y-%m-%d)
f=${HOSTNAME}_${TODAY}.txt
cd ../ana/
go test -bench . -benchmem -run=^$ >$f
