import Data.Char
import Data.List
import qualified Data.Map as Map

type Grid = [[Int]]

aboves :: Int -> Int -> Grid -> [Int]
aboves x y grid = reverse $ take y $ map (\line -> line !! x) grid

lefts :: Int -> Int -> Grid -> [Int]
lefts x y grid = reverse $ take x (grid !! y)

rights :: Int -> Int -> Grid -> [Int]
rights x y grid = drop (x+1) (grid !! y)

belows :: Int -> Int -> Grid -> [Int]
belows x y grid = drop (y+1) $ map (\line -> line !! x) grid

higherThanAll n numbers = and $ map (n>) numbers

isVisible :: Int -> Int -> Grid -> Bool
isVisible x y grid = or [higherThanAll n (f x y grid) | f <- funs]
    where n = (grid !! y) !! x
          funs = [aboves, lefts, rights, belows]

parse :: String -> Grid
parse s = map (map (\c -> read [c]::Int)) $ lines s

main = do
    c <- getContents
    let grid = parse c
    let rows = [0..(length grid)-1]
    let cols = [0..(length (head grid))-1]

    print $ length $ filter (True==) [ isVisible x y grid | x <- cols, y <- rows ]
