def sort_array(l: list) -> list:
    i = 0
    j = len(l)

    answer = []
    while i < j:
        if j * j > i * i:
            answer.append(j * j)
            j -= 1
        else:
            answer.append(i * i)
            i -= 1

    return answer[::-1]


print(sort_array([-2, 0, 1, 3, 5]))
