def get_top(l: list, k: int) -> list:
    counter = dict()

    for num in l:
        if num not in counter:
            counter[num] = 1
            continue

        counter[num] += 1

    sorted_freq = [[] for _ in range(len(l))]
    for key, value in counter.items():
        sorted_freq[value].append(key)

    top = []
    for bucket in sorted_freq[::-1]:
        if bucket:
            for number in bucket:
                top.append(number)

    return top[:k]


print(get_top([1, 1, 1, 1, 2, 3, 3, 3, 4, 4], 3))
