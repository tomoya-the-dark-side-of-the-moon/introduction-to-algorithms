def select_items(weights, prices, limit=50):
    unit_prices = {}
    for i, (w, p) in enumerate(zip(weights, prices)):
        unit_prices[i] = p / w

    sorted_unit_prices = dict(
        sorted(unit_prices.items(), key=lambda x: x[1], reverse=True)
    )

    rest = limit
    items = {}
    for idx in sorted_unit_prices:
        items[idx] = weights[idx] if rest - weights[idx] > 0 else rest

        rest -= items[idx]

    return items


if __name__ == "__main__":
    weights = [20, 30, 10]
    prices = [100, 120, 60]

    print("  i |", end="")
    for i in range(len(weights)):
        print(f"{i:^3}|", end="")
    print()
    print("w[i]|", end="")
    for w in weights:
        print(f"{w:>3}|", end="")
    print()
    print("p[i]|", end="")
    for p in prices:
        print(f"{p:>3}|", end="")
    print()

    items = select_items(weights, prices)

    print("--- result ---")
    total = 0.0
    for idx in items:
        print(f"index = {idx}, weight = {items[idx]}")

        total += (prices[idx] / weights[idx]) * items[idx]
    print(f"total = {total}")
