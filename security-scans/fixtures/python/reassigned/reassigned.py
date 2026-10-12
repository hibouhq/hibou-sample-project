import os
import subprocess


def inspect():
    command = os.getenv("COMMAND")
    command = "echo fixture"
    return subprocess.run(command, shell=True)