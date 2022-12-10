import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

type Pos = (Int, Int)
type Visited = Set.Set Pos
type Rope = [Pos]
type State = (Rope, Visited)

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

moveTailFullyToHead :: Pos -> Pos -> Pos
moveTailFullyToHead head tail = if moved
    then moveTailFullyToHead head tail'
    else tail
    where (tail', moved) = moveTailTowardsHead head tail

move :: Char -> Pos -> Pos
move 'R' (x, y) = (x+1, y)
move 'U' (x, y) = (x, y-1)
move 'L' (x, y) = (x-1, y)
move 'D' (x, y) = (x, y+1)

moveAllFullyTowardTheHead :: Rope -> Rope -> Visited -> ([Pos], Visited)
moveAllFullyTowardTheHead soFar theRest visited
    | null theRest = (soFar, visited)
    | otherwise    = moveAllFullyTowardTheHead (soFar ++ [moved]) theRest' visited'
        where target   = last soFar
              toMove   = head theRest
              theRest' = tail theRest
              moved    = moveTailFullyToHead target toMove
              isTail   = length theRest == 1
              visited' = if isTail then Set.insert moved visited else visited

moveHeadAndTail :: Char -> State -> State
moveHeadAndTail dir (rope, visited) = moveAllFullyTowardTheHead [h'] (tail rope) visited
    where h  = head rope
          h' = move dir h

moveAll :: [Char] -> State -> State
moveAll moves st = foldl (\st' dir -> moveHeadAndTail dir st') st moves

moveAllAndCountVisited :: [Char] -> Int
moveAllAndCountVisited moves = Set.size visited
    where (_, visited) = moveAll moves (rope, Set.singleton s)
          s = (0, 0)
          rope = take 10 $ repeat s

parse :: String -> [Char]
parse s = concat [take n (repeat dir) | line <- lines s, let dir = head line, let n = read (drop 2 line)::Int]

main = do
    c <- getContents

    print $ moveAllAndCountVisited $ parse c
