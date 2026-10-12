import os
import subprocess


def execute(command):
    return subprocess.run(command, shell=True)


def inspect():
    return execute(os.getenv("COMMAND"))