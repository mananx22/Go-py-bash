# Task 1) make a variable and echo it out while simultaneously commenting out this line of comment.

"""
make single line comment
make a block comment
make a variable
print the variable on screen
make a global variable
overrride global variable from local variable
"""

x = "manan"

def main():
    x = "sankhla"
    print("Hello, World insidefunc ! -> " + x)

main()
print("Hello, World outsidefunc ! -> " + x)