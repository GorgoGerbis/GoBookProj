// fibonacci.monkey - recursive fibonacci implementation

let fibonacci = fn(x) {
  if (x == 0) {
    0
  } else {
    if (x == 1) {
      1
    } else {
      fibonacci(x - 1) + fibonacci(x - 2);
    }
  }
};

// test some values
let result1 = fibonacci(10);  // should be 55
let result2 = fibonacci(15);  // should be 610

// show the results
puts("fibonacci(10) = ");
puts(result1);
puts("fibonacci(15) = ");
puts(result2); 