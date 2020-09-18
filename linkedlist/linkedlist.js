/*
 Note. All functions are in-order

 ** Singly Linked List 
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
    this.head = { value, next: null };
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

    const node = { value, next: null };
    this.tail.next = node;
    this.tail = node;
    this.length++;
    return this;
  };

  prepend = (value) => {
    if (this.isEmpty()) {
      this.initialize(value);
      return this;
    }

    const node = { value, next: null };
    node.next = this.head;
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
        str += ` --> ${curr.value}`;
      }
      console.log(str);
    }
  };

  traverse = (index) => {
    if (this.isEmpty()) {
      return { follower: null, curr: null, leader: null };
    }

    let follower = null;
    let curr = this.head;

    for (let i = 0; i < index && i < this.length - 1; i++) {
      follower = curr;
      curr = curr.next;
    }

    return { follower, curr, leader: curr ? curr.next : null };
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
    node.next = leader;
    this.length++;
    return this;
  };

  remove = (index) => {
    if (!this.isEmpty()) {
      const { follower, curr, leader } = this.traverse(index);
      if (!follower) {
        this.head = leader;
        curr.next = null;
      } else if (!leader) {
        this.tail = follower;
        follower.next = null;
      } else {
        follower.next = leader;
        curr.next = null;
      }
    }

    return this;
  };

  reverse = () => {
    if (!this.isEmpty() && this.length > 1) {
      if (this.length === 2) {
        this.tail = this.head;
        this.head = this.tail.next;
        this.head.next = this.tail;
        this.tail.next = null;
        return;
      }

      const holder = this.head;
      this.head = this.tail;
      this.tail = holder;

      let follower = this.tail;
      let curr = follower.next;

      while (curr !== null) {
        const leader = curr.next;
        curr.next = follower;
        follower = curr;
        curr = leader;
      }

      this.tail.next = null;
    }
  };
}

const ll = new LinkedList(1);
ll.append(2);
ll.append(3);
ll.append(4);
ll.append(5);
ll.print();
ll.reverse();
ll.print();
