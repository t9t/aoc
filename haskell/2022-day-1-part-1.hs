import Data.Char


main = do
    c <- getContents
    print $ maximum [sum $ toInts ss | ss <- (partitionOn "" (lines c))]
    where toInts strings = [read n :: Integer | n <- strings ]
          partitionOn sep s = case dropWhile null s of
                [] -> []
                s' -> p : partitionOn sep s''
                    where (p, s'') = break null s'
