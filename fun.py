# Function definition with default parameters
def greet(name="World", message="Hello"):
    # Function with multiple parameters
    print(f"{message}, {name}!")

# Calling function without parameters - uses default values
greet()

# Calling function with positional arguments
greet("Alice", "Hi")

# Calling function with keyword arguments
greet(message="Good morning", name="Bob")

# Function with variable number of arguments
def sum_all(*args):
    total = 0
    for num in args:
        total += num
    return total

print("Sum of 1, 2, 3:", sum_all(1, 2, 3))

# Lambda expression to create an anonymous function
multiply = lambda x, y: x * y
print("Multiply 4 and 5:", multiply(4, 5))