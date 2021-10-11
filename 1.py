def main():
    j, s = input(), input()
    hm = 0
    for si in s:
        if si in j:
            hm += 1
    return hm


if __name__ == '__main__':
    print(main())
