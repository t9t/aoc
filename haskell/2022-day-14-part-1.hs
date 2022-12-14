import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

data Tile = Air | NotAir deriving (Eq, Show) -- Rock | Sand
type Pos = (Int, Int)
type Grid = Map.Map Pos Tile
data DripOutcome = CameToRest Pos | FlowedIntoAbyss deriving (Eq, Show)

-- 498,4
parsePos :: String -> Pos
parsePos s = (toInt l, toInt $ tail r)
    where (l, r) = break (==',') s
          toInt = \n -> read n :: Int

-- 498,4 -> 498,6
getLineSegment :: Pos -> Pos -> [Pos]
getLineSegment l r = if sx == ex then
        [(sx, y) | y <- [(min sy ey)..(max sy ey)]]
    else if sy == ey then
        [(x, sy) | x <- [(min sx ex)..(max sx ex)]]
    else error $ "invalid start/end " ++ (show (sx, sy)) ++ "/" ++ (show (ex, ey))
    where (sx, sy) = l
          (ex, ey) = r

getAllLineSegments :: [Pos] -> [Pos]
getAllLineSegments l
    | length l == 1 = []
    | otherwise     = (getLineSegment first second) ++ (getAllLineSegments $ drop 1 l)
    where first = l !! 0
          second = l !! 1

-- 498,4 -> 498,6 -> 496,6
parseLinePoints :: String -> [Pos]
parseLinePoints "" = []
parseLinePoints s = pos : parseLinePoints (drop 4 r)
    where (l, r)   = break (==' ') s -- "498,4" / " -> 498,6 -> 496,6"
          pos      = parsePos l

parse :: String -> Grid
parse s = foldl (\m pos -> Map.insert pos NotAir m) Map.empty rocks
    where l     = lines s
          rocks = concat $ map (\line -> getAllLineSegments $ parseLinePoints line) l

get :: Pos -> Grid -> Tile
get pos grid = case Map.lookup pos grid of
    Just tile -> tile
    Nothing -> Air

dripSand :: Int -> Pos -> Grid -> DripOutcome
dripSand maxY (x, y) grid = if y > maxY then FlowedIntoAbyss else case nextDrip of
    Just dripPos -> dripSand maxY dripPos grid
    Nothing      -> CameToRest (x, y)
    where isAir          = \pos -> Air == get pos grid
          below          = (x, y+1)
          belowIsAir     = isAir below
          downLeft       = (x-1, y+1)
          downLeftIsAir  = isAir downLeft
          downRight      = (x+1, y+1)
          downRightIsAir = isAir downRight
          nextDrip       = if belowIsAir then Just below
                           else if downLeftIsAir then Just downLeft
                           else if downRightIsAir then Just downRight
                           else Nothing

dripSandUntilSandStartsFlowingIntoTheAbyss :: Int -> Grid -> Int
dripSandUntilSandStartsFlowingIntoTheAbyss n grid = if dripped == FlowedIntoAbyss then n else dripSandUntilSandStartsFlowingIntoTheAbyss (n+1) grid'
    where dripped = dripSand 10_000 (500, 0) grid
          grid'   = case dripped of
                    FlowedIntoAbyss    -> grid
                    CameToRest restPos -> Map.insert restPos NotAir grid
          

main = do
    c <- getContents

    let grid = parse c
    print $ dripSandUntilSandStartsFlowingIntoTheAbyss 0 grid
