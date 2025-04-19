# The function that prints "hello"
def print_hello():
    print("hello")

# Test function to capture printed output and compare it
def test_print_hello(capsys):
    print_hello()  # Call the function that prints "hello"
    
    # Capture the printed output
    captured = capsys.readouterr()
    
    # Check if the output is "hello"
    assert captured.out.strip() == "hello"
