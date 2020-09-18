/*
 Note. slice returns a new array
 while pop, shift and splice do not

 ** Arrays
 ** New - = []
 ** Insert (at start) - .unshift(n)
 ** Insert (at end) - .splice(arr.length, 0, n)
 ** Insert (at middle) - .splice(pos, 0, n)
 ** Lookup - [n]
 ** Remove (at start) - .shift()
 ** Remove (at end) - .pop()
 ** Remove (at middle) - .splice(pos, 1)
 ** Slice - .slice(start, end) # start is inclusive and end is exclusive
 */

// =============================================

function generateArray(n) {
  const arr = [];
  for (let i = 0; i < n; i++) {
    arr.push(i);
  }
  return arr;
}

let size = 100000000;

let small = generateArray(size);
console.time('splice');
small.splice(small.length - 1, 1);
console.timeEnd('splice');

small = generateArray(size);
console.time('pop');
small.pop();
console.timeEnd('pop');
