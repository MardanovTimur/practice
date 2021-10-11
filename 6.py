
def distance(x1, x2):
    return abs(x1[0] - x2[0]) + abs(x1[1] - x2[1])


def minimal_distance(current, end, fuel, cities, visited, all_dists):
    if current == end:
        all_dists.add(len(visited))
        return

    if len(all_dists) and len(visited) >= min(all_dists):
        return

    if fuel >= distance(cities[current - 1], cities[end - 1]):
        all_dists.add(len(visited) + 1)
        return

    for id, i in enumerate(cities):
        if id + 1 == current or id + 1 in visited or id + 1 == end:
            continue

        if fuel >= distance(cities[current - 1], cities[id]):
            minimal_distance(id + 1, end, fuel, cities, visited + [id + 1], all_dists)


def main():
    n = int(input())

    cities = []

    for i in range(n):
        cities.append([int(a) for a in input().split(" ")])

    fuel = int(input())
    start, end = [int(a) for a in input().split(" ")]
    all_dists = set()
    minimal_distance(start, end, fuel, cities, [start], all_dists)

    if len(all_dists) == 0:
        print(-1)
    else:
        print(min(all_dists) - 1)


if __name__ == '__main__':
    main()
