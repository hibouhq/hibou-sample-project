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
