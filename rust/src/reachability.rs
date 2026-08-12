//! Crate-reference fixture: pairs a REFERENCED vulnerable crate with a
//! present-but-NEVER-referenced one, so `hibou analyze reachability
//! --lang rust` (crate-level: cargo metadata + source scan) has both verdicts
//! to report.
//!
//! - REACHABLE  — `time` 0.1.45 (RUSTSEC-2020-0071): named right here (and in
//!   insecure.rs), so the analyzer reports it as referenced.
//! - NOT reachable — `arrayvec` 0.4.10 (RUSTSEC-2019-0011): a dependency in
//!   Cargo.toml that NO source file ever names. It is in the lockfile and the
//!   SBOM — the CVE is matched — but the crate-reference scan proves the
//!   workspace never uses it.
//!
//! Keep the contract intact: do not `use arrayvec` anywhere in this crate, or
//! the unreachable fixture flips.

/// Names the vulnerable `time` crate so it is referenced at crate level.
pub fn epoch_seconds() -> f64 {
    time::precise_time_s()
}
