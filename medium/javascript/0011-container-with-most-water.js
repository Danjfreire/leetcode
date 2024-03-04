/**
 * O(n^2)
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let max = 0;

  for (let i = 0; i < height.length; i++) {
    for (let j = i + 1; j < height.length; j++) {
      if (i === j) {
        continue;
      }

      const area = Math.min(height[i], height[j]) * (j - i);
      if (area > max) {
        max = area;
      }
    }
  }

  return max;
};

/**
 * O(n)
 * @param {number[]} height
 * @return {number}
 */
var maxArea = function (height) {
  let left = 0;
  let right = height.length - 1;
  let max = -Infinity;
  while (left < right) {
    const area = Math.min(height[left], height[right]) * (right - left);
    max = Math.max(max, area);
    if (height[left] > height[right]) {
      right--;
    } else {
      left++;
    }
  }
  return max;
};
