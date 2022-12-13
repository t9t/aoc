import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

data PacketValue = IntValue Int | ListValue [PacketValue] deriving Show
data Order = RightOrder | WrongOrder | SameOrder deriving (Eq, Show)

parseNextIntValue :: String -> (Int, String)
parseNextIntValue s = (read ns :: Int, rest)
    where (ns, rest) = break (`elem` ",]") s

parseNextListValue :: String -> ([PacketValue], String)
parseNextListValue s
    | head s == ']' = ([], drop 1 s)
    | head s == ',' = parseNextListValue $ tail s
    | otherwise     = (v : next, rest')
    where (v, rest)     = parseNextValue s
          (next, rest') = parseNextListValue rest

parseNextValue :: String -> (PacketValue, String)
parseNextValue s
    | head s == '[' = (ListValue l, lrest)
    | otherwise     = (IntValue n, nrest)
    where (l, lrest) = parseNextListValue $ tail s
          (n, nrest) = parseNextIntValue s

parseFullPacket :: String -> [PacketValue]
parseFullPacket s = case fst $ parseNextValue s of
    ListValue l -> l
    _ -> error "bad input"

comparePacketValues :: PacketValue -> PacketValue -> Order
comparePacketValues (ListValue []) (ListValue []) = SameOrder
comparePacketValues (ListValue []) (ListValue _) = RightOrder
comparePacketValues (ListValue _) (ListValue []) = WrongOrder
comparePacketValues (ListValue l) (ListValue r) = if first == SameOrder then comparePacketValues (ListValue $ tail l) (ListValue $ tail r) else first
    where first = comparePacketValues (head l) (head r)
comparePacketValues (IntValue l) (IntValue r) = if l < r then RightOrder else if l > r then WrongOrder else SameOrder
comparePacketValues (ListValue l) (IntValue r) = comparePacketValues (ListValue l) (ListValue [IntValue r])
comparePacketValues (IntValue l) (ListValue r) = comparePacketValues (ListValue [IntValue l]) (ListValue r)

comparePacketStrings :: String -> String -> Order
comparePacketStrings l r = comparePacketValues l'' r''
    where l'  = parseFullPacket l
          r'  = parseFullPacket r
          l'' = ListValue l'
          r'' = ListValue r'

sumOfPacketIndicesInTheRightOrder :: Int -> [String] -> Int
sumOfPacketIndicesInTheRightOrder _ [] = 0
sumOfPacketIndicesInTheRightOrder idx s = n + (sumOfPacketIndicesInTheRightOrder (idx+1) s')
    where next3    = take 3 s
          packet1  = next3 !! 0
          packet2  = next3 !! 1
          compared = comparePacketStrings packet1 packet2
          right    = compared == RightOrder
          n        = if right then idx else 0
          s'       = drop 3 s


main = do
    c <- getContents

    print $ sumOfPacketIndicesInTheRightOrder 1 $ lines c
