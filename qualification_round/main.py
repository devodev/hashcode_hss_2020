#
#

class Manager:

    def __init__(
        self, total_book_count, lib_count, day_count,
        book_scores, libraries
    ):
        self.total_book_count = total_book_count
        self.lib_count = lib_count
        self.day_count = day_count
        self.book_scores = book_scores
        self.libraries = libraries

    @classmethod
    def from_input(cls):
        libraries = list()

        total_book_count, lib_count, day_count = read_ints()
        book_scores = read_ints()
        for _ in range(lib_count):
            book_count, sig_day_count, book_count_per_day = read_ints()
            books = read_ints()
            libraries.append(Library(book_count, sig_day_count, book_count_per_day, books))
        return cls(total_book_count, lib_count, day_count, book_scores, libraries)

    def __repr__(self):
        return '\n'.join((
            '',
            'total_book_count: {}'.format(self.total_book_count),
            'lib_count: {}'.format(self.lib_count),
            'day_count: {}'.format(self.day_count),
            'book_scores: {}'.format(self.book_scores),
            '\n[libraries]\n---------------',
            *map('{}\n---------------'.format, self.libraries),
        ))


class Library:

    def __init__(self, book_count, sig_day_count, book_count_per_day, books):
        self.book_count = book_count
        self.sig_day_count = sig_day_count
        self.book_count_per_day = book_count_per_day
        self.books = books

    def __repr__(self):
        return '\n'.join((
            'book_count: {}'.format(self.book_count),
            'sig_day_count: {}'.format(self.sig_day_count),
            'book_count_per_day: {}'.format(self.book_count_per_day),
            'books: {}'.format(self.books),
        ))


def read_ints():
    return list(map(int, input().split()))


if __name__ == '__main__':
    manager = Manager.from_input()
    print(manager)