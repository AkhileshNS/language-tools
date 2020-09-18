/*
 Note. All functions are in-order

 ** Doubly Linked List 
 ** New - new LinkedList(value)
 ** Insert (at start) - .prepend(value)
 ** Insert (at end) - .append(value)
 ** Insert (at middle) - .insert(index, value)
 ** Lookup - .get(index)
 ** Remove - .remove(index)
 */

// =============================================

class LinkedList {
  constructor(value) {
    this.initialize(value);
  }

  initialize = (value) => {
    this.head = { value, next: null, prev: null };
    this.tail = this.head;
    this.length = 1;
  };

  isEmpty = () => {
    if (!this.head) {
      console.log('List is empty');
      return true;
    }
    return false;
  };

  append = (value) => {
    if (this.isEmpty()) {
      this.initialize(value);
      return this;
    }

    const node = { value, next: null, prev: null };
    this.tail.next = node;
    node.prev = this.tail;
    this.tail = node;
    this.length++;
    return this;
  };

  prepend = (value) => {
    if (this.isEmpty()) {
      this.initialize(value);
      return this;
    }

    const node = { value, next: null, prev: null };
    node.next = this.head;
    this.head.prev = node;
    this.head = node;
    this.length++;
    return this;
  };

  print = () => {
    if (!this.isEmpty()) {
      let curr = this.head;
      let str = curr.value + '';
      while (curr.next != null) {
        curr = curr.next;
        str += ` <-> ${curr.value}`;
      }
      console.log(str);
    }
  };

  printBackwards = () => {
    if (!this.isEmpty()) {
      let curr = this.tail;
      let str = curr.value + '';
      while (curr.prev != null) {
        curr = curr.prev;
        str += ` <-> ${curr.value}`;
      }
      console.log(str);
    }
  };

  traverse = (index) => {
    if (this.isEmpty()) {
      return { follower: null, curr: null, leader: null };
    }

    let curr = this.head;
    if (index <= this.length / 2) {
      for (let i = 0; i < index && i < this.length - 1; i++) {
        curr = curr.next;
      }
    } else {
      curr = this.tail;
      for (let i = this.length - 1; i > index && i > 0; i--) {
        curr = curr.prev;
      }
    }

    return { follower: curr.prev, curr, leader: curr.next };
  };

  get = (index) => {
    if (!this.isEmpty()) {
      return this.traverse(index).curr.value;
    }
    return null;
  };

  insert = (index, value) => {
    if (this.isEmpty()) {
      this.initialize(value);
      return this;
    }
    const node = { value, next: null };

    const { _, curr, leader } = this.traverse(index);
    curr.next = node;
    node.prev = curr;
    node.next = leader;
    if (leader) {
      leader.prev = node;
    } else {
      this.tail = node;
    }
    this.length++;
    return this;
  };

  remove = (index) => {
    if (!this.isEmpty()) {
      const { follower, curr, leader } = this.traverse(index);
      if (!follower) {
        this.head = leader;
        leader.prev = null;
        curr.next = null;
      } else if (!leader) {
        this.tail = follower;
        follower.next = null;
        curr.prev = null;
      } else {
        follower.next = leader;
        leader.prev = follower;
        curr.next = null;
        curr.prev = null;
      }
      this.length--;
    }

    return this;
  };
}

const ll = new LinkedList(5);
ll.print();
ll.printBackwards();
ll.append(15);
ll.print();
ll.printBackwards();
ll.prepend(11);
ll.print();
ll.printBackwards();
