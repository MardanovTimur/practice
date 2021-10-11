def check_and_move(a, b):
    return 1 if "".join(sorted(list(a))) == "".join(sorted(list(b))) else 0


def main():
    a = input()
    b = input()

    if len(a) != len(b):
        return 0

    return check_and_move(a.lower(), b.lower())


if __name__ == '__main__':
    print(main())
