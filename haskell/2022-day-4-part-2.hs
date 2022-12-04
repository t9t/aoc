import Data.Char
import Data.List

breakOn c s = (l, tail r) where (l, r) = break (c==) s
mapBoth fn (l, r) = (fn l, fn r)
eitherContainsOther (l, r) = not . null $ l `intersect` r
toRange (l, r) = [l..r]
parse line = mapBoth (\p -> toRange(mapBoth (\s -> (read s :: Integer)) (breakOn '-' p))) (breakOn ',' line)

main = do
    c <- getContents
    print $ length $ filter eitherContainsOther $ map parse $ lines c
