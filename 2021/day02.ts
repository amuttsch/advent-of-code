import {getReadingsAsArray} from "./utils";
import assert = require("assert");

type Direction = 'up' | 'down' | 'forward';

type Command = {
    readonly direction: Direction;
    readonly units: number;
}

const CommandRegex = new RegExp("([a-z]*) (\\d*)")

const transform = (line: string): Command => {
    const result = line.match(CommandRegex);
    if (!result) {
        throw new Error(`Line "${line}" invalid"`);
    }

    const [_, direction, units] = result;

    return {
        direction: direction as Direction,
        units: parseInt(units, 10)
    }
}

const part1 = (commands: Command[]) => {
    let depth = 0;
    let horizontalPosition = 0;

    commands.forEach(command => {
        switch (command.direction) {
            case "up":
                depth -= command.units;
                break;
            case "down":
                depth += command.units;
                break;
            case "forward":
                horizontalPosition += command.units;
                break;
        }
    });

    return depth * horizontalPosition;
}

const part2 = (commands: Command[]) => {
    let depth = 0;
    let horizontalPosition = 0;
    let aim = 0;

    commands.forEach(command => {
        switch (command.direction) {
            case "up":
                aim -= command.units;
                break;
            case "down":
                aim += command.units;
                break;
            case "forward":
                horizontalPosition += command.units;
                depth += command.units * aim;
                break;
        }
    });

    return depth * horizontalPosition;
}

getReadingsAsArray<Command>('day02.txt', transform).then(readings => {
    const result1 = part1(readings);
    console.log(`Part 1: ${result1}`);
    assert(result1 === 2091984, "Part 1: Solution is 2091984");

    const result2 = part2(readings);
    console.log(`Part 2: ${result2}`);
    assert(result2 === 2086261056, "Part 2: Solution is 2086261056")
});
