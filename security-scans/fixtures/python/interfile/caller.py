import os
from sink import execute


def inspect():
    return execute(os.getenv("COMMAND"))