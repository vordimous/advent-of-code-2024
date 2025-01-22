Interview

Candidate:
    AJ Danelz

Interviewers:
    Jeff McBride
    Zach Heath

Agenda:
* Introductions
* Code Exercise
* Search Experience Technical Questions
* Your Questions (~15 min)


The Bowling Game

Task: Create code that can calculate the score for a complete bowling game for 1 player


We encourage you to ask questions as we go!

Bowling rules

A Bowling Game is comprised of 10 Frames

In each Frame
- Previous Frame's pins are cleared
- 10 new pins are setup
- You get up to 2 rolls to knock down as many pins as you can.
- If you knock down all 10 pins on your first roll it is considered a Strike and that frame is done (don't use a second roll)
- If you knock down all 10 pins using both rolls it is considered a Spare (ex 8 + 2)
- If you knock down less that 10 pins with both rolls it is considered an Open Frame ( ex 8 + 1)
- The last (10th) Frame is special - We'll get to that

Scoring:

If a Frame is an open frame you score the number of pins knocked down during that frame

If a Frame is a Spare you score 10 points plus the number of pins knocked down on your next roll (from the next frame)

If a Frame is a Strike you score 10 points plus the number of pins knocked down on your next 2 rolls (from the next frame(s))


The last frame is special because there are no following Frames.
However, if you get a Spare or a Strike you still get your bonus rolls, they are just part of the 10th frame.
1 Bonus roll for a Spare and 2 Bonus rolls for a Strike


Examples:
[0,6]
Frame 1 Score 6 - Running Game Score 6

[0,6] [4,6] [2,5]
Frame 1 Score 6
Frame 2 Score 12
Frame 3 Score 7 - Running Game Score 25

[0,6] [4,6] [2,5] [10,] [10,] [2,5]
Frame 1 Score 6
Frame 2 Score 12
Frame 3 Score 7
Frame 4 Score 22
Frame 5 Score 17
Frame 6 Score 7 - Running Game Score 71



Code Exercise:

Task: Create code that can calculate the score for a complete bowling game for 1 player

There should be a class with the following interface

void rollBall(int pins);

This is called each time the player rolls a ball and does not return anything.
It takes in the number of pins 0-10 knocked down for that roll.

int getScore();

This gets called after the game is complete and should return the players total score.

For testing purposes, you can return whatever you want when the game is not yet complete.

You can have as many additional variables, data structures, additional classes as needed.

We are NOT looking for highly optimized, high performant, tricky code.

We are looking for clean code that works and is easy to understand and is well tested.

Preferably we work through the problem together in an incremental TDD manner.

Relax: We get it...  interviews are stressful... we both went through this

Messy, incomplete, wrong to begin with is fine / how we refactor/improve as we go is what we care about

Ask questions / ask for help / think out loud!
