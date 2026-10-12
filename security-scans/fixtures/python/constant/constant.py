import os
import subprocess


def inspect():
    ignored = os.getenv("COMMAND")
    return subprocess.run("echo fixture", shell=True)