/**
 * O(n)
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var maxOperations = function (nums, k) {
  const map = new Map();
  let operations = 0;

  for (let i = 0; i < nums.length; i++) {
    const dif = k - nums[i];
    const mappedDif = map.get(dif);

    if (mappedDif !== undefined && mappedDif.length > 0) {
      operations++;
      mappedDif.pop();
    } else {
      const mapped = map.get(nums[i]);
      mapped ? mapped.push(i) : map.set(nums[i], [i]);
    }
  }

  return operations;
};
