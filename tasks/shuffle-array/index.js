/**
 * @param {number[]} nums
 * @param {number} n
 * @return {number[]}
 */
var shuffle = function (nums, n) {
    let newArr = []

    for (let i = 0; i < n; i++) {
        newArr.push(nums[i])
        newArr.push(nums[i + n])
    }

    return newArr
};

function main() {
    shuffle([2, 5, 1, 3, 4, 7], 3)
    shuffle([1, 2, 3, 4, 4, 3, 2, 1], 4)
    shuffle([1, 1, 2, 2], 2)
}

console.log(main())