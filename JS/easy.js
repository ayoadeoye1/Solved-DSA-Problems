
/**
 * 
 * @name Make-Array-Elements-Equal-To-Zero
 * 
 * @description Start by selecting a starting position curr such that nums[curr] == 0, and choose a movement direction of either left or right.

After that, you repeat the following process:

If curr is out of the range [0, n - 1], this process ends.
If nums[curr] == 0, move in the current direction by incrementing curr if you are moving right, or decrementing curr if you are moving left.
Else if nums[curr] > 0:
Decrement nums[curr] by 1.
Reverse your movement direction (left becomes right and vice versa).
Take a step in your new direction.
A selection of the initial position curr and movement direction is considered valid if every element in nums becomes 0 by the end of the process.

Note: Please do not copy the description during the contest to maintain the integrity of your submissions.©leetcode
 */
var countValidSelections = function(nums) {
    function canMakeAllZeros(start, direction) {
        let temp = [...nums];
        let n = temp.length;
        let curr = start;

        while (curr >= 0 && curr < n) {
          if (temp[curr] === 0) {
            curr += direction;
          } else {
            temp[curr] -= 1;
            direction *= -1;
            curr += direction;
          }
        }
        return temp.every(num => num === 0);
      }

      let validSelections = 0;
    
      for (let i = 0; i < nums.length; i++) {
        if (nums[i] === 0) {
          if (canMakeAllZeros(i, -1)) {
            validSelections++;
          }
          if (canMakeAllZeros(i, 1)) {
            validSelections++;
          }
        }
      }

  return validSelections;
};