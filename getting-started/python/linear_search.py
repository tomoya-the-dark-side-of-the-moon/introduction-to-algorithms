import random


def search(data, key):
    result = -1
    for i, d in enumerate(data):
        if d == key:
            result = i
            break

    return result


if __name__ == "__main__":
    data = [random.randint(1, 20) for _ in range(20)]

    print("--- data ---")
    print(data)

    print("input key > ", end="")
    key = int(input())

    result = search(data, key)
    if result != -1:
        print(f"result = {result}")
    else:
        print(f"{key} is not found.")
