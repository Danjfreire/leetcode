/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var longestOnes = function (nums, k) {
  let maxOnes = 0;
  let currentOnes = 0;
  let leftIdx = 0;
  let rightIdx = 0;
  let flippedZeroes = 0;

  while (rightIdx < nums.length) {
    if (nums[rightIdx] === 1) {
      currentOnes++;
      rightIdx++;
    } else {
      if (flippedZeroes < k) {
        flippedZeroes++;
        rightIdx++;
        currentOnes++;
      } else {
        currentOnes--;
        if (nums[leftIdx] === 0) {
          flippedZeroes--;
        }
        leftIdx++;
      }
    }

    if (currentOnes > maxOnes) {
      maxOnes = currentOnes;
    }
  }

  return maxOnes;
};

var longestOnes = function (nums, k) {
  let left = 0;
  let right = 0;

  while (right < nums.length) {
    if (nums[right] === 0) {
      k--;
    }
    if (k < 0) {
      if (nums[left] === 0) {
        k++;
      }
      left++;
    }
    right++;
  }
  return right - left;
};

longestOnes([0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3);
