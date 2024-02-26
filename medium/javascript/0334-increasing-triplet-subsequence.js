/**
 * O(n) time and O(1) space
 * @param {number[]} nums
 * @return {boolean}
 */
var increasingTriplet = function (nums) {
  let minNum = Infinity;
  let midNum = Infinity;

  for (let currentNum of nums) {
    if (currentNum > midNum) {
      return true;
    }

    if (currentNum > minNum) {
      midNum = currentNum;
    } else {
      minNum = currentNum;
    }
  }

  return false;
};

console.log(increasingTriplet([1, 5, 2, 6, 3]));
