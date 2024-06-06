
.PHONY: gen-openapi
gen-openapi: clean_gen gen-openapi-server tidy


.PHONY: clean_gen
clean_gen:
	@rm -rf ./api/gen


.PHONY: gen-openapi-server
gen-openapi-server:
	mkdir -p ./api/gen/xiaohan
	bin/swagger generate server -t ./api/gen/xiaohan -f ./api/spec/xiaohan.yaml -s server --exclude-main -A Xiaohan

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run_mongo
run_mongo:
	docker run -d -p 27017:27017 --name xiaohan_mongo mongo:4.4.6