/**
 * O(n) time and O(1) extra space
 * @param {character[]} chars
 * @return {number}
 */
var compress = function (chars) {
  let idx = 0;
  let resultIdx = 0;
  while (idx < chars.length) {
    const char = chars[idx];
    let sequenceLen = 1;

    while (chars[idx + sequenceLen] === char) {
      sequenceLen++;
    }

    idx += sequenceLen;

    chars[resultIdx] = char;
    if (sequenceLen > 1) {
      resultIdx++;
      let lenStr = sequenceLen.toString();
      for (let i = 0; i < lenStr.length; i++) {
        chars[resultIdx] = lenStr[i];
        resultIdx++;
      }
    } else {
      resultIdx++;
    }
  }

  return resultIdx;
};
