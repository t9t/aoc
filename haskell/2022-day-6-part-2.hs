import Data.Char
import Data.List

areAllDifferent s = (nub s) == s

brrrr :: String -> String -> Int -> Int
brrrr l r n
    | null r            = error "no solution found"
    | null l            = brrrr (take 14 r) (drop 14 r) 14
    | areAllDifferent l = n
    | otherwise         = brrrr ((tail l) ++ [(head r)]) (tail r) n+1

main = do
    c <- getContents
    print $ brrrr "" (head (lines c))  0
