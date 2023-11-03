generate:
	@buf mod update;
	@buf lint --config buf.yaml;
	@echo "generating stubs for Go and Python"
	@buf generate --config buf.yaml --template buf.gen.yaml