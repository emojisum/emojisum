import hashlib
import json
import mmap
import os
import re
import sys
from pathlib import Path
from typing import NamedTuple, Optional

from rich.console import Console

SHA1SUM_PATTERN = re.compile(r"(?P<digest>\S*)\s+(?P<path>.*)")
OPENSSL_SHA1_PATTERN = re.compile(r"SHA1\((?P<path>.*)\)=\s+(?P<digest>.*)")
OUTPUT_PATTERNS = [SHA1SUM_PATTERN, OPENSSL_SHA1_PATTERN]


class Sha1Result(NamedTuple):
    hex_digest: str
    path: str


def load_table():
    with (Path(__file__).parent / "emojimap.json").open() as f:
        data = json.load(f)

    return data["emojiwords"]


def _emojify(hex_digest):
    for i in range(0, len(hex_digest), 2):
        yield load_table()[int(hex_digest[i : i + 1], 0x10)][0]


def emojify(hex_digest: str):
    return "".join(_emojify(hex_digest))


def sha1sum(path: Path):
    h = hashlib.sha1()
    with path.open("rb") as f:
        with mmap.mmap(f.fileno(), 0, access=mmap.ACCESS_READ) as mm:
            h.update(mm)
    return h.hexdigest()


def parse_sha1_output(output: str) -> Optional[Sha1Result]:
    for pattern in OUTPUT_PATTERNS:
        match = pattern.match(output)
        if match:
            return Sha1Result(hex_digest=match["digest"], path=match["path"])

    return None


def print_result(result: Sha1Result):
    path = os.path.normpath(result.path)
    hex_digest = result.hex_digest

    markdown = emojify(hex_digest=hex_digest)

    Console().print(
        f"SHA1({path}) = {hex_digest}", emoji=False, markup=False, highlight=False
    )
    Console().print(
        f"SHA1({path}) = {markdown}", emoji=False, markup=False, highlight=False
    )
    Console().print(
        f"SHA1({path}) = {markdown}", emoji=True, markup=False, highlight=False
    )


def calculate_sum(path: Path) -> Sha1Result:
    return Sha1Result(path=str(path), hex_digest=sha1sum(path))


def main(path: Optional[Path] = None):
    if path:
        sha1_result = calculate_sum(path)

    elif not sys.stdin.isatty():
        sha1_result = parse_sha1_output(sys.stdin.read())
    else:
        return

    if not sha1_result:
        return

    print_result(sha1_result)


def entry():
    if len(sys.argv) == 2:
        path = Path(sys.argv[1])
    else:
        path = None

    main(path)


if __name__ == "__main__":
    entry()
