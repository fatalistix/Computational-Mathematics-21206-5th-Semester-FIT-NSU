import math


def foo(x):
    return x * x - 5


def diffoo(x):
    return 2 * x


if __name__ == "__main__":
    singleTangenrX = [3]
    newtonX = [3]
    secantsX = [3]
    target = math.sqrt(5)

    for i in range(5):
        print("<====================================================>")
        print(singleTangenrX[i], math.fabs(singleTangenrX[i] - target))
        print(newtonX[i], math.fabs(newtonX[i] - target))
        print(secantsX[i], math.fabs(secantsX[i] - target))

        singleTangenrX.append(
            singleTangenrX[i] - foo(singleTangenrX[i]) / diffoo(singleTangenrX[0])
        )
        newtonX.append(newtonX[i] - foo(newtonX[i]) / diffoo(newtonX[i]))

        if i == 0:
            secantsX.append(secantsX[i] - foo(secantsX[i]) / diffoo(secantsX[i]))
        else:
            secantsX.append(
                secantsX[i]
                - foo(secantsX[i])
                / (foo(secantsX[i - 1]) - foo(secantsX[i]))
                * (secantsX[i - 1] - secantsX[i])
            )
