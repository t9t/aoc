import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

data Tile = Air | NotAir deriving (Eq, Show) -- Rock | Sand
type Pos = (Int, Int)
type Grid = Map.Map Pos Tile
data DripOutcome = CameToRest Pos | Clogged deriving (Eq, Show)

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
dripSand floorY (x, y) grid = if nextDrip == Nothing && x == 500 && y == 0 then Clogged else case nextDrip of
    Just dripPos -> dripSand floorY dripPos grid
    Nothing      -> CameToRest (x, y)
    where isAir          = \pos -> Air == get pos grid
          below          = (x, y+1)
          belowIsFloor   = floorY == y+1
          belowIsAir     = if belowIsFloor then False else isAir below
          downLeft       = (x-1, y+1)
          downLeftIsAir  = if belowIsFloor then False else isAir downLeft
          downRight      = (x+1, y+1)
          downRightIsAir = if belowIsFloor then False else isAir downRight
          nextDrip       = if belowIsAir then Just below
                           else if downLeftIsAir then Just downLeft
                           else if downRightIsAir then Just downRight
                           else Nothing

dripSandUntilClogged :: Int -> Int -> Grid -> Int
dripSandUntilClogged floorY n grid = if dripped == Clogged then n else dripSandUntilClogged floorY (n+1) grid'
    where dripped = dripSand floorY (500, 0) grid
          grid'   = case dripped of
                    Clogged    -> grid
                    CameToRest restPos -> Map.insert restPos NotAir grid
          
-- TODO: slow as fuck
main = do
    c <- getContents

    let grid = parse c
    let maxY = maximum $ map snd $ Map.keys grid
    let floorY = maxY + 2
    print $ 1 + dripSandUntilClogged floorY 0 grid
