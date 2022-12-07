import Data.Char
import Data.List
import qualified Data.Map as Map

type Path = [String]
type DirSizes = Map.Map Path Int
type InContext = (Path, String, DirSizes, Int)
type OutContext = (Path, DirSizes, Int)

cdIn :: InContext -> OutContext
cdIn (cwd, op, sizes, totalSize) = (cwd ++ [drop 5 op], sizes, totalSize)

cdOut :: InContext -> OutContext
cdOut (cwd, op, sizes, totalSize) = (init cwd, sizes, totalSize)

addSizeToDir :: Int -> Path -> DirSizes -> DirSizes
addSizeToDir n inPath sizes =
    Map.insertWith (+) inPath n sizes

addSizeToAllDirs :: Int -> Path -> DirSizes -> DirSizes
addSizeToAllDirs n inPath sizes
    | null inPath = sizes
    | otherwise   = addSizeToAllDirs n (init inPath) (addSizeToDir n inPath sizes)

recordSize :: InContext -> OutContext
recordSize (cwd, op, sizes, totalSize) =
    (cwd, addSizeToAllDirs size cwd sizes, totalSize - size)
    where w = words op
          (l, r) = (head w, last w)
          size = read l :: Int

operate :: InContext -> OutContext
operate (cwd, op, sizes, totalSize)
    | op == "$ ls"        = (cwd, sizes, totalSize)
    | op == "$ cd .."     = cdOut (cwd, op, sizes, totalSize)
    | op == "$ cd /"      = (["/"], sizes, totalSize)
    | take 4 op == "$ cd" = cdIn (cwd, op, sizes, totalSize)
    | take 3 op == "dir"  = (cwd, sizes, totalSize)
    | otherwise           = recordSize (cwd, op, sizes, totalSize)


operateAll :: (Path, [String], DirSizes, Int) -> (DirSizes, Int)
operateAll (cwd, operations, sizes, totalSize)
    | null operations = (sizes, totalSize)
    | otherwise       = operateAll(cwd', tail operations, sizes', totalSize')
        where (cwd', sizes', totalSize') = operate (cwd, head operations, sizes, totalSize)

filterLowerThan :: (DirSizes, Int) -> DirSizes
filterLowerThan (sizes, n) = Map.filter (\v -> v >= (30000000 - n)) sizes

findSmallest :: DirSizes -> Int
findSmallest = minimum . Map.elems

main = do
    c <- getContents
    print $ findSmallest $ filterLowerThan $ operateAll ([], lines c, Map.empty, 70000000)
