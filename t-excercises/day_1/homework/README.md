
***Task 1***

Check if a given number or part of this number is a palindrome. For example, the number 1234437 is not a palindrome,
but its part 3443 is a palindrome. Numbers less than 10 are counted as invalid input.
```
Easy level:
Find if a number is a palindrome or not

Mid level:
You need to find only first subpalindrome

Hard:
Find all subpalindromes
```

Input parameters: number

Output: the palindrome/s extracted from the number, or 0 if the extraction failed or no palindromes was found.


***Task 2***

There are 2 ways to count lucky tickets:
1. Simple - if a six-digit number is printed on the ticket, and the sum of the first three digits is equal to the sum of the last three digits, then this
   ticket is lucky.
2. Hard - if the sum of the even digits of the ticket is equal to the sum of the odd digits of the ticket, then the ticket is considered lucky.
   Determine programmatically which variant of counting lucky tickets will give them a greater number at a given interval.
   
Task: Calculate how many tickets are lucky within provided min and max tickets

Input parameters: 2 values min digit of ticket and max digit of ticket
Output: information about the winning method and the number of lucky tickets for each counting method.

Input numbers contains exactly 6 digits. Not more or less


Example:
```
Min: 120123
Max: 320320

--Result--
EasyFormula: 11187
HardFormula: 5790
```
