import Data.Char
import Data.List

chunkAt n [] = []
chunkAt n l = c : chunkAt n l'
    where (c, l') = splitAt n l


intersectsAll :: String -> [String] -> String
intersectsAll l ll
    | length ll == 1 = nub(l `intersect` (head ll))
    | otherwise      = nub(l `intersect` ((head ll) `intersectsAll` (tail ll)))

intersectAll :: [String] -> String
intersectAll ll
    | length ll == 1 = head ll
    | otherwise      = (head ll) `intersectsAll` (tail ll)

main = do
    c <- getContents
    print $ sum [scores . head . intersectAll $ chunk | chunk <- chunkAt 3 (lines c)]
    where
        scores c = (head . findIndices (c==) $ ['a'..'z']++['A'..'Z']) + 1
