def decode(s: str):
    stack = []
    cur_str = ''
    cur_num = ''

    for char in s:
        if char.isdigit():
            cur_num += char
        elif char.isalpha():
            cur_str += char
        elif char == '[':
            stack.append(cur_str)
            stack.append(cur_num)
            cur_str = ''
            cur_num = ''
        elif char == ']':
            prev_num = int(stack.pop())
            prev_str = stack.pop()
            cur_str = prev_str + cur_str * prev_num

    return cur_str


print(decode('a1[b2[c]d12[e]]f0[z]'))
