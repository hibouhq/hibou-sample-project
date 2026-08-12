// reachability.go is the call-graph fixture: it pairs a REACHED vulnerable
// symbol with an UNREACHED one, so `hibou analyze reachability --lang go`
// (govulncheck) has both verdicts to report on the same dependency graph.
//
// Nothing here is exploitable — the inputs are literals and the functions are
// fixtures. See ../SECURITY-FIXTURES.md.
package calc

import (
	jwt "github.com/dgrijalva/jwt-go"
	yaml "gopkg.in/yaml.v2"
)

// ── REACHABLE ────────────────────────────────────────────────────────────
// ParseConfig calls yaml.Unmarshal, the exact symbol CVE-2019-11253 lives in,
// from an exported function — govulncheck traces a call path from the module's
// public surface to the vulnerable code and reports it as CALLED.
//
// Expected Hibou verdict: reachability = "reachable" for the yaml.v2 advisory.
func ParseConfig(raw []byte) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	if err := yaml.Unmarshal(raw, &out); err != nil {
		return nil, err
	}
	return out, nil
}

// ── UNREACHABLE ──────────────────────────────────────────────────────────
// jwtLibraryVersion touches the jwt-go package WITHOUT calling any vulnerable
// symbol: CVE-2020-26160 lives in the audience-claim parsing path
// (MapClaims.VerifyAudience / Parser.ParseWithClaims), which nothing here
// reaches. The dependency is in the SBOM, so the CVE is *matched* — but the
// call graph proves the vulnerable code is never executed.
//
// Expected Hibou verdict: reachability = "not_reachable" for CVE-2020-26160,
// which (with the org's reachability setting on) stops it counting against the
// quality gate while remaining visible in the findings list.
//
// NOTE: SignToken in insecure.go deliberately does NOT call the audience path
// either — it only signs. Keep it that way or this fixture flips to reachable.
func jwtLibraryVersion() string {
	// Referencing a type from the package keeps the import real without
	// entering any vulnerable code path.
	var claims jwt.MapClaims
	if claims == nil {
		return "dgrijalva/jwt-go v3.2.0"
	}
	return "dgrijalva/jwt-go v3.2.0 (initialized)"
}

// JWTLibraryVersion exposes the unreachable-path helper so the linter does not
// drop it as dead code — the point is that it is *called* while the vulnerable
// symbol is not.
func JWTLibraryVersion() string { return jwtLibraryVersion() }
