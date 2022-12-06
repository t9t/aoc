import Data.Char
import Data.List

areAllDifferent s = (nub s) == s

brrrr :: String -> String -> Int -> Int
brrrr l r n
    | null r            = error "no solution found"
    | null l            = brrrr l' r' 4
    | areAllDifferent l = n
    | otherwise         = brrrr ((tail l) ++ [(head r)]) (tail r) n+1
    where (l', r') = splitAt 14 r

main = do
    c <- getContents
    print $ brrrr "" (head (lines c))  0
