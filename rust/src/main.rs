mod day1;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;
mod day10;
mod day11;
mod day12;
mod day13;
/*mod newday*/

fn main() {
    let funs = [
        day1::part1,
        day1::part2,
        day2::part1,
        day2::part2,
        day3::part1,
        day3::part2,
        day4::part1,
        day4::part2,
        day5::part1,
        day5::part2,
        day6::part1,
        day6::part2,
        day7::part1,
        day7::part2,
        day8::part1,
        day8::part2,
        day9::part1,
        day9::part2,
        day10::part1,
        day10::part2,
        day11::part1,
        day11::part2,
        day12::part1,
        day12::part2,
        day13::part1,
        day13::part2, /*newday*/
    ];
    let args: Vec<String> = std::env::args().collect();

    if args.len() != 3 {
        println!("Usage:");
        println!("\t{} <day> <part>", args[0]);
        std::process::exit(1);
    }

    let day = args[1].parse::<u8>().unwrap();
    let part = args[2].parse::<u8>().unwrap();

    println!("Running Year: 2017; Day: {}; Part: {}", day, part);
    let input = read_input(2017, day).unwrap();
    let fun = funs[(((day - 1) * 2) + part - 1) as usize];
    let start = std::time::Instant::now();
    let result = fun(input.trim()).unwrap();
    let duration = start.elapsed();
    println!("Result ({:?}): {}", duration, result);
}

fn read_input(year: u16, day: u8) -> std::io::Result<String> {
    return std::fs::read_to_string(format!("../input/{}/{}.txt", year, day));
}
