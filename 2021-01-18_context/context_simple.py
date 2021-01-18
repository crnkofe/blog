import sys
from io import IOBase
from typing import Tuple, List


class SimpleContextManager(object):

    def __init__(self, file_name: str, method: str):
        self.file_obj = open(file_name, method)

    def __enter__(self) -> IOBase:
        return self.file_obj

    def __exit__(self, type, value, traceback):
        self.file_obj.close()
        # return False for leaving exception handling to user
        return True


if __name__ == "__main__":
    with SimpleContextManager(sys.argv[1], "r") as file:
        for line in file:
            print(line.strip())
