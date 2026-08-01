# Debian 10 (buster) fixture — INTENTIONALLY EOL so its dpkg packages carry
# known CVEs (deb ecosystem, release-scoped to Debian 10). Not for production.
# See ../../SECURITY-FIXTURES.md.
FROM debian:10

# buster is EOL: its packages now live on archive.debian.org and the normal
# mirror 404s, so repoint apt there and skip the valid-until check (archive
# Release files are intentionally past their expiry). buster-updates is empty
# on the archive, so drop it. Old openssl/libc packages are the intended
# findings; no upgrade on purpose so installed versions stay vulnerable.
RUN sed -i \
      -e 's|http://deb.debian.org/debian-security|http://archive.debian.org/debian-security|g' \
      -e 's|http://deb.debian.org/debian|http://archive.debian.org/debian|g' \
      -e '/buster-updates/d' \
      /etc/apt/sources.list \
    && apt-get -o Acquire::Check-Valid-Until=false update \
    && apt-get install -y --no-install-recommends openssl ca-certificates \
    && rm -rf /var/lib/apt/lists/*

CMD ["openssl", "version"]
