/**
 * O(n)
 * @param {number[]} candies
 * @param {number} extraCandies
 * @return {boolean[]}
 */
var kidsWithCandies = function (candies, extraCandies) {
  const result = new Array(candies.length).fill(false);
  const maxCandies = candies.reduce((prev, curr) => {
    return curr > prev ? curr : prev;
  }, 0);

  for (let i = 0; i < candies.length; i++) {
    if (candies[i] + extraCandies >= maxCandies) {
      result[i] = true;
    }
  }

  return result;
};

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * O(n) - more javascripty
 * @param {number[]} candies
 * @param {number} extraCandies
 * @return {boolean[]}
 */
var kidsWithCandies = function (candies, extraCandies) {
  const maxCandies = Math.max(...candies);
  return candies.map((candy) => candy + extraCandies >= maxCandies);
};
