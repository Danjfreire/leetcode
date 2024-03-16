/**
 * @param {number[]} arr
 * @return {boolean}
 */
var uniqueOccurrences = function (arr) {
  const occurrenceMap = new Map();

  for (let i = 0; i < arr.length; i++) {
    const occurrences = occurrenceMap.get(arr[i]);

    if (occurrences) {
      occurrenceMap.set(arr[i], occurrences + 1);
    } else {
      occurrenceMap.set(arr[i], 1);
    }
  }

  const occurrenceSet = new Set();

  for (const [_, v] of occurrenceMap) {
    if (occurrenceSet.has(v)) {
      return false;
    } else {
      occurrenceSet.add(v);
    }
  }

  return true;
};
