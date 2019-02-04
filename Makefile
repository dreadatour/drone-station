.PHONY: api web open-api-docs install-tests install-linter test lint lint-web

define API_DOCS_SERVE
import os, subprocess, sys
from SimpleHTTPServer import test
try:
	os.chdir("docs")
	subprocess.Popen(["python","-m","webbrowser","-t","http://localhost:8000"])
	test()
except KeyboardInterrupt:
	sys.exit(0)
endef
export API_DOCS_SERVE

api:
	@go run cmd/api/main.go -dotenv

web:
	@pushd web; yarn dev; popd

open-api-docs:
	@python -c "$$API_DOCS_SERVE"

install-tests:
	@go get -u github.com/smartystreets/goconvey

install-linter:
	@go get -u github.com/alecthomas/gometalinter
	@gometalinter --install

test:
	@go test -race -cover `go list ./...`

lint:
	@gometalinter --vendor ./...

lint-web:
	@pushd web; yarn lint; popd
