echo "Test,Iterations,ns/op,B/op,allocs/op" > bench.log.csv

cat bench.log | tail -n +4 | tr "\t" "," | sed  's@ns/op@@' \
 | sed 's@B/op@@' | sed 's@allocs/op@@' \
 | head -n -2  >> bench.log.csv