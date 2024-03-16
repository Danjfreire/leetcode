/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[][]}
 */
var findDifference = function (nums1, nums2) {
  let onlyInNums1 = [];
  let onlyInNums2 = [];

  const num1Set = new Set(nums1);
  const num2Set = new Set(nums2);

  for (const v of num1Set) {
    if (!num2Set.has(v)) {
      onlyInNums1.push(v);
    }
  }

  for (const v of num2Set) {
    if (!num1Set.has(v)) {
      onlyInNums2.push(v);
    }
  }

  return [onlyInNums1, onlyInNums2];
};
