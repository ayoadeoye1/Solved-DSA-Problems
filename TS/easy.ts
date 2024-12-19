


/**
 * 
 * @name 
 * 
 * @description 
 * 
 * 
 */
function canAliceWin(n: number): boolean {
    let arem: number = 10;
    let brem: number = 9;
    let remain: number = n;
    let isAliceTurn: boolean = true;
    while (remain > 0) {
        if (isAliceTurn) {
            remain -= arem;
            brem = arem - 1;
            isAliceTurn = false
            if(remain < 0) return false
            if(remain === 0) return true
        } else {
            remain -= brem;
            arem = brem - 1
            isAliceTurn = true
            if(remain < 0) return true
            if(remain === 0) return false
        }
    }
};

/**
 * 
 * @name permutation
 * 
 * @description 
 * 
 * 
 */
var permute = function(nums) {
    const res = [];
    function backTrack(start){
        if(start === nums.length){
            res.push([...nums]);
            return;
        }
        for(let i=start; i<nums.length; i++){
            [nums[start], nums[i]] = [nums[i], nums[start]];
            backTrack(start+1);
            [nums[start], nums[i]] = [nums[i], nums[start]];
        }
    }
    backTrack(0);
    return res;
};


/**
 * 
 * @name Diameter-Of-Binary-Tree
 * @alias BST DFS
 * @description 
 * 
 * 
 */
var diameterOfBinaryTree = function(root) {
    let maxLen = 0;
    const dfs = (root) => {
        if(!root){
            return 0;
        }
        let right = dfs(root.right);
        let left = dfs(root.left);
        maxLen = Math.max(maxLen, right+left);
        return 1+Math.max(right, left);
    }
    dfs(root);
    return maxLen;
};

