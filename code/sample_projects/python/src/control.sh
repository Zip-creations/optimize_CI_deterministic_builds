#!/usr/bin/bash

PYTHONPATH=src python3 -m pytest tests/tests.py --junit-xml="tests/test-reports/report.xml"
