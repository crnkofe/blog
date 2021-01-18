from __future__ import annotations

# with annotations a class can type constrain self
# ie. def __iter__(self) -> LineIterator:
# see https://www.python.org/dev/peps/pep-0563/

import argparse
import logging
from io import IOBase
from typing import List, Tuple


logging.basicConfig(level=logging.DEBUG)


def create_argument_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser()
    parser.add_argument("-f", "--filename", type=str, help="CSV filename")
    parser.add_argument(
        "-s", "--separator",
        type=str, help="CSV field separator character or string")
    parser.add_argument(
        "-r", "--header",
        action="store_true", help="Does CSV have a header")
    return parser


class LineIterator(object):

    def __init__(self,
                 file_object: IOBase,
                 has_header: bool = True,
                 separator: str = ";"):
        self.file_object = file_object
        self.has_header = has_header
        self.separator = separator
        self.first_line = True

    def __next__(self) -> List[str]:
        line = next(self.file_object)
        if self.first_line:
            next(self.file_object)
            self.first_line = False
        if line:
            return line.strip().split(self.separator)

    def __iter__(self) -> LineIterator:
        return self


class CsvFile(object):

    def __init__(self,
                 file_name: str,
                 method: str,
                 has_header: bool = True,
                 separator: str = ";"):
        self.has_header = has_header
        self.separator = separator
        self.file_name = file_name
        self.method = method
        self.file_obj = None

    def __enter__(self) -> Tuple[List[str], LineIterator]:
        if not self.file_obj:
            try:
                self.file_obj = open(self.file_name, self.method)
            except Exception as e:
                # if this is rethrown caller needs to handle the exception
                logging.error("Failed opening file")

        header = []
        if self.has_header:
            first_line = next(self.file_obj, None)
            if first_line:
                header = first_line.split(self.separator)
        lines = LineIterator(self.file_obj, self.has_header, self.separator)
        return header,  lines

    def __exit__(self, type, value, traceback):
        if value:
            logging.error("Failed processing CSV",
                          exc_info=(type, value, traceback))
        if self.file_obj:
            self.file_obj.close()
        # return False for leaving exception handling to user
        return True


if __name__ == "__main__":
    argument_parser = create_argument_parser()
    args = argument_parser.parse_args()
    with CsvFile(args.filename,
                 "r",
                 has_header=args.header,
                 separator=args.separator) as (csv_header, csv_file):
        for line in (x for i, x in enumerate(csv_file) if i < 5):
            print(line)
