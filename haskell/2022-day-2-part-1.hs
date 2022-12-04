import Data.Char
import Data.List

rock = "A"
paper = "B"
scissors = "C"

rock' = "X"
paper' = "Y"
scissors' = "Z"

baseScoreRock = 1
baseScorePaper = 2
baseScoreScissors = 3

scoreLose = 0
scoreDraw = 3
scoreWin = 6


points ("A", "X") = baseScoreRock + scoreDraw
points ("A", "Y") = baseScorePaper + scoreWin
points ("A", "Z") = baseScoreScissors + scoreLose

points ("B", "X") = baseScoreRock + scoreLose
points ("B", "Y") = baseScorePaper + scoreDraw
points ("B", "Z") = baseScoreScissors + scoreWin

points ("C", "X") = baseScoreRock + scoreWin
points ("C", "Y") = baseScorePaper + scoreLose
points ("C", "Z") = baseScoreScissors + scoreDraw

main = do
    c <- getContents
    print $ sum [points . splitAt 1 . delete ' ' $ line | line <- lines c]
