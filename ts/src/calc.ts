// A tiny arithmetic library, mirrored across every language in this fixture.

export function add(a: number, b: number): number {
  return a + b
}

export function sub(a: number, b: number): number {
  return a - b
}

export function mul(a: number, b: number): number {
  return a * b
}

export function divide(a: number, b: number): number {
  if (b === 0) {
    throw new Error('divide by zero')
  }
  return a / b
}

// The compound condition sits on one line and the tests exercise only some of
// its outcomes, so lcov reports the line as PARTIALLY covered (BRDA taken on a
// subset) — the fixture for partial-branch rendering.
export function clamp(n: number, lo: number, hi: number): number {
  if (n < lo || n > hi) {
    return n < lo ? lo : hi
  }
  return n
}

// The "negative" branch is intentionally left uncovered → keeps coverage < 100%.
export function classify(n: number): string {
  if (n < 0) {
    return 'negative'
  }
  if (n === 0) {
    return 'zero'
  }
  return 'positive'
}

// ---- Partial-branch fixtures ----
// Each function below puts several branch outcomes on one line and the tests
// exercise only a subset, so the line reports PARTIAL branch coverage.

// `&&` chain: tests never pass a valid port, so the right-hand conditions
// short-circuit away.
export function isValidPort(n: number, allowPrivileged = false): boolean {
  return Number.isInteger(n) && n > 0 && n < 65536 && (allowPrivileged || n >= 1024)
}

// Nullish coalescing: tests always pass a value, the fallback arm never runs.
export function withUnit(value: number, unit?: string): string {
  return `${value} ${unit ?? 'ms'}`
}

// Chained ternary: tests only ever reach the first arm.
export function grade(score: number): string {
  return score >= 90 ? 'A' : score >= 75 ? 'B' : score >= 50 ? 'C' : 'F'
}

// New in this PR — TEXTBOOK partial coverage: the test only ever passes n === 0,
// so the first operand is always true and `n === 999` is NEVER reached (the || short-
// circuits). The line runs (covered) but one branch outcome is missing → partial.
export function isSentinel(n: number): boolean {
  if (n === 0 || n === 999) {
    return true
  }
  return false
}
