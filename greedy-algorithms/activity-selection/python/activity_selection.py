def select_activities(s, f):
    results = [0]

    i = 1
    finished = f[0]
    while i < len(s):
        if finished <= s[i]:
            results.append(i)
            finished = f[i]

        i += 1

    return results


if __name__ == "__main__":
    s = [1, 3, 0, 5, 3, 5, 6, 7, 8, 2, 12]
    f = [4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16]

    print("  i |", end="")
    for i in range(len(s)):
        print(f"{i:>3}|", end="")
    print()
    print("s[i]|", end="")
    for _s in s:
        print(f"{_s:>3}|", end="")
    print()
    print("f[i]|", end="")
    for _f in f:
        print(f"{_f:>3}|", end="")
    print()

    results = select_activities(s, f)

    print("--- result ---")
    for i in results:
        print(f"activity = {i} / start = {s[i]}, finish = {f[i]}")
