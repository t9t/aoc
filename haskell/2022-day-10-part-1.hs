import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set


convert :: String -> [Int]
convert "noop" = [0]
convert line
    | take 4 line == "addx" = [0, n]
        where n = read (drop 5 line)::Int

convertLines :: [String] -> [Int]
convertLines lines = concat $ map convert lines

processAll :: Int -> Int -> [Int] -> [(Int, Int, Int)]
processAll _ _ [] = []
processAll c x operations =
    (c, x, x') : (processAll (c+1) x' (tail operations))
    where x' = x + (head operations)

kek [] = []
kek bla =
    top : kek rest
    where (first, rest) = splitAt 40 bla
          top = head first

ss (c, b, _) = c * b

main = do
    c <- getContents

    print $ sum $ map ss $ kek $ drop 19 $ processAll 1 1 $ convertLines $ lines c
