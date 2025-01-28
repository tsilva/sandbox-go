import time

def main():
    # Define the range of numbers to multiply
    num_elements = 10000000

    # Start timer
    start = time.time()

    # Perform the multiplication
    result = 1.0  # Using float to avoid overflow for demonstration
    for i in range(1, num_elements + 1):
        result *= i
        if result == 0.0:
            print("Result too large and became zero due to floating point precision.")
            break

    # Stop timer
    elapsed = time.time() - start
    elapsed *= 1000.0

    # Output the result and time elapsed
    print(f"Time elapsed: {elapsed:.6f}ms")
    print(f"Final result (or partial): {result}")

if __name__ == "__main__":
    main()

