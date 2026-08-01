# RHEL-family fixture — UBI 8.5, INTENTIONALLY pinned to an old build so its rpm
# packages carry known CVEs (rpm ecosystem, release-scoped to RHEL 8). Uses the
# public Red Hat UBI image, no subscription needed. Not for production.
# See ../../SECURITY-FIXTURES.md.
FROM redhat/ubi8:8.5

# The base image's rpms (openssl-libs, glibc, platform-python, …) at the 8.5
# snapshot are the intended findings; no `dnf update` on purpose.
CMD ["cat", "/etc/redhat-release"]
