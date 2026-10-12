import json
from pathlib import Path
import unittest
from unittest.mock import patch

from scan_harness import ROOT, check_results, expected_results, scan_case


class ScanHarnessTests(unittest.TestCase):
    def test_direct_flow_required_in_every_mode(self):
        for mode in ("basic", "intrafile", "interfile"):
            self.assertEqual(
                [("poc.java.command-flow", "Direct.java")],
                expected_results("java", "direct", "Direct.java", mode),
            )

    def test_cross_file_requires_interfile(self):
        self.assertEqual([], expected_results("java", "interfile", "Sink.java", "basic"))
        self.assertEqual([], expected_results("java", "interfile", "Sink.java", "intrafile"))
        self.assertEqual(
            [("poc.java.command-flow", "Sink.java")],
            expected_results("java", "interfile", "Sink.java", "interfile"),
        )

    def test_same_file_helper_requires_cross_function_mode(self):
        self.assertEqual([], expected_results("python", "intrafile", "helpers.py", "basic"))
        self.assertEqual(
            [("poc.python.command-flow", "helpers.py")],
            expected_results("python", "intrafile", "helpers.py", "intrafile"),
        )

    def test_negative_controls_never_expect_findings(self):
        for mode in ("basic", "intrafile", "interfile"):
            self.assertEqual([], expected_results("java", "constant", None, mode))
            self.assertEqual([], expected_results("java", "sanitized", None, mode))

    def test_empty_report_is_valid_negative_control(self):
        self.assertEqual([], check_results({"results": [], "errors": []}, []))

    def test_diagnostics_cannot_pass(self):
        with self.assertRaises(ValueError):
            check_results({"results": [], "errors": [{"type": "ParseError"}]}, [])

    def test_duplicate_findings_cannot_pass(self):
        result = {"check_id": "poc.java.command-flow", "path": "Direct.java"}
        with self.assertRaises(ValueError):
            check_results({"results": [result, result]}, [("poc.java.command-flow", "Direct.java")])

    def test_missing_expected_finding_cannot_pass(self):
        with self.assertRaises(ValueError):
            check_results({"results": []}, [("poc.java.command-flow", "Direct.java")])

    def test_prefixed_rule_id_is_normalized(self):
        result = {"check_id": "security-scans.poc.java.command-flow", "path": "src/Direct.java"}
        self.assertEqual(
            [("poc.java.command-flow", "Direct.java")],
            check_results({"results": [result]}, [("poc.java.command-flow", "Direct.java")]),
        )

    def test_extended_interfile_flow_requires_interfile(self):
        for case in ("interfile_return", "interfile_inheritance"):
            self.assertEqual([], expected_results("java", case, "Sink.java", "intrafile"))
            self.assertEqual(
                [("poc.java.command-flow", "Sink.java")],
                expected_results("java", case, "Sink.java", "interfile"),
            )

    def test_case_has_isolated_project_root(self):
        workspaces = []

        def scanner_output(arguments, timeout, cwd=None):
            self.assertNotEqual(ROOT.parent, Path(cwd))
            workspaces.append(Path(cwd))
            target = Path(arguments[-1])
            self.assertEqual(2, len(list(target.glob("*.java"))))
            return json.dumps({
                "results": [], "errors": [],
                "paths": {"scanned": [str(path) for path in target.glob("*.java")]},
            })

        with patch("scan_harness.run_process", side_effect=scanner_output):
            self.assertEqual(2, scan_case("scanner", "java", "interfile", "Sink.java", "basic", "", 10))
        self.assertFalse(workspaces[0].exists())

    def test_partial_source_coverage_cannot_pass(self):
        def scanner_output(arguments, timeout, cwd=None):
            return json.dumps({
                "results": [], "errors": [],
                "paths": {"scanned": [str(Path(arguments[-1]) / "Caller.java")]},
            })

        with patch("scan_harness.run_process", side_effect=scanner_output):
            with self.assertRaisesRegex(ValueError, "every fixture source"):
                scan_case("scanner", "java", "interfile", "Sink.java", "basic", "", 10)