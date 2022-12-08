import Data.Char
import Data.List
import qualified Data.Map as Map

type Grid = [[Int]]

aboves :: Int -> Int -> Grid -> [Int]
aboves x y grid = take y $ map (\line -> line !! x) grid

lefts :: Int -> Int -> Grid -> [Int]
lefts x y grid = take x (grid !! y)

rights :: Int -> Int -> Grid -> [Int]
rights x y grid = drop (x+1) (grid !! y)

belows :: Int -> Int -> Grid -> [Int]
belows x y grid = drop (y+1) $ map (\line -> line !! x) grid

arounds :: Int -> Int -> Grid -> [Int]
arounds x y grid =
    (aboves x y grid) ++
    (lefts x y grid) ++
    (rights x y grid) ++
    (belows x y grid)

higherThanAll n numbers = and $ map (n>) numbers

isVisible :: Int -> Int -> Grid -> Bool
isVisible 0 _ _ = True
isVisible _ 0 _ = True
isVisible x y grid
    | x+1 == length (head grid) = True
    | y+1 == length grid        = True
    | otherwise                 = (higherThanAll n (aboves x y grid)) || (higherThanAll n (lefts x y grid)) || (higherThanAll n (rights x y grid)) || (higherThanAll n (belows x y grid))
    where n = (grid !! y) !! x

parse :: String -> Grid
parse s = map (map (\c -> read [c]::Int)) $ lines s

main = do
    c <- getContents
    let grid = parse c
    let rows = [0..(length grid)-1]
    let cols = [0..(length (head grid))-1]

    print $ length $ filter (True==) [ isVisible x y grid | x <- cols, y <- rows ]
