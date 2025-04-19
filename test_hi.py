def print_hi():
    print("hi")

# Test function to capture printed output and compare it
def test_print_hi(capsys):
    print_hi()  # Call the function that prints "hello"
    
    # Capture the printed output
    captured = capsys.readouterr()
    
    # Check if the output is "hello"
    assert captured.out.strip() == "hi"
