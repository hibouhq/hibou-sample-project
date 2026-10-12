import os
import subprocess


def inspect():
    command = os.getenv("COMMAND")
    return subprocess.run(command, shell=True)