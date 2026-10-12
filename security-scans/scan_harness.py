import argparse
from collections import Counter
import json
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import time


ROOT = Path(__file__).resolve().parent
MODES = {"basic": [], "intrafile": ["--taint-intrafile"], "interfile": ["--taint-interfile"]}


def expected_results(language, case, filename, mode):
    active = filename is not None and (
        case in {"direct", "branch", "loop"}
        or (case == "intrafile" and mode != "basic")
        or (case.startswith("interfile") and mode == "interfile")
    )
    return [(f"poc.{language}.command-flow", filename)] if active else []


def observed_results(report):
    if report.get("errors"):
        raise ValueError(f"Scanner diagnostics: {report['errors']}")
    results = []
    for result in report["results"]:
        rule = result["check_id"]
        marker = rule.find("poc.")
        if marker < 0:
            raise ValueError(f"Unexpected rule ID: {rule}")
        results.append((rule[marker:], Path(result["path"]).name))
    return results


def check_results(report, expected):
    actual = observed_results(report)
    if Counter(actual) != Counter(expected):
        raise ValueError(f"Expected {expected}, received {actual}")
    return actual


def run_process(arguments, timeout, cwd=None):
    result = subprocess.run(arguments, capture_output=True, text=True, timeout=timeout, cwd=cwd)
    if result.returncode != 0:
        raise ValueError(f"Scanner exit {result.returncode}: {result.stderr[-3000:]}")
    return result.stdout


def scan_case(scanner, language, case, filename, mode, help_text, timeout):
    with tempfile.TemporaryDirectory(prefix="hibou-sast-") as workspace:
        target = Path(workspace) / "source"
        shutil.copytree(ROOT / "fixtures" / language / case, target)
        command = [
            scanner, "scan", "--config", str(ROOT / "rules.yaml"),
            "--json", "--disable-nosem", "--jobs=1", *MODES[mode], str(target),
        ]
        if "--metrics" in help_text:
            command.insert(2, "--metrics=off")
        report = json.loads(run_process(command, timeout, cwd=workspace))
        check_results(report, expected_results(language, case, filename, mode))
        paths = report.get("paths", {}).get("scanned")
        if not paths:
            raise ValueError("Scanner did not report any scanned files")
        expected_paths = {path.resolve() for path in target.rglob("*") if path.is_file()}
        observed_paths = {
            (Path(workspace) / path).resolve() if not Path(path).is_absolute()
            else Path(path).resolve() for path in paths
        }
        if observed_paths != expected_paths:
            raise ValueError("Scanner did not cover every fixture source file")
        return len(paths)


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--scanner", default="opengrep")
    parser.add_argument("--require-interfile", action="store_true")
    parser.add_argument("--timeout", type=int, default=120)
    parser.add_argument("--report", type=Path)
    arguments = parser.parse_args()
    if arguments.report:
        arguments.report.unlink(missing_ok=True)
    scanner = shutil.which(arguments.scanner)
    if not scanner:
        parser.error(f"Scanner unavailable: {arguments.scanner}")
    try:
        version = run_process([scanner, "--version"], arguments.timeout).strip()
        help_text = run_process([scanner, "scan", "--help"], arguments.timeout)
    except (ValueError, subprocess.TimeoutExpired) as error:
        parser.error(str(error))
    supported = {
        mode: not flags or flags[0] in help_text for mode, flags in MODES.items()
    }
    if arguments.require_interfile and not supported["interfile"]:
        parser.error("Required --taint-interfile is absent from scanner help")
    cases = json.loads((ROOT / "cases.json").read_text())
    outcomes = []
    for mode in MODES:
        for language, language_cases in cases.items():
            for case, filename in language_cases.items():
                outcome = {"mode": mode, "language": language, "case": case}
                if not supported[mode]:
                    outcome.update(status="skipped", reason="Scanner does not advertise mode")
                else:
                    started = time.monotonic()
                    try:
                        scanned_files = scan_case(
                            scanner, language, case, filename, mode, help_text, arguments.timeout,
                        )
                        outcome.update(status="passed", scanned_files=scanned_files)
                    except (ValueError, KeyError, subprocess.TimeoutExpired) as error:
                        outcome.update(status="failed", reason=str(error))
                    outcome["seconds"] = round(time.monotonic() - started, 3)
                outcomes.append(outcome)
                print(f"{outcome['status'].upper()} {mode}/{language}/{case}")
                if outcome.get("reason"):
                    print(f"  {outcome['reason']}")
    summary = {
        "scanner": scanner, "version": version, "supported_modes": supported,
        "counts": dict(Counter(outcome["status"] for outcome in outcomes)),
        "outcomes": outcomes,
    }
    if arguments.report:
        arguments.report.write_text(json.dumps(summary, indent=2) + "\n")
    print(json.dumps({"version": version, "counts": summary["counts"]}))
    return 1 if summary["counts"].get("failed") else 0


if __name__ == "__main__":
    sys.exit(main())