/**
 * @param {number[]} asteroids
 * @return {number[]}
 */
var asteroidCollision = function (asteroids) {
  const stack = [];

  for (let i = 0; i < asteroids.length; i++) {
    if (stack.length === 0 || asteroids[i] > 0) {
      stack.push(asteroids[i]);
      continue;
    }

    // just for asteroids goind left
    while (true) {
      const lastAsteroid = stack[stack.length - 1];
      if (lastAsteroid < 0) {
        // the last asteroid was also going left
        stack.push(asteroids[i]);
        break;
      } else if (lastAsteroid === -asteroids[i]) {
        // the last asteroid is the same size and going right
        stack.pop();
        break;
      } else if (lastAsteroid > -asteroids[i]) {
        // the last asteroid is bigger and going right
        break;
      } else {
        // the last asteroid is smaller and going right
        stack.pop();
        if (stack.length === 0) {
          // add it to the stack if there are no more remaining asteroids to collide
          stack.push(asteroids[i]);
          break;
        }
      }
    }
  }

  return stack;
};
