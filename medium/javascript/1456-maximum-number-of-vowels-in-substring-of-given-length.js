/**
 * @param {string} s
 * @param {number} k
 * @return {number}
 */
var maxVowels = function (s, k) {
  let currentVowels = 0;
  const vowels = new Set(["a", "e", "i", "o", "u"]);

  for (let idx = 0; idx < k; idx++) {
    if (vowels.has(s[idx])) {
      currentVowels++;
    }
  }

  max = currentVowels;

  for (let idx = k; idx < s.length; idx++) {
    if (vowels.has(s[idx])) {
      currentVowels++;
    }

    if (vowels.has(s[idx - k])) {
      currentVowels--;
    }

    if (currentVowels > max) {
      max = currentVowels;
    }
  }

  return max;
};
