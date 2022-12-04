import Data.Char
import Data.List

main = do
    c <- getContents
    print $ sum [scores . findFirstIntersection . splitAt ((length line) `div` 2) $ line | line <- lines c]
    where
        findFirstIntersection (a, b) = take 1 (a `intersect` b) !! 0
        scores c = (findIndices (c==) (['a'..'z']++['A'..'Z']) !! 0) + 1
