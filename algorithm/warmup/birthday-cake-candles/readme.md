# Birthday Cake Candles

You are in-charge of the cake for your niece's birthday and have decided the cake will have one candle for each year of her total age. When she blows out the candles, she’ll only be able to blow out the tallest ones. Your task is to find out how many candles she can successfully blow out.

For example, if your niece is turning **4** years old, and the cake will have **4** candles of height **3**, **2**, **1**, **3**, she will be able to blow out **2** candles successfully, since the tallest candle is of height **3** and there are **2** such candles.

Complete the function `birthdayCakeCandles` that takes your niece's age and an integer array containing height of each candle as input, and return the number of candles she can successfully blow out.

### Input Format

The first line contains a single integer, , denoting the number of candles on the cake. 
The second line contains *n* space-separated integers, where each integer *i* describes the height of candle .

### Constraints

 * 1 <= n <= 10^5
 * 1 <= *height*i <= 10^7


### Output Format

Print the number of candles the can be blown out on a new line.

### Sample Input 0

	4
	3 2 1 3

### Sample Output 0

	2

### Explanation 0

We have one candle of height **1**, one candle of height **2**, and two candles of height **3**. Your niece only blows out the tallest candles, meaning the candles where **height = 3**. Because there are **2** such candles, we print **2** on a new line.