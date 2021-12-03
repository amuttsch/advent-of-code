import {getReadingsAsArray} from "./utils";
import assert = require("assert");

const getMostCommonBit = (lines: string[], position: number) => {
    const bits = {"0": 0, "1": 0};

    lines.forEach(line => {
        const bit = line[position];
        // @ts-ignore
        bits[bit]++;
    });

    if (bits["0"] > bits["1"]) {
        return "0"
    } else {
        return "1"
    }
}

const part1 = (lines: string[]) => {
    const mostCommonBits = [];

    for (let i=0;i<lines[0].length; i++) {
        mostCommonBits.push(getMostCommonBit(lines, i));
    }

    const gamma = parseInt(mostCommonBits.join(""), 2);
    const epsilon = parseInt(mostCommonBits.map(bit => {
        switch (bit) {
            case "0": return "1";
            case "1": return "0";
        }
    }).join(""), 2);

    return gamma * epsilon;
}

const part2 = (lines: string[]) => {
    let oxygenLines = [...lines];
    let co2Lines = [...lines];

    for (let i=0;i<lines[0].length; i++) {
        const mostCommonBitOxygen = getMostCommonBit(oxygenLines, i);
        const mostCommonBitCo2 = getMostCommonBit(co2Lines, i);

        if (oxygenLines.length !== 1) {
            oxygenLines = oxygenLines.filter(line => line[i] === mostCommonBitOxygen);
        }

        if (co2Lines.length !== 1) {
            co2Lines = co2Lines.filter(line => line[i] !== mostCommonBitCo2);
        }
    }

    return parseInt(oxygenLines[0], 2) * parseInt(co2Lines[0], 2);
}

getReadingsAsArray<string>('day03.txt').then(readings => {
    const result1 = part1(readings);
    console.log(`Part 1: ${result1}`);
    assert(result1 === 2648450, "Part 1: Solution is 2648450");

    const result2 = part2(readings);
    console.log(`Part 2: ${result2}`);
    assert(result2 === 2845944, "Part 2: Solution is 2845944")
});
