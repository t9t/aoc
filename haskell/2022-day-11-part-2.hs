import Data.Char
import Data.List
import qualified Data.Map as Map
import qualified Data.Set as Set

type Operation = Int -> Int
instance Show Operation where show _ = show "(operation)"
type Test = Int -> Bool
instance Show Test where show _ = show "(test)"
type MonkeyMap = Map.Map Int Monkey

data Monkey = Monkey {
    number :: Int,
    items :: [Int],
    operation :: Operation,
    test :: Test,
    divisor :: Int,
    ifTrueThrowTo :: Int,
    ifFalseThrowTo :: Int,
    inspected :: Int
} deriving (Show)

partitionOn :: (Eq a) => a -> [a] -> [[a]]
partitionOn sep s = case dropWhile (==sep) s of
    [] -> []
    s' -> p : partitionOn sep s''
        where (p, s'') = break (==sep) s'

-- "Monkey 0:"
parseNumber :: String -> Int
parseNumber s = read (init $ drop 7 s) :: Int

-- "  Starting items: 54, 65, 75, 74"
parseItems :: String -> [Int]
parseItems s = map (\n -> read n :: Int) $ partitionOn ',' $ drop 18 s

parseOperator :: Char -> (Int -> Int -> Int)
parseOperator '+' = (+)
parseOperator '*' = (*)

parseOperand :: String -> (Int -> Int)
parseOperand " old" = (\old -> old)
parseOperand s = (\_ -> (read s) :: Int)

-- "  Operation: new = old + 6"
-- TODO: old * old
parseOperation :: String -> Operation
parseOperation s = (\old -> old `op` (oper old))
    where r    = drop 23 s
          op   = parseOperator $ head r
          oper = parseOperand $ tail r

-- "  Test: divisible by 19"
parseTest :: String -> Test
parseTest s = (\i -> (i `mod` n) == 0)
    where n = (read $ drop 21 s) :: Int

-- "  Test: divisible by 19"
parseDivisor :: String -> Int
parseDivisor s = n
    where n = (read $ drop 21 s) :: Int

-- ": throw to monkey 2"
parseThrowTo :: String -> Int
parseThrowTo s = (read $ drop 18 s) :: Int

-- "    If true: throw to monkey 2"
parseIfTrue :: String -> Int
parseIfTrue s = parseThrowTo (drop 11 s)

-- "    If false: throw to monkey 0"
parseIfFalse :: String -> Int
parseIfFalse s = parseThrowTo (drop 12 s)

{-
Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3
["Monkey 2:","  Starting items: 79, 60, 97","  Operation: new = old * old","  Test: divisible by 13","    If true: throw to monkey 1","    If false: throw to monkey 3"]
-}
parseMonkey :: [String] -> Monkey
parseMonkey lines = Monkey {
    number = parseNumber (lines !! 0),
    items = parseItems (lines !! 1),
    operation = parseOperation (lines !! 2),
    test = parseTest (lines !! 3),
    divisor = parseDivisor (lines !! 3),
    ifTrueThrowTo = parseIfTrue (lines !! 4),
    ifFalseThrowTo = parseIfFalse (lines !! 5),
    inspected = 0
}

get :: MonkeyMap -> Int -> Monkey
get monkeys num = case Map.lookup num monkeys of
    Just monkey -> monkey
    Nothing -> error "Oh no"

addItem :: Monkey -> Int -> Monkey
addItem monkey item = monkey { items = (items monkey) ++ [item] }

inspectItems :: Int -> Monkey -> Monkey
inspectItems d monkey = monkey { items = items'', inspected = inspected' }
    where items' = map (operation monkey) $ items monkey
          items'' = map (`mod` d) items'
          inspected' = (inspected monkey) + (length items')

throwItems :: Monkey -> MonkeyMap -> MonkeyMap
throwItems monkey monkeys
    | null $ items monkey = monkeys
    | otherwise = throwItems monkey' monkeys''
        where item = head $ items monkey
              monkey' = monkey { items = tail $ items monkey }
              tested  = test monkey item
              trueTarget  = ifTrueThrowTo monkey
              falseTarget = ifFalseThrowTo monkey
              targetNumber = if tested then trueTarget else falseTarget
              target  = get monkeys targetNumber
              target' = addItem target item
              monkeys'  = Map.insert (number monkey) monkey' monkeys
              monkeys'' = Map.insert targetNumber target' monkeys'


turn :: Int -> Int -> MonkeyMap -> MonkeyMap
turn d monkeyNumber monkeys = case items monkey of
    [] -> monkeys
    items -> monkeys'
    where monkey = get monkeys monkeyNumber
          monkey' = inspectItems d monkey
          monkeys' = throwItems monkey' monkeys

turnAll :: Int -> Int -> MonkeyMap -> MonkeyMap
turnAll d monkeyNumber monkeys
    | monkeyNumber == Map.size monkeys = monkeys
    | otherwise = turnAll d (monkeyNumber+1) monkeys'
    where monkeys' = turn d monkeyNumber monkeys

processRound :: Int -> MonkeyMap -> MonkeyMap
processRound d monkeys = turnAll d 0 monkeys

rounds :: Int -> Int -> MonkeyMap -> MonkeyMap
rounds d n monkeys
    | n == 0    = monkeys
    | otherwise = rounds d (n-1) monkeys'
    where monkeys' = processRound d monkeys

toMap :: [Monkey] -> MonkeyMap
toMap l = Map.fromList [(number monkey, monkey) | monkey <- l]

business :: MonkeyMap -> Int
business monkeys = top * second
    where l = Map.elems monkeys
          inspecteds = map inspected l
          sorted = reverse $ sort inspecteds
          top = head sorted
          second = head $ tail sorted

main = do
    c <- getContents

    let monkeyList = map parseMonkey $ partitionOn "" $ lines $ c
    let monkeys = toMap monkeyList
    let chillFactor = product $ map divisor monkeyList

    putStrLn $ show chillFactor

    let monkeys' = rounds chillFactor 10000 monkeys
    putStrLn "Input:"
    putStrLn $ unlines $ map show $ Map.elems monkeys
    putStrLn "Output:"
    putStrLn $ unlines $ map show $ Map.elems monkeys'
    print $ business monkeys'
