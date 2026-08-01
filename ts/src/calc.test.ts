import { describe, it, expect } from 'vitest'
import { add, sub, mul, divide, classify, clamp, isValidPort, withUnit, grade } from './calc'

describe('calc', () => {
  // Deliberately use only the below-range case: `n < lo` short-circuits, so
  // the `n > hi` branch remains unvisited on this executed source line.
  it('clamps below range', () => {
    expect(clamp(-3, 0, 10)).toBe(0)
  })

  // Each test below exercises only a subset of the branch outcomes on the
  // corresponding one-liner, keeping those lines PARTIALLY covered.
  it('rejects a non-integer port', () => expect(isValidPort(1.5)).toBe(false))
  it('formats with an explicit unit', () => expect(withUnit(3, 's')).toBe('3 s'))
  it('grades an A', () => expect(grade(95)).toBe('A'))

  it('adds', () => expect(add(2, 3)).toBe(5))
  it('subtracts', () => expect(sub(5, 3)).toBe(2))
  it('multiplies', () => expect(mul(4, 3)).toBe(12))
  it('divides', () => expect(divide(10, 2)).toBe(5))
  it('throws on divide by zero', () => expect(() => divide(1, 0)).toThrow('divide by zero'))

  it('classifies', () => {
    expect(classify(0)).toBe('zero')
    expect(classify(5)).toBe('positive')
    // classify(negative) intentionally untested → keeps coverage < 100%.
  })
})
