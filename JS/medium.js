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