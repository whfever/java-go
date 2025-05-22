# ====== 1. Basic Function ======
def greet(name="Anonymous"):
    """Demonstrates function definition and docstring"""
    return f"Hello, {name}!"


# ====== 2. Parameter Types ======
def calculate(a, b=5, *args, **kwargs):
    """Shows multiple parameter types:
    - Positional (a)
    - Default (b)
    - Arbitrary positional (*args)
    - Arbitrary keyword (**kwargs)
    """
    result = a + b
    for arg in args:
        result += arg
    for key in kwargs:
        result += kwargs[key]
    return result


# ====== 3. Return Multiple Values ======
def circle_calc(radius):
    """Returns multiple values as tuple"""
    area = 3.14159 * radius**2
    circumference = 2 * 3.14159 * radius
    return area, circumference


# ====== 4. Lambda Functions ======
square = lambda x: x ** 2
add_numbers = lambda a, b: a + b


# ====== 5. Function Scope ======
global_var = 100

def scope_demo():
    local_var = 50
    global global_var
    global_var += 1
    return local_var, global_var


# ====== 6. Decorators ======
def simple_decorator(func):
    def wrapper():
        print("Before function call")
        func()
        print("After function call")
    return wrapper

@simple_decorator
def say_hello():
    print("Hello from decorated function!")


# ====== 7. Generator Functions ======
def fibonacci_gen(n):
    """Generates Fibonacci sequence"""
    a, b = 0, 1
    for _ in range(n):
        yield a
        a, b = b, a + b


# ====== 8. Recursive Functions ======
def factorial(n):
    if n == 0:
        return 1
    return n * factorial(n-1)


# ====== 9. Type Hints ======
def type_hinted_func(name: str, age: int) -> str:
    return f"{name} is {age} years old"


# ====== Demonstration Block ======
if __name__ == "__main__":
    # 1. Basic
    print(greet("Alice"))  # Hello, Alice!
    
    # 2. Parameters
    print(calculate(2))               # 7 (2+5)
    print(calculate(2, 3, 4, 5, x=6)) # 20 (2+3+4+5+6)
    
    # 3. Multiple returns
    area, circ = circle_calc(5)
    print(f"Area: {area:.2f}, Circumference: {circ:.2f}")
    
    # 4. Lambda
    print(square(4))      # 16
    print(add_numbers(3,5)) # 8
    
    # 5. Scope
    print(scope_demo())  # (50, 101)
    
    # 6. Decorator
    say_hello()
    
    # 7. Generator
    print(list(fibonacci_gen(5)))  # [0, 1, 1, 2, 3]
    
    # 8. Recursion
    print(factorial(5))  # 120
    
    # 9. Type hints
    print(type_hinted_func("Bob", 30))