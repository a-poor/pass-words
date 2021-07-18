import re
import functools as ft
from pathlib import Path
from typing import Set

SRC_DIR  = "./books"
DST_PATH = "./vocab.txt"

def tokenize_book(book_text: str) -> Set[str]:
    """Tokenize a book into a set of words.

    :param book_text: The raw text of a book.
    :returns: A normalized set of words in the book. All words
        are lowercase, and non-letter characters are removed.
    """
    text = book_text.lower().strip()
    text = re.sub(r"[^a-z]", " ", text)
    text = re.sub(r"\s+", " ", text)
    words = re.split(r"\s+", text)
    return set(words)

if __name__ == "__main__":
    # Get all books
    book_files = Path(SRC_DIR).rglob("*.txt")

    # Read the books
    books = (p.read_text() for p in book_files)

    # Tokenize the books
    book_vocabs = (tokenize_book(book) for book in books)

    # Process all the books
    vocab = sorted(ft.reduce(lambda a, b: a | b, book_vocabs))

    if "" in vocab:
        vocab.remove("")

    # Write the vocab to file
    Path(DST_PATH).write_text("\n".join(vocab))

    # Output the results
    print(f"Done. Found {len(vocab):,d} unique words.")

