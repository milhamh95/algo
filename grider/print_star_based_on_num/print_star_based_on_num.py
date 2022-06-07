def printstarbasedonnum(num):
    result = ""
    space = " "
    for i in range(1, num + 1, 1):
        if i == num:
            space = ""

        if i % 5 == 2 or i % 5 == 0:
            result = result + "*" + space
            continue

        result = result + str(i) + space

    return result