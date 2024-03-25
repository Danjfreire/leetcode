/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteMiddle = function (head) {
  let length = 1;
  let cur = head;

  while (cur.next != undefined) {
    cur = cur.next;
    length++;
  }

  const middle = Math.floor(length / 2);

  if (middle === 0) {
    return null;
  }

  cur = head;
  let idx = 0;

  while (idx !== middle - 1 && cur.next != undefined) {
    cur = cur.next;
    idx++;
  }

  const removed = cur.next;
  cur.next = removed.next;
  removed.next = undefined;

  return head;
};
