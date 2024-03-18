/**
 * @param {number[][]} grid
 * @return {number}
 */
var equalPairs = function (grid) {
  const rowMap = new Map();
  const n = grid.length;

  for (let i = 0; i < n; i++) {
    let rowStr = "";
    for (let j = 0; j < n; j++) {
      rowStr = rowStr.concat(`|${grid[i][j]}|`);
    }

    const occurrences = rowMap.get(rowStr);
    if (occurrences) {
      rowMap.set(rowStr, occurrences + 1);
    } else {
      rowMap.set(rowStr, 1);
    }
  }

  let equalPairs = 0;

  for (let i = 0; i < n; i++) {
    let colStr = "";
    for (let j = 0; j < n; j++) {
      colStr = colStr.concat(`|${grid[j][i]}|`);
    }
    const occurrences = rowMap.get(colStr);
    if (occurrences) {
      equalPairs += occurrences;
    }
  }

  return equalPairs;
};
