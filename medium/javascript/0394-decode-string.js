/**
 * @param {string} s
 * @return {string}
 */
var decodeString = function (s) {
  const stack = [];

  for (let i = 0; i < s.length; i++) {
    if (s.charAt(i) !== "]") {
      stack.push(s.charAt(i));
      continue;
    }

    let str = "";
    let currChar = stack.pop();
    while (currChar !== "[") {
      str = currChar + str;
      currChar = stack.pop();
    }

    let k = "";
    currChar = stack.pop();
    while (!Number.isNaN(Number(currChar))) {
      k = currChar + k;
      currChar = stack.pop();
    }

    stack.push(currChar);
    stack.push(str.repeat(Number(k)));
  }

  return stack.join("");
};
