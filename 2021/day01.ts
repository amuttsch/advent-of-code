import assert = require("assert");
import {getReadingsAsArray} from "./utils";

async function part1(readings: number[]) {
    let previousReading: number | undefined = undefined;
    let timesDepthIncreased = 0;

    readings.forEach(currentReading => {
        if (previousReading && currentReading > previousReading) {
            timesDepthIncreased++;
        }

        previousReading = currentReading;
    });

    return timesDepthIncreased;
}

async function part2(readings: number[]) {
    const measurementSums = [];

    for (let i=0; i<readings.length-2; i++) {
        measurementSums.push(readings[i] + readings[i+1] + readings[i+2])
    }

    return part1(measurementSums);
}

getReadingsAsArray('day01.txt', (line => parseInt(line, 10))).then(readings => {
    part1(readings).then(result => {
        console.log(`Part 1: ${result}`);
        assert(result === 1624, "Part 1: Solution is 1624")
    });

    part2(readings).then(result => {
        console.log(`Part 2: ${result}`);
        assert(result === 1653, "Part 2: Solution is 1653")
    });
});


