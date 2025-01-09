


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
 * @name Has Cycle
 * 
 * @description Linked List Cycle
 * 
 * 
 */
class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
function hasCycle(head: ListNode | null): boolean {
    if(!head || !head.next){
        return false;
    }
    let slow: ListNode = head;
    let fast: ListNode = head.next;
    while(slow !== fast){
        if(!fast || !fast.next){
            return false
        }
        slow = slow.next;
        fast = fast.next.next;
    }
    return true;
};
