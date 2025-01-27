/**
 * 
 * @name Can Finish
 * @alias Graph
 * @description 
 * 
 * 
 */
var canFinish = function(n, prerequisites) {
    let graph = Array.from({length: n}, () => [])
    let inDegree = Array(n).fill(0)
    for(let [c, p] of prerequisites){
        graph[p].push(c);
        inDegree[c]++
    }

    let stack = [];
    for(let i=0; i<n; i++){
        if(inDegree[i]===0){
            stack.push(i)
        }
    }

    let completed = 0;
    while(stack.length > 0){
        const course = stack.shift()
        completed++

        for(let neig of graph[course]){
            inDegree[neig]--
            if(inDegree[neig]===0){
                stack.push(neig)
            }
        }
    }
    return n===completed
};


/**
* 
* @name Binary-Tree-To-Linked-List
* @alias BST
* @description 
* 
* 
*/
var flatten = function(root) {
    if(!root){
        return 
    }
    let right = root.right;
    let left = root.left;
    flatten(right);
    flatten(left);
    let subRight = root.right;
    root.right = root.left;
    root.left=null;
    let current = root;
    while(current.right){
        current = current.right;
    }
    current.right = subRight;
};


//BruteForce
var bulbSwitch = function(n) {
    let arr = new Array(n).fill(0)
    for(let i=0; i<n; i++){
        const num = i+1;
        for(let j=1; j<=n; j++){
            if(j%num === 0){
                arr[j] ^=1;
            }
        }
    }
    let count = 0;
    console.log(arr);
    for(let i=0; i<arr.length; i++){
        count += arr[i];
    }
    console.log(count);
    return count;
};

//Bulbs that are tggled in number of odd times remains on
//for ith of n, all its factors are toggled
//Tags: Factors Even Odd
var bulbSwitch = function(n) {
    return Math.floor(Math.sqrt(n));
};