#!/usr/bin/env python

import argparse
import solver

def read_input(filename):
    content = []
    with open(filename) as f:
        content = f.readlines()
    content = [x.strip() for x in content]
    return content

def write_output(answers, filename):
    with open(filename, 'w') as f:
        for index, line in enumerate(answers):
            f.write("Case #{}: {}\n".format(index+1, line))

def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("-i", "--input", default="input.txt", help="Filename of the input")
    parser.add_argument("-o", "--output", default="output.txt", help="Filename of the output")

    args = parser.parse_args()

    problem_cases = read_input(args.input)

    answers = solver.solve(problem_cases)

    write_output(answers, args.output)

if __name__ == "__main__":
    main()
