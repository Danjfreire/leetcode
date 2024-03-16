/**
 * @param {string} word1
 * @param {string} word2
 * @return {boolean}
 */
var closeStrings = function (word1, word2) {
  if (word1.length !== word2.length) {
    return false;
  }

  const letters = new Set();
  const word1OccurrenceMap = new Map();
  const word2OccurrenceMap = new Map();

  for (let i = 0; i < word1.length; i++) {
    const word1Occurrence = word1OccurrenceMap.get(word1[i]);
    const word2Occurrence = word2OccurrenceMap.get(word2[i]);

    const word1NewOccurrence = word1Occurrence ? word1Occurrence + 1 : 1;
    const word2NewOccurrence = word2Occurrence ? word2Occurrence + 1 : 1;

    word1OccurrenceMap.set(word1[i], word1NewOccurrence);
    word2OccurrenceMap.set(word2[i], word2NewOccurrence);

    letters.add(word1[i]);
    letters.add(word2[i]);
  }

  // check if the words have the same letters
  for (const letter of letters) {
    if (!word1OccurrenceMap.get(letter) || !word2OccurrenceMap.get(letter)) {
      return false;
    }
  }

  // check if the occurrences have the same distribution
  const word1Occurrences = new Array(...word1OccurrenceMap.values()).sort(
    (a, b) => a - b
  );
  const word2Occurrences = new Array(...word2OccurrenceMap.values()).sort(
    (a, b) => a - b
  );

  for (let i = 0; i < word1Occurrences.length; i++) {
    if (word1Occurrences[i] !== word2Occurrences[i]) {
      return false;
    }
  }

  return true;
};

closeStrings("abbzzca", "babzzcz");
