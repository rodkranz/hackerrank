# Super Reduced String

Steve has a string, *s*, consisting of *n* lowercase English alphabetic letters. In one operation, he can delete any pair of adjacent letters with same value. For example, string `aabcc` would become either `aab` or `bcc` after **1** operation.

Steve wants to reduce  as much as possible. To do this, he will repeat the above operation as many times as it can be performed. Help Steve out by finding and printing 's non-reducible form!

**Note**: If the final string is empty, print `Empty String`.

### Input Format

A single string, *s*.

### Constraints

 * 1 <= n <= 100

### Output Format

If the final string is empty, print `Empty String`; otherwise, print the final non-reducible string.

### Sample Input 0

    aaabccddd
    
### Sample Output 0

    abd
    
### Explanation 0

Steve can perform the following sequence of operations to get the final string:

   1) `aaabccddd` → `abccddd`
   2) `abccddd` → `abddd`
   3) `abddd` → `abd`
   
Thus, we print `abd`.