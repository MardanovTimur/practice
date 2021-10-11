def main():
    x = []
    n = int(input())
    for _ in range(n):
        x.append(int(input()))

    temp_max_length = 0
    max_length = 0
    for v in x:
        if v == 1:
            temp_max_length += 1
        else:
            if temp_max_length > max_length:
                max_length = temp_max_length
            temp_max_length = 0

    if temp_max_length > max_length:
        max_length = temp_max_length

    return max_length



if __name__ == '__main__':
    print(main())
