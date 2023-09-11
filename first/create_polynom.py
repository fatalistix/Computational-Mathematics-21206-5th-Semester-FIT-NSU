a_buf, b_buf, c_buf = map(str, input().split())

a = a_buf[:-1]
b = b_buf[:-1]
c = c_buf[:-1]

print("(x^3)+(", a, "*x^2)+(", b, "*x)+(", c, ")", sep="")
