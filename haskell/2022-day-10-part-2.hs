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

getPixel :: Int -> Int -> Char
getPixel screenX xRegister =
    if xRegister == screenX || xRegister == (screenX-1) || xRegister == (screenX+1) then
        '#'
    else 
        '.'

crTick :: (Int, Int, Int) -> (Int, Int) -> (Int, Int, Char)
crTick (cycle, xBefore, xAfter) (screenX, screenY) =
    (screenX', screenY', pixel)
    where screenX' = if screenX == 39 then 0 else screenX+1
          screenY' = if screenX' == 0 then screenY + 1 else screenY 
          pixel    = getPixel screenX xBefore

crTickAll :: (Int, Int) -> [Char] -> [(Int, Int, Int)] -> [Char]
crTickAll _ crt [] = crt
crTickAll (x, y) crt instructions =
    pixel : (crTickAll (x', y') crt rest)
    where (c, b, a) = head instructions
          rest      = tail instructions
          (x', y', pixel) = crTick (c, b, a) (x, y)

partitionTo :: Int -> [a] -> [[a]]
partitionTo n s
    | len <= n  = [s]
    | otherwise = l : partitionTo n s'
    where len = length s
          (l, s') = splitAt n s

main = do
    c <- getContents

    putStrLn $ unlines $ partitionTo 40 $ crTickAll (0, 0) [] $ processAll 1 1 $ convertLines $ lines c
