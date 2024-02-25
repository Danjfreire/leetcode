/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function (s) {
  const words = s.split(/\s+/);
  let result = "";

  for (let i = words.length - 1; i >= 0; i--) {
    result += words[i];
    result += " ";
  }

  return result.trim();
};
