# gobowling
 A bowling game implementation in Go (interview question)

## Prompt

Write a class "Bowling" that allows one to solve rolls and scores.

---------------------------
--    RULES OF BOWLING   --
---------------------------

- 10 frames per game (Think of frames as turns)
- Every frame is scored independently
- Total score of the game is sum(frame scores)
- 10 Pins available to knock down in each frame
- Up to 2 rolls per frame (Think of rolls as attempts. You get two attempts per frame to knock down all pins)

---------------------------
--    SCORING A FRAME    --
---------------------------
1. If the 2 Rolls sum to less than 10, then the score of the Frame is the sum of the Rolls. Nothing special.
2. If you roll a Spare, the score of the Frame is 10 + the next Roll
3. If you roll a Strike, the score of the Frame is 10 + the next 2 Rolls
4. The tenth frame is different, there is the possibility of 3 rolls
    A. If you score a strike, you get two more rolls (3 rolls total)
    B. If you score a spare, you get one extra roll (3 rolls total)
    C. If you don't score a strike or spare, just a regular frame (2 rolls total)

---------------------------
--       EXAMPLE 1       --
---------------------------
Frame 1: Rolls => 3 & 4
Frame 1 has a score of 7 => 3 + 4

---------------------------
--       EXAMPLE 2       --
---------------------------
Frame 1: Rolls => 6 & 4
Frame 2: Rolls => 3 & 4

Frame 1 has a score of 13 => (6+4) + 3
Frame 2 has a score of 7 => 3 + 4

The total score so far is 20 => (13 + 7)

---------------------------
--       EXAMPLE 3       --
---------------------------
Frame 1: Rolls => 10
Frame 2: Rolls => 3 & 4

Frame 1 has a score of 17 => 10 + (3 + 4)
Frame 2 has a score of 7 => 3 + 4

The total score so far is 24 => 17 + 7

---------------------------
--       EXAMPLE 4       --
---------------------------
Frame 1: Rolls => 10
Frame 2: Rolls => 10
Frame 3: Rolls => 3 & 4

Frame 1 has a score of 23 => 10 + 10 + 3
Frame 2 has a score of 17 => 10 + (3 + 4)
Frame 3 has a score of 7 => 3 + 4

The total score so far is 47 => 23 + 17 + 7


---------------------------
-- Tenth Frame Example 1 --
---------------------------
Frame 10, Roll 1: 10  // This is a strike so two more rolls are allowed
Frame 10, Roll 2: 5
Frame 10, Roll 3: 5

Total score for 10th frame is 20 => 10 + 5 + 5


---------------------------
-- Tenth Frame Example 2 --
---------------------------
Frame 10, Roll 1: 5
Frame 10, Roll 2: 5  // This is a spare, one more roll are allowed
Frame 10, Roll 3: 10

Total score for 10th frame is 20 => 5 + 5 + 10


---------------------------
-- Tenth Frame Example 3 --
---------------------------
Frame 10, Roll 1: 2
Frame 10, Roll 2: 5  // NOT a strike or spare, game is over

Total score for 10th frame is 7 => 2 + 5
