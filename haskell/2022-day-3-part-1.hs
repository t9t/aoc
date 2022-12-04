import Data.Char
import Data.List

main = do
    c <- getContents
    print $ sum [scores . findFirstIntersection . splitAt ((length line) `div` 2) $ line | line <- lines c]
    where
        findFirstIntersection (a, b) = head . take 1 $ a `intersect` b
        scores c = (head . findIndices (c==) $ ['a'..'z']++['A'..'Z']) + 1
