def main():
    n = int(input())

    for ind, _ in enumerate(range(n)):
        inp = int(input())
        if ind == 0:
            prev_val = inp
            print(inp)
            continue

        if inp == prev_val:
            continue
        else:
            prev_val = inp
            print(inp)


if __name__ == '__main__':
    main()
