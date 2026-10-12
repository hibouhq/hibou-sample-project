import subprocess


def execute(command):
    return subprocess.run(command, shell=True)