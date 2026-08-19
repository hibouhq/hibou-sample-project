# Alpine 3.9 fixture — INTENTIONALLY old so the OS packages carry known CVEs
# (apk ecosystem, release-scoped to Alpine 3.9). Also installs a vulnerable
# dockerize from GitHub releases to exercise binary detection. Not for
# production. See ../../SECURITY-FIXTURES.md.
FROM alpine:3.9

# dockerize v0.6.1: bundles an old golang.org/x/net / x/crypto with published
# advisories, and is itself EOL. Pinned so the finding is reproducible.
ARG DOCKERIZE_VERSION=v0.6.1
RUN apk add --no-cache openssl curl \
    && curl -fsSL "https://github.com/jwilder/dockerize/releases/download/${DOCKERIZE_VERSION}/dockerize-alpine-linux-amd64-${DOCKERIZE_VERSION}.tar.gz" \
       | tar -xz -C /usr/local/bin \
    && chmod +x /usr/local/bin/dockerize

ENTRYPOINT ["/usr/local/bin/dockerize"]
