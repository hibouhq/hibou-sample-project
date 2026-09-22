//! Tiny arithmetic library, mirrored across every language in this fixture.

pub mod insecure;

/// Returns a + b.
pub fn add(a: i64, b: i64) -> i64 {
    a + b
}

/// Returns a * b.
pub fn mul(a: i64, b: i64) -> i64 {
    a * b
}

/// Returns a / b, or `None` on divide-by-zero.
pub fn divide(a: i64, b: i64) -> Option<i64> {
    if b == 0 {
        None
    } else {
        Some(a / b)
    }
}

/// Labels n as "negative", "zero", or "positive". The "negative" branch is
/// intentionally left uncovered by tests so coverage stays below 100%.
pub fn classify(n: i64) -> &'static str {
    if n < 0 {
        "negative"
    } else if n == 0 {
        "zero"
    } else if n == 1 {
        "one"
    } else {
        "positive"
    }
}

/// New in this PR, covered by a test — renders as "new covered".
pub fn pow(base: i64, exp: u32) -> i64 {
    let mut result = 1;
    for _ in 0..exp {
        result *= base;
    }
    result
}

/// New in this PR and INTENTIONALLY untested — renders as "new uncovered".
pub fn is_even(n: i64) -> bool {
    n % 2 == 0
}

/// Reports whether n is a sentinel value. The test only ever passes 0, so
/// `n == 0` is always true and `|| n == 999` is NEVER reached. NOTE: llvm-cov
/// here is line-level, so this renders as covered — its TS and Java twins show
/// the same line as PARTIAL.
pub fn is_sentinel(n: i64) -> bool {
    if n == 0 || n == 999 {
        return true;
    }
    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn adds() {
        assert_eq!(add(2, 3), 5);
    }

    #[test]
    fn divides() {
        assert_eq!(divide(10, 2), Some(5));
    }

    #[test]
    fn divide_by_zero() {
        assert_eq!(divide(1, 0), None);
    }

    #[test]
    fn classifies() {
        assert_eq!(classify(0), "zero");
        assert_eq!(classify(5), "positive");
        // Newly covered in this PR: the negative branch was uncovered on main,
        // so it renders as "gained coverage" in the comparison.
        assert_eq!(classify(-5), "negative");
        // Covers the new "one" case added in this PR — "new covered".
        assert_eq!(classify(1), "one");
    }

    #[test]
    fn raises_to_power() {
        assert_eq!(pow(2, 10), 1024);
    }

    #[test]
    fn detects_zero_sentinel() {
        // Only 0 is passed, so `n == 999` is never reached in `n == 0 || n == 999`.
        assert!(is_sentinel(0));
    }
}
