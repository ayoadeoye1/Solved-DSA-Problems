


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
