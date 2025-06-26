// closures.monkey - demonstrating closures and lexical scoping

let makeAdder = fn(x) {
  fn(y) { x + y };
};

let add5 = makeAdder(5);
let add10 = makeAdder(10);

puts("add5(3) = ");
puts(add5(3));    // should be 8

puts("add10(3) = ");
puts(add10(3));   // should be 13

// counter example
let makeCounter = fn() {
  let count = 0;
  fn() {
    count = count + 1;
    count;
  };
};

let counter = makeCounter();
puts("First call: ");
puts(counter());  // 1
puts("Second call: ");
puts(counter()); // 2
puts("Third call: ");
puts(counter()); // 3 