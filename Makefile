# hibou-sample-project — polyglot CI fixture.
# Every target degrades gracefully: a language whose toolchain is missing is
# skipped with a notice instead of failing the whole run.

.DEFAULT_GOAL := help
.PHONY: help test coverage ci sbom sarif clean \
        go-test go-cov ts-test ts-cov java-test java-cov rust-test rust-cov

have = $(shell command -v $(1) >/dev/null 2>&1 && echo yes)

help: ## Show targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN{FS=":.*?## "}{printf "\033[36m%-12s\033[0m %s\n",$$1,$$2}'

test: go-test ts-test java-test rust-test ## Run all tests

coverage: go-cov ts-cov java-cov rust-cov ## Run all tests + emit coverage artifacts

ci: coverage sbom sarif ## Full local CI: coverage + SBOM + SARIF

# ---- Go ----
go-test: ## Go tests
	@[ "$(call have,go)" = yes ] && (cd go && go test ./...) || echo "SKIP go: toolchain not found"

go-cov: ## Go coverage (coverage.out + junit)
	@if [ "$(call have,go)" != yes ]; then echo "SKIP go: toolchain not found"; else \
		cd go && if command -v gotestsum >/dev/null 2>&1; then \
			gotestsum --junitfile junit-go.xml --format pkgname -- -coverprofile=coverage.out ./...; \
		else \
			go test -coverprofile=coverage.out ./...; \
		fi; \
		go tool cover -func=coverage.out | tail -1; \
	fi

# ---- TypeScript ----
ts-test: ## TS tests
	@[ "$(call have,npm)" = yes ] && (cd ts && npm ci --silent && npm test) || echo "SKIP ts: npm not found"

ts-cov: ## TS coverage (lcov + junit)
	@[ "$(call have,npm)" = yes ] && (cd ts && npm ci --silent && npm run coverage) || echo "SKIP ts: npm not found"

# ---- Java ----
java-test: ## Java tests
	@[ "$(call have,mvn)" = yes ] && (cd java && mvn -q -B test) || echo "SKIP java: mvn not found"

java-cov: ## Java coverage (jacoco.xml)
	@[ "$(call have,mvn)" = yes ] && (cd java && mvn -q -B verify) || echo "SKIP java: mvn not found"

# ---- Rust ----
rust-test: ## Rust tests
	@[ "$(call have,cargo)" = yes ] && (cd rust && cargo test) || echo "SKIP rust: cargo not found"

rust-cov: ## Rust coverage (lcov.info)
	@if [ "$(call have,cargo)" != yes ]; then echo "SKIP rust: cargo not found"; \
	elif ! cargo llvm-cov --version >/dev/null 2>&1; then echo "SKIP rust cov: cargo-llvm-cov not installed"; cd rust && cargo test; \
	else cd rust && cargo llvm-cov --lcov --output-path lcov.info; fi

# ---- Security (optional tools) ----
sbom: ## CycloneDX SBOMs per language (needs trivy)
	@if [ "$(call have,trivy)" != yes ]; then echo "SKIP sbom: trivy not found"; else \
		trivy fs --format cyclonedx --output go-sbom.cdx.json go/; \
		trivy fs --format cyclonedx --output ts-sbom.cdx.json ts/; \
		trivy fs --format cyclonedx --output java-sbom.cdx.json java/; \
		trivy fs --format cyclonedx --output rust-sbom.cdx.json rust/; \
	fi

sarif: ## SAST + vuln SARIF (needs gosec + trivy)
	@command -v gosec >/dev/null 2>&1 && (cd go && gosec -fmt sarif -out ../gosec.sarif ./... || true) || echo "SKIP gosec: not found"
	@[ "$(call have,trivy)" = yes ] && (trivy fs --format sarif --output trivy-fs.sarif .) || echo "SKIP trivy: not found"

clean: ## Remove build + coverage artifacts
	@rm -rf go/coverage.out go/junit-go.xml go/app \
		ts/coverage ts/junit-front.xml ts/node_modules \
		java/target rust/target rust/lcov.info \
		*.sarif *.cdx.json
	@echo "cleaned"
