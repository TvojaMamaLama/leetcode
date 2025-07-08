def merge(intervals: list[list[int]]) -> list[list[int]]:
    intervals.sort(key=lambda x: x[0])

    answer = []

    for interval in intervals:
        if not answer or answer[-1][1] < interval[0]:
            answer.append(interval)
        else:
            answer[-1][1] = max(interval[1], answer[-1][1])

    return answer


print(merge([[1, 3], [2, 6], [8, 10], [15, 18]]))
