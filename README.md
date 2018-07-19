Happy Numbers Tester

Happy Numbers are numbers that when placed through algorithm eventually result in the value of 1.
All other numbers are unhappy and will eventually converge to the value of 4.

Happy Numbers Algorithm:
  1) Split the numbers digits apart
  2) Square the value of each individual digit.
  3) Sum the squares.
  4) If the sum of the squares is 1 or 4, you are complete.
  5) If the sum of the squares is any other value, return to step 1 with the new value.

Project Components:

Happy Numbers library
 - src\github.com\kbtibbs\happy_numbers\happynumbers.go
Happy Numbers console app utilizing the Terminal Colors library
 - src\github.com\kbtibbs\kount
Happy Numbers Test File - creates or appends results to access.log
 - src\github.com\kbtibbs\happy_numbers\happynumbers_test.go

 Known issues:
  - Some windows applications (Notepad) do not respect the linefeed used by Go test. More flexible applications like wordpad or Sublime Text 2 handle them correctly.