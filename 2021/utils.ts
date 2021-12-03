const fs = require('fs');
const readline = require('readline');

export async function getReadingsAsArray<T>(filename: string, transform?: (line: string) => T): Promise<T[]> {
    const fileStream = fs.createReadStream(filename);

    const rl = readline.createInterface({
        input: fileStream,
        crlfDelay: Infinity
    });

    const readings = []

    for await (const line of rl) {
        readings.push(transform ? transform(line) : line);
    }

    return readings;
}