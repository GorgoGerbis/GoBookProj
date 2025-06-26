// arrays_and_functions.monkey - working with arrays and functions

let map = fn(arr, f) {
  let iter = fn(arr, accumulated) {
    if (len(arr) == 0) {
      accumulated
    } else {
      iter(rest(arr), push(accumulated, f(first(arr))));
    }
  };
  
  iter(arr, []);
};

let reduce = fn(arr, initial, f) {
  let iter = fn(arr, result) {
    if (len(arr) == 0) {
      result
    } else {
      iter(rest(arr), f(result, first(arr)));
    }
  };
  
  iter(arr, initial);
};

// example usage
let a = [1, 2, 3, 4];
let double = fn(x) { x * 2 };
let add = fn(x, y) { x + y };

let doubled = map(a, double);     // [2, 4, 6, 8]
let sum = reduce(a, 0, add);      // 10

puts("Original array: ");
puts(a);
puts("Doubled: ");
puts(doubled);
puts("Sum: ");
puts(sum); 