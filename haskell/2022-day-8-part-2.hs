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

arounds :: Int -> Int -> Grid -> [Int]
arounds x y grid =
    (aboves x y grid) ++
    (lefts x y grid) ++
    (rights x y grid) ++
    (belows x y grid)

higherThanAll n numbers = and $ map (n>) numbers

kek n l
    | null l = []
    | h < n  = h : (kek n r)
    | h > n  = h : []    
    | h == n = h : []
    where h = head l
          r = tail l

viewingDistance :: Int -> [Int] -> Int
viewingDistance n numbers = length $ kek n numbers

asdf :: Int -> Int -> Grid -> [Int]
asdf x y grid = [(viewingDistance n (aboves x y grid)), (viewingDistance n (lefts x y grid)), (viewingDistance n (rights x y grid)), (viewingDistance n (belows x y grid))]
    where n = (grid !! y) !! x

treeScore x y grid = product $ asdf x y grid

parse :: String -> Grid
parse s = map (map (\c -> read [c]::Int)) $ lines s

main = do
    c <- getContents
    let grid = parse c
    let rows = [0..(length grid)-1]
    let cols = [0..(length (head grid))-1]

    print $ maximum [ treeScore x y grid | x <- cols, y <- rows ]
