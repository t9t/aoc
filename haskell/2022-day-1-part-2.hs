import Data.Char
import Data.List

main = do
    c <- getContents
    print . sum . take 3 . reverse . sort $ [sum $ toInts ss | ss <- (partitionOn "" (lines c))]
    where toInts strings = [read n :: Integer | n <- strings ]
          partitionOn sep s = case dropWhile null s of
                [] -> []
                s' -> p : partitionOn sep s''
                    where (p, s'') = break null s'
