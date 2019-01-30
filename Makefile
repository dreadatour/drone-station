.PHONY: api

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

open_api_docs:
	@python -c "$$API_DOCS_SERVE"
