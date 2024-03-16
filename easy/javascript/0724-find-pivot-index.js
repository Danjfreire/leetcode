/**
 * @param {number[]} nums
 * @return {number}
 */
var pivotIndex = function (nums) {
  let rightSumMap = new Map();
  rightSumMap.set(nums.length - 1, 0);

  for (let i = nums.length - 2; i >= 0; i--) {
    rightSumMap.set(i, rightSumMap.get(i + 1) + nums[i + 1]);
  }

  let pivotIndex = -1;
  let leftSumMap = new Map();

  for (let i = 0; i < nums.length; i++) {
    const leftSum = i == 0 ? 0 : leftSumMap.get(i - 1) + nums[i - 1];

    if (leftSum === rightSumMap.get(i)) {
      pivotIndex = i;
      break;
    } else {
      leftSumMap.set(i, leftSum);
    }
  }

  return pivotIndex;
};
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
/**
 * @param {number[]} nums
 * @return {number}
 */
var pivotIndex = function (nums) {
  let totalSum = 0;
  for (let i = 0; i < nums.length; i++) {
    totalSum += nums[i];
  }

  let leftSum = 0;
  for (let i = 0; i < nums.length; i++) {
    const rightSum = totalSum - nums[i] - leftSum;
    if (rightSum === leftSum) {
      return i;
    } else {
      leftSum += nums[i];
    }
  }

  return -1;
};
