import System.Environment
import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set
import Data.Maybe

type Pos = (Int, Int)

distance :: Pos -> Pos -> Int
distance (x1, y1) (x2, y2) = (abs $ x2-x1) + (abs $ y2-y1)

-- "x=2, y=18"
parsePos :: String -> Pos
parsePos s = (parseInt l', parseInt r')
    where (l, r)   = break (==',') s
          l'       = drop 2 l
          r'       = drop 4 r
          parseInt = (\n -> read n :: Int)

-- "Sensor at x=2, y=18: closest beacon is at x=-2, y=15"
parseLine :: String -> (Pos, Pos)
parseLine s = (sensor, beacon)
    where s'     = drop 10 s
          (l, r) = break (==':') s'
          sensor = parsePos l
          r'     = drop 23 r
          beacon = parsePos r'

xMinAndMax :: Int -> Int -> Pos -> Maybe (Int, Int)
xMinAndMax y closestDelta (sx, sy) = if dy > closestDelta then Nothing else Just (sx-dxMax, sx+dxMax)
    where dy    = abs (y - sy)
          dxMax = closestDelta - dy

main = do
    c <- getContents

    args <- getArgs
    let y = case (listToMaybe $ map (\s -> read s :: Int) $ take 1 args) of
              Just n -> n
              Nothing -> 2000000
    let sensorsWithClosestBeacons = map parseLine (lines c)
    let xMinsAndMaxes = catMaybes [xMinAndMax y (distance s b) s | (s, b) <- sensorsWithClosestBeacons]
    let beaconXsAtY = nub $ map fst $ filter (\(_, by) -> by == y) $ map snd sensorsWithClosestBeacons
    let (minX, maxX) = foldl1 (\(lmin, lmax) (rmin, rmax) -> (min lmin rmin, max lmax rmax)) xMinsAndMaxes
    print $ maxX - minX + 1 - (length beaconXsAtY)
