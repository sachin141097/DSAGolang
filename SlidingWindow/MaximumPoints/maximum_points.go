/*
There are several cards arranged in a row, and each card has an associated number of points. The points are given in the integer array cardPoints.

In one step, you can take one card from the beginning or from the end of the row. You have to take exactly k cards.

Your score is the sum of the points of the cards you have taken.

Given the integer array cardPoints and the integer k, return the maximum score you can obtain.

Example 1:

Input: cardPoints = [1,2,3,4,5,6,1], k = 3
Output: 12
Explanation: After the first step, your score will always be 1. However, choosing the rightmost card first will maximize your total score. The optimal strategy is to take the three cards on the right, giving a final score of 1 + 6 + 5 = 12.
Example 2:

Input: cardPoints = [2,2,2], k = 2
Output: 4
Explanation: Regardless of which two cards you take, your score will always be 4.
*/

/*
For picking up k=3 cards these are the possible ways:
[1,2,3,4,5,6,1]
1. Pick the first 3 cards (all from left) and zero from right hand side--> 1+2+3=6
2. Pick the first 2 cards (all from left) and 1 from right hand side --> 1+2+1=4
3. Pick the first 1 card (all from left) and 2 from right hand side --> 1+6+1=8
4. Pick the first 0 card (all from left) and 3 from right hand side --> 1+5+6=12

So, the maximum points we can get is 12.
*/
//Time Complexity: O(k)
//Space Complexity: O(1)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	leftSum := 0
	maxSum := 0
	for i := 0; i < k; i++ {
		leftSum += cardPoints[i]
		maxSum = max(maxSum, leftSum)
	}
	rightSum := 0
	rightIndex := n - 1
	for i := k - 1; i >= 0; i-- {
		leftSum -= cardPoints[i]
		rightSum += cardPoints[rightIndex]
		maxSum = max(maxSum, leftSum+rightSum)
		rightIndex--
	}
	return maxSum
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter card points space separated:")
	scanner.Scan()
	cardInput := strings.Fields(scanner.Text())
	var cardPoints []int
	for _, numStr := range cardInput {
		num, _ := strconv.Atoi(numStr)
		cardPoints = append(cardPoints, num)
	}
	fmt.Println("Enter k:")
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	fmt.Println("Maximum points:", maxScore(cardPoints, k))

}
