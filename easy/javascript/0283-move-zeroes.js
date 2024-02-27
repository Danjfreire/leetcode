// /**
//  * sorting - O(n^2)
//  * @param {number[]} nums
//  * @return {void} Do not return anything, modify nums in-place instead.
//  */
// var moveZeroes = function (nums) {
//   for (let i = 0; i < nums.length; i++) {
//     if (nums[i] !== 0) {
//       continue;
//     }

//     for (let j = i + 1; j < nums.length; j++) {
//       if (nums[j] !== 0) {
//         const temp = nums[j];
//         nums[j] = nums[i];
//         nums[i] = temp;
//         break;
//       }
//     }
//   }
// };

/**
 * O(n)
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var moveZeroes = function (nums) {
  let lastNonZeroIdx = 0;

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== 0) {
      nums[lastNonZeroIdx] = nums[i];
      lastNonZeroIdx++;
    }
  }

  for (let i = lastNonZeroIdx; i < nums.length; i++) {
    nums[i] = 0;
  }
};
