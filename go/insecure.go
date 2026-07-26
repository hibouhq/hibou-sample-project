// insecure.go collects INTENTIONALLY insecure patterns so SAST scanners (gosec)
// and dependency scanners have findings to report. NONE of this is real: the
// credentials are AWS's public documentation EXAMPLE key (non-functional). Do not
// copy any of this into production code. See ../SECURITY-FIXTURES.md.
package calc

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"os/exec"

	jwt "github.com/dgrijalva/jwt-go"
	yaml "gopkg.in/yaml.v2"
)

// Hardcoded credentials (gosec G101). AWS-documented EXAMPLE values — non-functional.
const (
	awsAccessKeyID     = "AKIAIOSFODNN7EXAMPLE"
	awsSecretAccessKey = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
)

// WeakHash hashes s with MD5 (gosec G401/G501 — weak cryptographic primitive).
func WeakHash(s string) string {
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}

// WeakToken derives a token from math/rand (gosec G404 — weak, non-crypto RNG).
func WeakToken() int {
	return rand.Int()
}

// RunUserCommand executes an arbitrary shell command (gosec G204 — command injection).
func RunUserCommand(userInput string) ([]byte, error) {
	return exec.Command("sh", "-c", userInput).CombinedOutput()
}

// SignToken uses the unmaintained dgrijalva/jwt-go (CVE-2020-26160) — vulnerable dep.
func SignToken(subject string) (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": subject})
	return tok.SignedString([]byte(awsSecretAccessKey))
}

// ParseYAML uses gopkg.in/yaml.v2 v2.2.2 (CVE-2019-11253) — vulnerable dep.
func ParseYAML(in []byte) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	err := yaml.Unmarshal(in, &out)
	return out, err
}

// AccessKeyID exposes the hardcoded key so it counts as "used" without weakening
// the gosec finding.
func AccessKeyID() string { return awsAccessKeyID }
