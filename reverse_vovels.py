def reverse_vowels(s: str):
    vowels = {'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
    pointer1 = 0
    pointer2 = len(s) - 1
    l = list(s)

    while pointer1 < pointer2:
        while pointer1 < pointer2 and l[pointer1] not in vowels:
            pointer1 += 1

        while pointer2 > pointer1 and l[pointer2] not in vowels:
            pointer2 -= 1

        l[pointer1], l[pointer2] = l[pointer2], l[pointer1]
        pointer1 += 1
        pointer2 -= 1

    return ''.join(l)


print(reverse_vowels('coco E cici'))