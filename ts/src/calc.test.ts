import { describe, it, expect } from 'vitest'
import { add, sub, mul, divide, classify } from './calc'

describe('calc', () => {
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
