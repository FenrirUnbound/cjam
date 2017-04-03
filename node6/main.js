#!/usr/bin/env node

'use strict';

const fs = require('fs');
const solver = require('./solver');

/**
 * Parses argv for the input and output file names.
 * @param  {String[]} argv    process.argv
 * @return {Object}           Object containing keys:
 *                                "input"  - input filename
 *                                "output" - output filename
 */
function parseArgs(argv) {
  const flags = {
    '-i': 'input',
    '--input': 'input',
    '-o': 'output',
    '--output': 'output'
  };
  const result = {
    input: '',
    output: ''
  };
  let saveNext = '';

  argv.slice(2).forEach((item) => {
    if (saveNext) {
      result[saveNext] = item;
      saveNext = '';
    } else if (flags[item]) {
      saveNext = flags[item];
    }
  });

  // default values
  result.input = result.input || 'input.txt';
  result.output = result.output || 'output.txt';

  return result;
}

/**
 * Read the input problem cases from a local file.
 * As a side effect, it will prune the last line in a file, if one is
 * present
 * @param  {String} filename  local filename of the problem cases
 * @return {Promise}          Promise that resolves into a string
 *                                array of the problem cases
 */
function readInput(filename) {
  return new Promise((resolve, reject) => {
    fs.readFile(filename, 'utf8', (err, data) => {
      if (err) {
        return reject(err);
      }

      const result = data.split('\n');
      const length = parseInt(data[0], 10);
      const checkLast = result[result.length - 1];

      return resolve((checkLast) ? result : result.slice(0, length+1));
    });
  });
}

/**
 * Write the answers to the designated output file, according to the
 * correct answer format
 * @param  {String}   filename    filename to store the answers to
 * @param  {String[]} answers     Answers to each problem each
 * @return {Promise}              Promise that resolves when data is
 *                                    written to the designated file
 */
function writeOutput(filename, answers) {
  return new Promise((resolve, reject) => {
    const results = answers.map((answer, index) => {
      return `Case #${index+1}: ${answer}`;
    });

    fs.writeFile(filename, results.join('\n'), 'utf8', (err) => {
      if (err) {
        return reject(err);
      }

      return resolve();
    });
  });
}

/**
 * Calls the solver function to process the given problem set.
 * Has additional complexity for simple Promise-based compatibility. The
 * checks are simplistic in nature, and cover only the desired outputs
 * of the Solver.
 * @param  {String[]} problemCases    The problem cases read in from the file
 * @return {Promise}                  Promise that resolves to the answers
 *                                       given from the solver
 */
function callSolver(problemCases) {
  return new Promise((resolve, reject) => {
    const maybePromise = solver(problemCases, (err, data) => {
      if (err) {
        return reject(err);
      }

      return resolve(data);
    });

    // compatability with callback-style users
    if (maybePromise) {
      if (Array.isArray(maybePromise) || typeof maybePromise.then === 'function') {
        return resolve(maybePromise);
      }
    }
  });
}

/**
 * Main function. Is the function called when the script is ran
 * As a side effect, will exit with code 0 upon success, or 1 upon
 * failure.
 * @return {Promise}    Promise that resolves upon completion.
 */
function main() {
  const filenames = parseArgs(process.argv);

  return readInput(filenames.input)
  .then(callSolver)
  .then(answers => writeOutput(filenames.output, answers))
  .then(() => process.exit(0))
  .catch((err) => {
    console.error('Error encountered');
    console.error(err);
    process.exit(1);
  });
}

return main();
