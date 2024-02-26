/**
 * O(n^2)
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  const result = new Array(nums.length);

  for (let i = 0; i < result.length; i++) {
    result[i] = 1;
    for (let j = 0; j < nums.length; j++) {
      if (j !== i) {
        result[i] *= nums[j];
      }
    }
  }

  return result;
};

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * O(n)
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  let totalProduct = 1;
  for (const num of nums) {
    totalProduct *= num;
  }

  const result = new Array(nums.length);

  for (let i = 0; i < result.length; i++) {
    result[i] = totalProduct / nums[i];
  }

  return result;
};

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * O(n) without division
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  const leftProduct = new Array(nums.length);

  for (let i = 0; i < leftProduct.length; i++) {
    if (i === 0) {
      leftProduct[i] = 1;
    } else {
      leftProduct[i] = nums[i - 1] * leftProduct[i - 1];
    }
  }

  const rightProduct = new Array(nums.length);

  for (let i = nums.length - 1; i >= 0; i--) {
    if (i === nums.length - 1) {
      rightProduct[i] = 1;
    } else {
      rightProduct[i] = nums[i + 1] * rightProduct[i + 1];
    }
  }

  const result = new Array(nums.length);

  for (let i = 0; i < result.length; i++) {
    result[i] = rightProduct[i] * leftProduct[i];
  }

  return result;
};

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * O(n) without division and with O(1) extra space
 * @param {number[]} nums
 * @return {number[]}
 */
var productExceptSelf = function (nums) {
  const result = new Array(nums.length);
  let leftProduct = 1;

  for (let i = 0; i < nums.length; i++) {
    result[i] = leftProduct;
    leftProduct *= nums[i];
  }

  let rightProduct = 1;

  for (let i = nums.length - 1; i >= 0; i--) {
    result[i] *= rightProduct;
    rightProduct *= nums[i];
  }

  return result;
};

console.log(productExceptSelf([1, 2, 3, 4]));
