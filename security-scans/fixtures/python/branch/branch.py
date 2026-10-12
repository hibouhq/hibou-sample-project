import os
import subprocess


def inspect(use_input):
    command = "echo fixture"
    if use_input:
        command = os.getenv("COMMAND")
    return subprocess.run(command, shell=True)