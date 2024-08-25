# 401. Binary Watch
- Easy, Topics, Companies, Hint

A binary watch has 4 LEDs on the top to represent the hours (0-11), and 6 LEDs on the bottom to represent the minutes (0-59).
Each LED represents a zero or one, with the least significant bit on the right.

    For example, the below binary watch reads "4:51".

Given an integer turnedOn which represents the number of LEDs that are currently on (ignoring the PM), return all possible times the watch could represent.
You may return the answer in any order.

The hour must not contain a leading zero.

    For example, "01:00" is not valid. It should be "1:00".

The minute must consist of two digits and may contain a leading zero.

    For example, "10:2" is not valid. It should be "10:02".


## Example 1:

```
Input: turnedOn = 1
Output: ["0:01","0:02","0:04","0:08","0:16","0:32","1:00","2:00","4:00","8:00"]
```

## Example 2:

```
Input: turnedOn = 9
Output: []
```

### Constraints:

- 0 <= turnedOn <= 10


## Solution
- Data struct for the binaryWatch: just two sets one for the hours and another for the minutes 

- Enumerate all possible hours and minutes combination giver the amount of leds turned on
- May use backtrackint to traverse all de permutations

###  Backtracking
1. Choice: At any giver hour that I'm in I can chose beetween all the remaning minutes that were not already chosen in this hour
2. Cronstaint: If there is no more minutes to put in this hour than enumate all the possibilites in the next hour
3. Goal: Giver a valid time (h:mm) store that as a possible result

3

Start with hour 0

go to first minute
