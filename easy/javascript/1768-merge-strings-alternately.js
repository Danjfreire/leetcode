/**
 * O(n) where n = max length between words
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function (word1, word2) {
  let result = "";
  const maxLength = Math.max(word1.length, word2.length);

  for (let i = 0; i < maxLength; i++) {
    if (i < word1.length) {
      result += word1[i];
    }

    if (i < word2.length) {
      result += word2[i];
    }
  }

  return result;
};

///////////////////////////////////////////////////////////////////////////////////////////

/**
 * O(n) where n = min length between words, also less constant operations
 * @param {string} word1
 * @param {string} word2
 * @return {string}
 */
var mergeAlternately = function (word1, word2) {
  let result = "";
  const minLength = Math.min(word1.length, word2.length);

  for (let i = 0; i < minLength; i++) {
    result += word1[i];
    result += word2[i];
  }

  if (word1.length === minLength) {
    result += word2.substring(minLength);
  } else {
    result += word1.substring(minLength);
  }

  return result;
};
