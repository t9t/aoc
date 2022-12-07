import Data.Char
import Data.List
import qualified Data.Map as Map

type Path = [String]
type DirSizes = Map.Map Path Int
type InContext = (Path, String, DirSizes)
type OutContext = (Path, DirSizes)

cdIn :: InContext -> OutContext
cdIn (cwd, op, sizes) = (cwd ++ [drop 5 op], sizes)

cdOut :: InContext -> OutContext
cdOut (cwd, op, sizes) = (init cwd, sizes)

addSizeToDir :: Int -> Path -> DirSizes -> DirSizes
addSizeToDir n inPath sizes =
    Map.insertWith (+) inPath n sizes

addSizeToAllDirs :: Int -> Path -> DirSizes -> DirSizes
addSizeToAllDirs n inPath sizes
    | null inPath = sizes
    | otherwise   = addSizeToAllDirs n (init inPath) (addSizeToDir n inPath sizes)

recordSize :: InContext -> OutContext
recordSize (cwd, op, sizes) =
    (cwd, addSizeToAllDirs size cwd sizes)
    where w = words op
          (l, r) = (head w, last w)
          size = read l :: Int

operate :: InContext -> OutContext
operate (cwd, op, sizes)
    | op == "$ ls"        = (cwd, sizes)
    | op == "$ cd .."     = cdOut (cwd, op, sizes)
    | op == "$ cd /"      = (["/"], sizes)
    | take 4 op == "$ cd" = cdIn (cwd, op, sizes)
    | take 3 op == "dir"  = (cwd, sizes)
    | otherwise           = recordSize (cwd, op, sizes)


operateAll :: ([String], [String], DirSizes) -> DirSizes
operateAll (cwd, operations, sizes)
    | null operations = sizes
    | otherwise       = operateAll(cwd', tail operations, sizes')
        where (cwd', sizes') = operate (cwd, head operations, sizes)

sumSmallish :: DirSizes -> Int
sumSmallish sizes = sum (Map.elems $ Map.filter (\v -> v <= 100000) sizes)

main = do
    c <- getContents
    print $ sumSmallish $ operateAll ([], lines c, Map.empty)
