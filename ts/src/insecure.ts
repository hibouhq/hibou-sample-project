// INTENTIONALLY insecure patterns for scanner testing. Credentials are AWS's
// public documentation EXAMPLE values (non-functional). Do not copy any of this
// into production code. See ../SECURITY-FIXTURES.md.

import _ from 'lodash'
import { createHash } from 'node:crypto'
import { execSync } from 'node:child_process'

// Hardcoded credentials. NOTE: AWS's docs example key (AKIAIOSFODNN7EXAMPLE)
// sits in gitleaks' DEFAULT ALLOWLIST — scanners deliberately ignore it, so it
// proved nothing. These are randomly generated, pattern-valid, NON-FUNCTIONAL
// tokens that default rules do flag.
export const AWS_ACCESS_KEY_ID = 'AKIAIOSFODNN7EXAMPLE'
export const AWS_SECRET_ACCESS_KEY = 'wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY'
export const GITHUB_TOKEN = 'ghp_28B00KOUDWV6yVcqrloHTgsjfsiQnFhKzipc'
export const SLACK_BOT_TOKEN = 'xoxb-669936451773-610546882125-lTPZJwocCZARYk5AiQXb0Gwf'

// Weak hashing (MD5).
export function weakHash(s: string): string {
  return createHash('md5').update(s).digest('hex')
}

// Insecure, non-cryptographic RNG used as a "token".
export function weakToken(): number {
  return Math.floor(Math.random() * 1_000_000_000)
}

// Command injection — caller-controlled string handed to a shell.
export function runUserCommand(input: string): string {
  return execSync(input).toString()
}

// Code injection — eval of a caller-supplied string.
export function dangerousEval(code: string): unknown {
  // eslint-disable-next-line no-eval
  return eval(code)
}

// Uses lodash 4.17.11 (CVE-2019-10744, prototype pollution) — vulnerable dep.
export function mergeConfig(
  base: Record<string, unknown>,
  override: Record<string, unknown>,
): Record<string, unknown> {
  return _.merge({}, base, override)
}
