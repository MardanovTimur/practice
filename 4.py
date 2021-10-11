

def generate_sequence(n, arr, count_op_arr, count_cl_arr):
    if count_op_arr < n:
        generate_sequence(n, arr + '(', count_op_arr + 1, count_cl_arr)

    if count_cl_arr < n and count_op_arr > count_cl_arr:
        generate_sequence(n, arr + ')', count_op_arr, count_cl_arr + 1)

    if len(arr) == n * 2:
        print(arr)


def main():
    n = int(input())
    generate_sequence(n, '(', 1, 0)


if __name__ == '__main__':
    main()
