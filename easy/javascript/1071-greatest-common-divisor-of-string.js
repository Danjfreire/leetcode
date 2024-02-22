/**
 * n = str1.len
 * m = str2.len
 * O(min(n,m) * (m+n))
 * @param {string} str1
 * @param {string} str2
 * @return {string}
 */
var gcdOfStrings = function (str1, str2) {
  let minStr = str1.length < str2.length ? str1 : str2;

  while (minStr.length > 0) {
    if (isGcd(str1, str2, minStr)) {
      break;
    }

    minStr = minStr.substring(0, minStr.length - 1);
  }

  return minStr;
};

function isGcd(str1, str2, gcd) {
  const len1 = str1.length;
  const len2 = str2.length;
  const gcdLen = gcd.length;

  if (len1 % gcdLen > 0 || len2 % gcdLen > 0) {
    return false;
  }

  const divideStr1 = gcd.repeat(len1 / gcdLen) === str1;
  const divideStr2 = gcd.repeat(len2 / gcdLen) === str2;

  return divideStr1 && divideStr2;
}
