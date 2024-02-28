/**
 * O(n)
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isSubsequence = function (s, t) {
  let subIdx = 0;

  for (let i = 0; i < t.length; i++) {
    if (t[i] === s[subIdx]) {
      subIdx++;
    }
  }

  return subIdx === s.length;
};

console.log(isSubsequence("axc", "ahbgdc"));
