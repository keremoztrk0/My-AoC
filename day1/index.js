import fs from 'fs'

const main = async () => {
    const input = fs.readFileSync("input_day1.txt").toString().split("\n");

    const start1 = performance.now();
    const part1_result = day1_part1(input);
    const end1 = performance.now();
    console.log(`${part1_result} / Execution time: ${end1 - start1} ms`);

    const start2 = performance.now()
    const part2_result = day1_part2(input);
    const end2 = performance.now()
    console.log(`${part2_result} / Execution time: ${end2 - start2} ms`);


}

const matchedNumbers = [
    { text: "one", number: 1 },
    { text: "two", number: 2 },
    { text: "three", number: 3 },
    { text: "four", number: 4 },
    { text: "five", number: 5 },
    { text: "six", number: 6 },
    { text: "seven", number: 7 },
    { text: "eight", number: 8 },
    { text: "nine", number: 9 }
];
const day1_part1 = (input) => {
    let store = 0;
    for (let line of input) {
        let leftNum = -1;
        let rightNum = -1;
        for (let i = 0; i < line.length; i++) {
            const num = parseInt(line[i]);
            if (!isNaN(num)) {
                if (leftNum == -1) {
                    leftNum = num;
                    rightNum = num;
                }
                else {
                    rightNum = num;
                }
            }
        }
        store += parseInt(leftNum + '' + rightNum)
    }
    return store;
}
const day1_part2 = (input) => {
    let store = 0;
    for (let line of input) {
        let leftNum = -1;
        let rightNum = -1;
        for (let i = 0; i < line.length; i++) {
            let num = parseInt(line[i]);
            let isNum = !isNaN(num);
            let isStringNum = false;
            if (!isNum) {
                let matchedNum = numFromSpelledNum(line, i);
                isStringNum = matchedNum != -1
                if (isStringNum) {
                    num = matchedNum.number;
                }
            }
            if (isStringNum || isNum) {

                if (leftNum == -1) {
                    leftNum = num;
                    rightNum = num;
                }
                else {
                    rightNum = num;
                }
            }

        }
        store += parseInt(leftNum + '' + rightNum)

    }
    return store;
    function numFromSpelledNum(line, currentIndex) {
        line = line.slice(currentIndex);
        var result = -1;
        for (const matchedNumber of matchedNumbers) {
            if (line.startsWith(matchedNumber.text)) {
                result = matchedNumber
            }
        }
        return result

    }
}
await main();
