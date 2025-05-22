# ====== 1. Basic Class ======
class Animal:
    """Base class demonstrating class basics"""
    
    def __init__(self, name: str):
        self.name = name
        
    def speak(self) -> str:
        return "Generic animal sound"

# ====== 2. Inheritance ======
class Dog(Animal):
    """Child class overriding methods"""
    
    def speak(self) -> str:
        return f"{self.name} says: Woof!"

# ====== 3. Class Attributes & Static Methods ======
class MathOperations:
    PI = 3.14159  # Class attribute
    
    @staticmethod
    def add(a: int, b: int) -> int:
        return a + b

# ====== 4. Property Decorators ======
class Circle:
    def __init__(self, radius: float):
        self._radius = radius
        
    @property
    def radius(self) -> float:
        return self._radius
        
    @radius.setter
    def radius(self, value: float):
        if value > 0:
            self._radius = value
            
    @property
    def area(self) -> float:
        return self._radius ** 2 * MathOperations.PI

# ====== 5. Polymorphism ======
class Cat(Animal):
    def speak(self) -> str:
        return f"{self.name} says: Meow!"

# ====== 6. Magic Methods ======
class Vector:
    def __init__(self, x: float, y: float):
        self.x = x
        self.y = y
        
    def __add__(self, other):
        return Vector(self.x + other.x, self.y + other.y)
        
    def __str__(self):
        return f"Vector({self.x}, {self.y})"

# ====== 7. Abstract Base Class ======
from abc import ABC, abstractmethod

class Shape(ABC):
    @abstractmethod
    def area(self) -> float:
        pass

class Square(Shape):
    def __init__(self, side: float):
        self.side = side
        
    def area(self) -> float:
        return self.side ** 2

# ====== Demonstration Block ======
if __name__ == "__main__":
    # 1. Basic class
    generic = Animal("Generic")
    print(generic.speak())  # Generic animal sound
    
    # 2. Inheritance
    dog = Dog("Buddy")
    print(dog.speak())      # Buddy says: Woof!
    
    # 3. Class attributes
    print(MathOperations.PI)           # 3.14159
    print(MathOperations.add(2, 3))    # 5
    
    # 4. Property decorators
    c = Circle(5)
    print(c.area)           # 78.53975
    c.radius = 10
    print(c.area)           # 314.159
    
    # 5. Polymorphism
    cat = Cat("Whiskers")
    print(cat.speak())      # Whiskers says: Meow!
    
    # 6. Magic methods
    v1 = Vector(2, 3)
    v2 = Vector(1, 4)
    print(v1 + v2)          # Vector(3, 7)
    
    # 7. Abstract class
    square = Square(4)
    print(square.area())    # 16