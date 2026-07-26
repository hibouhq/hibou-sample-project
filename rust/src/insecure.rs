//! INTENTIONALLY insecure patterns for scanner testing. Credentials are AWS's
//! public documentation EXAMPLE values (non-functional). Do not copy any of this
//! into production code. See ../SECURITY-FIXTURES.md.

use std::process::{Command, Output};

// Hardcoded credentials — secret scanners flag these.
pub const AWS_ACCESS_KEY_ID: &str = "AKIAIOSFODNN7EXAMPLE";
pub const AWS_SECRET_ACCESS_KEY: &str = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY";

/// Command injection — caller-controlled string handed to a shell.
pub fn run_user_command(input: &str) -> std::io::Result<Output> {
    Command::new("sh").arg("-c").arg(input).output()
}

/// Uses the vulnerable `time` 0.1.45 crate (RUSTSEC-2020-0071).
pub fn timestamp() -> f64 {
    time::precise_time_s()
}

/// Insecure, predictable "token" derived from the clock rather than a CSPRNG.
pub fn weak_token() -> i64 {
    let t = time::now();
    i64::from(t.tm_sec) * 1000 + i64::from(t.tm_nsec) / 1_000_000
}

/// An unnecessary `unsafe` raw-pointer dereference (clippy / audit signal).
pub fn unsafe_identity(x: i64) -> i64 {
    let p = &x as *const i64;
    unsafe { *p }
}
