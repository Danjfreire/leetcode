/**
 * O(n * 10) => O(n)
 * @param {string} s
 * @return {string}
 */
var reverseVowels = function (s) {
  const vowels = ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"];
  const strVowels = [];

  for (const letter of s) {
    if (vowels.includes(letter)) {
      strVowels.push(letter);
    }
  }

  let result = "";

  for (let i = 0; i < s.length; i++) {
    if (vowels.includes(s[i])) {
      result += strVowels.pop();
    } else {
      result += s[i];
    }
  }

  return result;
};
