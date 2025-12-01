import fs from "fs";

function main() {
  const filePath: string = "input.txt";
  try {
    const fileContent: string = fs.readFileSync(filePath, "utf8");

    const records: string[] = fileContent.split(/\r?\n/);

    let p1: number = partOne(records);
    let p2: number = partTwo(records);

    console.log("Password for Part 1: ", p1);
    console.log("Password for Part 2: ", p2);
  } catch (err) {
    console.error("Error reading file: ", err);
  }
}

// Part One =====================================================

function partOne(records: string[]): number {
  let res: number = 0;
  let cur: number = 50;

  for (let i in records) {
    let line: string = records[i];
    let val: number = parseInt(line.slice(1));

    val %= 100;
    if (line[0] == "L") {
      val = -val;
    }

    cur += val;

    if (cur > 99) cur -= 100;
    else if (cur < 0) cur += 100;

    if (cur === 0) res++;
  }

  return res;
}

// Part Two =====================================================

function partTwo(records: string[]): number {
  let res: number = 0;
  let cur: number = 50;
  let prev: number = 0;

  for (let i in records) {
    let line: string = records[i];
    let numString: string = line.slice(1);
    let val: number = parseInt(numString, 10);

    res += Math.floor(val / 100);

    val %= 100;
    if (line[0] === "L") {
      val = -val;
    }

    cur += val;

    if (cur > 99) cur -= 100;
    else if (cur < 0) cur += 100;

    if (cur === 0) {
      res++;
    } else if (cur >= prev && line[0] == "L" && prev != 0) {
      res++;
    } else if (cur <= prev && line[0] == "R" && prev != 0) {
      res++;
    }

    prev = cur;
  }

  return res;
}

main();
