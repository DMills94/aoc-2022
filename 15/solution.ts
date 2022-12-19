import * as fs from "fs";
import { range } from "lodash";

type Coords = [number, number];

const dist = (start: Coords, end: Coords): number => {
  return Math.abs(start[0] - end[0]) + Math.abs(start[1] - end[1]);
};

// const input = fs.readFileSync("./example.txt", "utf-8");
const input = fs.readFileSync("./input.txt", "utf-8");

const signals: Set<[...Coords, number]> = new Set();
const beacons: Set<string> = new Set();
for (let pair of input.split("\n")) {
  const [s, b] = pair.split(":");
  const smatch = s.match("x=(-?[0-9]+), y=(-?[0-9]+)");
  const bmatch = b.match("x=(-?[0-9]+), y=(-?[0-9]+)");
  if (smatch && bmatch) {
    const [, sx, sy] = smatch;
    const [, bx, by] = bmatch;
    const d = dist([+sx, +sy], [+bx, +by]);
    signals.add([+sx, +sy, d]);
    beacons.add(`${bx},${by}`);
  }
}

const isValid = (
  x: number,
  y: number,
  S: Set<[...Coords, number]>
): boolean => {
  for (let [sx, sy, d] of S) {
    const dxy = dist([x, y], [sx, sy]);
    if (dxy <= d) return false;
  }
  return true;
};

let validSpots = 0;
const y = 2e6;

for (let x = -1e7; x < 1e7; x++) {
  if (!isValid(x, y, signals) && !beacons.has(`${x},${y}`)) {
    validSpots++;
  }
}

console.log(`Valid spots: ${validSpots}`);

const max = 4e6;
const calcFreq = ([x, y]: number[]): number => {
  return x * max + y;
};

let found = false;
let beacon = [0, 0];
while (!found) {
  for (let [sx, sy, d] of signals) {
    for (let dx of range(d + 2)) {
      const dy = d + 1 - dx;
      const x = sx + dx * -1;
      const y = sy + dy * -1;
      if (x < 0 || x > max || y < 0 || y > max) {
        continue;
      }
      if (dist([x, y], [sx, sy]) === d + 1 && isValid(x, y, signals)) {
        beacon = [x, y];
        found = true;
      }
    }
  }
}
console.log(`Tuning frequency: ${calcFreq(beacon)}`);
