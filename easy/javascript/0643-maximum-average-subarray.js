/** O(n)
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function (nums, k) {
  let currentSum = 0;
  let topSum = -Infinity;

  for (let idx = 0; idx < k; idx++) {
    currentSum += nums[idx];
    topSum = currentSum;
  }

  for (let idx = k; idx < nums.length; idx++) {
    currentSum -= nums[idx - k];
    currentSum += nums[idx];

    if (currentSum > topSum) {
      topSum = currentSum;
    }
  }

  return topSum / k;
};
