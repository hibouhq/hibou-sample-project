//! Tiny arithmetic library, mirrored across every language in this fixture.

pub mod insecure;

/// Returns a + b.
pub fn add(a: i64, b: i64) -> i64 {
    a + b
}

/// Returns a - b.
pub fn sub(a: i64, b: i64) -> i64 {
    a - b
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
    } else {
        "positive"
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn adds() {
        assert_eq!(add(2, 3), 5);
    }

    #[test]
    fn subtracts() {
        assert_eq!(sub(5, 3), 2);
    }

    #[test]
    fn multiplies() {
        assert_eq!(mul(4, 3), 12);
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
        // classify(negative) intentionally untested → keeps coverage < 100%.
    }
}
