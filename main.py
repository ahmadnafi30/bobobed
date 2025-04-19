# main.py

def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def string_to_uppercase(s):
    return s.upper()

if __name__ == "__main__":
    print("Penjumlahan 3 + 2 =", add(3, 2))
    print("Pengurangan 5 - 3 =", subtract(5, 3))
    print('Uppercase dari "hello" =', string_to_uppercase("hello"))
