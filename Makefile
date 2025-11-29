.PHONY: generate gen-buf gen-rust

generate: gen-buf gen-rust

gen-buf:
	@echo "Running buf generate..."
	buf generate

gen-rust:
	@echo "Running rust generator..."
	cd tools/gen-rust && cargo run
