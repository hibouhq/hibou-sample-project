# Apache Tomcat fixture — INTENTIONALLY old so the bundled jars carry known
# CVEs (maven ecosystem: org.apache.tomcat:*). Unlike the alpine/debian/rhel
# fixtures this one is NOT about OS packages: Tomcat ships as jars under
# /usr/local/tomcat/lib, so it only reaches Hibou through an *image* SBOM —
# `trivy fs` over a source tree never sees it. Not for production.
# See ../../SECURITY-FIXTURES.md.
#
# TOMCAT_TAG picks which side of the experiment this build represents:
#   11.0.22 → DEFAULT. A recent release whose 2026 CVEs carry no maven
#             coordinates in GHSA (type: unreviewed, empty affected list) or OSV
#             (GIT commit ranges only) — only an NVD CPE binding. An empty result
#             here is that coordinate gap, not a Hibou matching bug.
#   11.0.2  → positive control, built by CI via --build-arg. Last release before
#             CVE-2025-24813 (RCE) was fixed in 11.0.3; OSV/GHSA carry ~20
#             advisories for org.apache.tomcat:tomcat-catalina at this version,
#             so expect a non-empty CRITICAL/HIGH set.
ARG TOMCAT_TAG=11.0.22-jre21-temurin
FROM tomcat:${TOMCAT_TAG}

# Serve the bundled ROOT app, so this is a plausible app image and the manager
# webapp jars stay in the scanned layer.
RUN mv /usr/local/tomcat/webapps.dist/ROOT /usr/local/tomcat/webapps/ROOT

EXPOSE 8080
CMD ["catalina.sh", "run"]
