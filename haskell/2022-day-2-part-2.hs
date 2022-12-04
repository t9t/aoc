import Data.Char
import Data.List

baseScoreRock = 1
baseScorePaper = 2
baseScoreScissors = 3

scoreLose = 0
scoreDraw = 3
scoreWin = 6

points ("A", "X") = baseScoreScissors + scoreLose
points ("A", "Y") = baseScoreRock + scoreDraw
points ("A", "Z") = baseScorePaper + scoreWin

points ("B", "X") = baseScoreRock + scoreLose
points ("B", "Y") = baseScorePaper + scoreDraw
points ("B", "Z") = baseScoreScissors + scoreWin

points ("C", "X") = baseScorePaper + scoreLose
points ("C", "Y") = baseScoreScissors + scoreDraw
points ("C", "Z") = baseScoreRock + scoreWin

main = do
    c <- getContents
    print $ sum [points . splitAt 1 . delete ' ' $ line | line <- lines c]
