import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

type Pos = (Int, Int)
type Visited = Set.Set Pos
type State = (Pos, Pos, Visited)

moveTailTowardsHead :: Pos -> Pos -> (Pos, Bool)
moveTailTowardsHead head tail
    | dMax <= 1 = (tail, False)
    | dx == 0   = ((tx, ty + sdy), True)
    | dy == 0   = ((tx + sdx, ty), True)
    | otherwise = ((tx + sdx, ty + sdy), True)
    where (hx, hy) = head
          (tx, ty) = tail
          dx       = hx - tx
          dy       = hy - ty
          dx'      = abs dx
          dy'      = abs dy
          sdx      = dx `div` dx'
          sdy      = dy `div` dy'
          dMax     = max dx' dy'

moveTailFullyToHead :: Pos -> (Pos, Visited) -> (Pos, Visited)
moveTailFullyToHead head (tail, visited) = if moved
    then moveTailFullyToHead head (tail', Set.insert tail' visited)
    else (tail, visited)
    where (tail', moved) = moveTailTowardsHead head tail

move :: Char -> Pos -> Pos
move 'R' (x, y) = (x+1, y)
move 'U' (x, y) = (x, y-1)
move 'L' (x, y) = (x-1, y)
move 'D' (x, y) = (x, y+1)

moveHeadAndTail :: Char -> State -> State
moveHeadAndTail dir (head, tail, visited) =
    (head', tail', visited')
    where head'             = (move dir head)
          (tail', visited') = moveTailFullyToHead head' (tail, visited)

moveAll :: [Char] -> State -> State
moveAll moves st = foldl (\st' dir -> moveHeadAndTail dir st') st moves

moveAllAndCountVisited :: [Char] -> Int
moveAllAndCountVisited moves = Set.size visited
    where (_, _, visited) = moveAll moves (s, s, Set.singleton s)
          s = (0, 0)

parse :: String -> [Char]
parse s = concat [take n (repeat dir) | line <- lines s, let dir = head line, let n = read (drop 2 line)::Int]

main = do
    c <- getContents

    print $ moveAllAndCountVisited $ parse c
