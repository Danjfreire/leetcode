/**
 * O(n)
 * @param {number[]} flowerbed
 * @param {number} n
 * @return {boolean}
 */
var canPlaceFlowers = function (flowerbed, n) {
  let plantedFlowers = 0;

  for (let i = 0; i < flowerbed.length; i++) {
    if (flowerbed[i] === 1) {
      continue;
    }

    const hasLeft = i > 0 && flowerbed[i - 1] === 1;
    const hasRight = i < flowerbed.length - 1 && flowerbed[i + 1] === 1;

    if (!hasLeft && !hasRight) {
      flowerbed[i] = 1;
      plantedFlowers++;

      // early return here for a slighly more optimized implementation
    }
  }

  return plantedFlowers >= n;
};

console.log(canPlaceFlowers([1, 0, 0, 0, 1], 2));
