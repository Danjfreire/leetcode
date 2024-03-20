/**
 * @param {string} senate
 * @return {string}
 */
var predictPartyVictory = function (senate) {
  const queue = [];

  for (let i = 0; i < senate.length; i++) {
    queue.push(senate[i]);
  }

  const direQueue = [];
  const radiantQueue = [];

  while (true) {
    if (direQueue.length >= queue.length) {
      return "Dire";
    } else if (radiantQueue.length >= queue.length) {
      return "Radiant";
    }
    const current = queue.shift();

    if (current === "R") {
      if (direQueue.shift() === undefined) {
        radiantQueue.push("R");
        queue.push("R");
      }
    } else if (current === "D") {
      if (radiantQueue.shift() === undefined) {
        direQueue.push("D");
        queue.push("D");
      }
    }
  }
};

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * @param {string} senate
 * @return {string}
 */
var predictPartyVictory = function (senate) {
  const queue = [];

  for (let i = 0; i < senate.length; i++) {
    queue.push(senate[i]);
  }

  let direQueued = 0;
  let radiantQueued = 0;

  while (true) {
    if (direQueued >= queue.length) {
      return "Dire";
    } else if (radiantQueued >= queue.length) {
      return "Radiant";
    }
    const current = queue.shift();

    if (current === "R") {
      if (direQueued === 0) {
        radiantQueued++;
        queue.push("R");
      } else {
        direQueued--;
      }
    } else if (current === "D") {
      if (radiantQueued === 0) {
        direQueued++;
        queue.push("D");
      } else {
        radiantQueued--;
      }
    }
  }
};
