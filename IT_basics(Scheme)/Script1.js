function getRandomInt(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

function getRandomStrings(count, length) {
    let left = 33,
        right = 127;

    let strings = Array(count).fill('')

    strings.forEach((_, index) => {
        for (let _ = 0; _ < length; _++)
            strings[index] += String.fromCharCode(getRandomInt(left, right))
    })
    return strings
}

module.exports = getRandomStrings