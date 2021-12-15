# Task 2)

"""
take user input as string
convert string to int
write fibonacci sequence till that
"""

x = int(input("Please enter your destination number:"))

a, b = 0, 1

def main():
    global a, b  # this global tells python to refer global variables created outside the func
    for i in range(0,x,1):
        print(a)
        a, b = b, a+b
main()
