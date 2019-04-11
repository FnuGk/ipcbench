PROTOS = simple.pb.go nested.pb.go
PLOTS = Unmarshal-Nested-plot.png Marshal-Nested-plot.png

%.pb.go: %.proto
	protoc --go_out=. $^


.PHONY: bench.log bench.log.csv $(PLOTS) all

bench.log: $(PROTOS)
	go test -bench=. -benchmem -benchtime 5s -count 30 -test.timeout 0 > bench.log


bench.log.csv: bench.log
	bash ./cleanup.sh


# $(PLOTS): bench.log.csv
# 	python3 plot.py

plots: bench.log.csv
	python3 plot.py

all: plots

# all: $(PLOTS)
