/*
 ** Queue
 ** New - new Queue(value)
 ** Enqueue - .enqueue(value)
 ** Dequeue - .dequeue()
 ** Peek - .peek()
 */

// =============================================
class Queue {
  constructor(value) {
    this.initialize(value);
  }

  initialize(value) {
    this.front = { value, next: null };
    this.back = this.front;
    this.length = 1;
  }

  isEmpty() {
    if (!this.front) {
      console.log('Queue is empty!');
      return true;
    }
    return false;
  }

  enqueue(value) {
    if (this.isEmpty()) {
      this.initialize(value);
      return this;
    }

    const node = { value, next: null };
    this.back.next = node;
    this.back = node;
    this.length++;
    return this;
  }

  dequeue() {
    if (this.isEmpty()) {
      return null;
    }
    const node = this.front;
    this.front = this.front.next;
    node.next = null;
    this.length--;
    return node;
  }

  peek() {
    if (!this.isEmpty()) {
      return this.front.value;
    } else {
      return null;
    }
  }

  print() {
    if (!this.isEmpty()) {
      let curr = this.front;
      let str = this.front.value + '';
      curr = curr.next;
      while (curr) {
        str += ` <-- ${curr.value}`;
        curr = curr.next;
      }
      console.log(str);
    }
  }
}

const q = new Queue('Akhilesh');
q.print();
q.enqueue('Anirudh');
q.enqueue('Anirban');
q.print();
console.log(q.peek());
console.log(q.dequeue());
q.print();
