/**
 * @param {number[]} nums
 * @return {number}
 */
var longestSubarray = function (nums) {
  let left = 0;
  let right = 0;
  let lastZeroIdx = -1;
  let maxLen = 0;

  while (right < nums.length) {
    if (nums[right] === 0 && lastZeroIdx === -1) {
      lastZeroIdx = right;
    } else if (nums[right] === 0) {
      maxLen = Math.max(maxLen, right - 1 - left); // don't count the current value
      left = lastZeroIdx + 1;
      lastZeroIdx = right;
    } else {
      maxLen = Math.max(maxLen, right - left);
    }
    right++;
  }

  return maxLen;
};
