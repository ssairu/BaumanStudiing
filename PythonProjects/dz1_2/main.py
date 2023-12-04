def is_valid(candidate, group, matrix):
    for c in group:
        if matrix[candidate][c] == '+':
            return False
    return True

def find_groups(matrix, group1, group2, candidates):
    if not candidates:
        return group1, group2

    candidate = candidates.pop()
    if is_valid(candidate, group1, matrix):
        group1.append(candidate)
        result1, result2 = find_groups(matrix, group1, group2, candidates)
        group1.pop()
    else:
        result1 = None

    if is_valid(candidate, group2, matrix):
        group2.append(candidate)
        result3, result4 = find_groups(matrix, group1, group2, candidates)
        group2.pop()
    else:
        result3 = None

    if result1 is None and result3 is None:
        return None, None
    elif result1 is None:
        return result3
    elif result3 is None:
        return result1

    # Compare the sizes of the groups and return the one with the minimum size
    if len(result1) <= len(result3):
        return result1
    else:
        return result3


def main():
    N = int(input("Введите количество кандидатов (N): "))

    # Считываем матрицу
    matrix = []
    for _ in range(N):
        row = input().strip()
        matrix.append(list(row))

    # Создаем список кандидатов и вызываем функцию для поиска групп
    candidates = list(range(N))
    group1, group2 = [], []
    result_group1, result_group2 = find_groups(matrix, group1, group2, candidates)

    if result_group1 is None or result_group2 is None:
        print("No solution")
    else:
        print("Кандидаты первой группы:", result_group1)
        print("Кандидаты второй группы:", result_group2)


if __name__ == "__main__":
    main()