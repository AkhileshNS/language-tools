/*
 ** Stack
 ** New - new Stack(value)
 ** Push - .push(value)
 ** Pop - .pop()
 ** Peek - .peek()
 */

// =============================================

class Stack {
  constructor(value) {
    this.stack = [value];
  }

  peek() {
    if (this.stack.length > 0) {
      return this.stack[this.stack.length - 1];
    } else {
      console.log('Stack is empty!');
      return null;
    }
  }

  push(value) {
    this.stack.push(value);
  }

  pop() {
    if (this.stack.length > 0) {
      return this.stack.pop();
    } else {
      console.log('Stack is empty!');
      return null;
    }
  }
}

const s = new Stack(1);
console.log(s.stack);
s.push(10);
console.log(s.stack);
console.log(s.pop());
console.log(s.peek());
console.log(s.stack);
