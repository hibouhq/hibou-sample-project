import os
import subprocess


def allow_command(value):
    return "echo fixture"


def inspect():
    return subprocess.run(allow_command(os.getenv("COMMAND")), shell=True)