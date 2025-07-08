def validate(s: str) -> bool:
    mapper = {
        ')': '(',
        '}': '{',
        ']': '[',
    }
    stack = []
    for char in s:
        if char in '({[':
            stack.append(char)

        if char in mapper and stack and stack[-1] == mapper[char]:
            stack.pop()

    if stack:
        return False

    return True


print(validate('(( { [ ] } { } ))'))