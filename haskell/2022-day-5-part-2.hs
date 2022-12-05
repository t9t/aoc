import Data.Char
import Data.List

breakOn c s = (l, tail r) where (l, r) = break (c==) s
cleanse s = [c | c <- s, c `elem` ['A'..'Z']]
zipAll ss = if length ss == 1 then head ss else zipWith (++) (head ss) (zipAll (tail ss))
toInt s = (read s :: Int)

partitionTo n s
    | len <= n  = [s]
    | otherwise = l : partitionTo n s'
    where len = length s
          (l, s') = splitAt n s

parse s = parseInputTuple $ breakOn "" $ lines s
parseStacks l = zipAll $ map (map cleanse . partitionTo 4) $ init l 

parseInputTuple :: ([String], [String]) -> ([String], [(Int, Int, Int)])
parseInputTuple (stacks, procedures) = (parseStacks stacks, map parseProcedure procedures)
readTops stack = map head stack


parseProcedure :: String -> (Int, Int, Int)
parseProcedure s = (toInt (w !! 1), toInt (w !! 3), toInt (w !! 5))
    where w = words s


applyProcedure :: (String, Int) -> String -> String -> Int -> Int -> String
applyProcedure (stack, number) fromHead fromTail from to
    | from == number = fromTail
    | to == number   = fromHead ++ stack
    | otherwise      = stack

-- move 2 from 4 to 9
-- (2,4,9)
process :: [String] -> (Int, Int, Int) -> [String]
process stacks (n, from, to) = 
    let fromStack            = stacks !! (from-1)
        toStack              = stacks !! (to-1)
        (fromHead, fromTail) = splitAt n (stacks !! (from-1))
        withIndices          = zip stacks [1..]
    in [applyProcedure tup fromHead fromTail from to | tup <- withIndices]


processAll :: ([String], [(Int, Int, Int)]) -> [String]
processAll (stacks, []) = stacks
processAll (stacks, procedures) = 
    let processed = process stacks (head procedures)
    in  processAll (processed, (tail procedures))

main = do
    c <- getContents
    print $ readTops $ processAll (parse c)
