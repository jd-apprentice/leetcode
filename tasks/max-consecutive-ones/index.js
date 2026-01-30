var findMaxConsecutiveOnes = function (nums) {
    return Math.max(
        ...nums.map(String)
            .join("")
            .split("0")
            .map(num => num.toString().length)
    );
};