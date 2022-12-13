import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

data PacketValue = IntValue Int | ListValue [PacketValue] deriving (Eq, Show)
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

parsePacket :: String -> PacketValue
parsePacket s = fst $ parseNextValue s

parseFullPacket :: String -> [PacketValue]
parseFullPacket s = case parsePacket s of
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

parseAllPackets :: String -> [PacketValue]
parseAllPackets s = map parsePacket packetLines
    where packetLines = filter (not . null) $ lines s

packetValueOrdering :: PacketValue -> PacketValue -> Ordering
packetValueOrdering l r = case comparePacketValues l r of
    RightOrder -> LT
    SameOrder -> EQ
    WrongOrder -> GT

main = do
    c <- getContents

    let parsed = parseAllPackets c
    let div1 = parsePacket "[[2]]"
    let div2 = parsePacket "[[6]]"
    let withDividerPackets = div1 : div2 : parsed
    let sorted = sortBy packetValueOrdering withDividerPackets
    let withIndices = zip sorted [1..]
    let dividers = filter (\(p, _) -> p == div1 || p == div2) withIndices
    let decoderKey = product $ map snd dividers
    print decoderKey
