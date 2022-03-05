use std::error::Error;

mod day1;
mod day10;
mod day11;
mod day12;
mod day13;
mod day14;
mod day15;
mod day16;
mod day17;
mod day18;
mod day19;
mod day2;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;
mod day20;
/*mod newday*/

fn main() {
    let args: Vec<String> = std::env::args().collect();

    if args.len() == 1 {
        run_all().unwrap();
        return;
    }

    if args.len() != 3 {
        println!("Usage:");
        println!("\t{} <day> <part>", args[0]);
        std::process::exit(1);
    }

    let day = args[1].parse::<u8>().unwrap();
    let part = args[2].parse::<u8>().unwrap();

    println!("Running Year: 2017; Day: {}; Part: {}", day, part);
    let input = read_input(2017, day).unwrap();
    let fun = get_fun(day, part).unwrap();
    let start = std::time::Instant::now();
    let result = fun(input.as_str()).unwrap();
    let duration = start.elapsed();
    println!("Result ({:?}): {}", duration, result);
}

fn read_input(year: u16, day: u8) -> std::io::Result<String> {
    let s = std::fs::read_to_string(format!("../input/{}/{}.txt", year, day))?;
    // trim() interferes with day 19, so only remove trailing newline, if it exists
    if s.ends_with('\n') {
        let mut chars = s.chars();
        chars.next_back();
        return Ok(chars.collect());
    }
    return Ok(s);
}

type DayFunc = fn(&str) -> Result<String, Box<dyn Error>>;

fn run_all() -> Result<(), Box<dyn Error>> {
    for day in 1..26 {
        let input = read_input(2017, day)?;
        for part in 1..3 {
            let fun_opt = get_fun(day, part);
            if let Some(fun) = fun_opt {
                let result = fun(input.as_str())?;
                println!("2017-{}-{}: {}", day, part, result);
            }
        }
    }
    return Ok(());
}

fn get_fun(day: u8, part: u8) -> Option<DayFunc> {
    let funs = all_funs();
    let idx = (((day - 1) * 2) + part - 1) as usize;
    return if idx < funs.len() {
        Some(funs[idx])
    } else {
        return None;
    };
}

fn all_funs() -> Vec<DayFunc> {
    return vec![
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
        day13::part2,
        day14::part1,
        day14::part2,
        day15::part1,
        day15::part2,
        day16::part1,
        day16::part2,
        day17::part1,
        day17::part2,
        day18::part1,
        day18::part2,
        day19::part1,
        day19::part2,
        day20::part1,
        day20::part2, /*newday*/
    ];
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[ignore] // use cargo test test_all to run
    fn test_all() {
        let results = std::fs::read_to_string("../input/2017/results.txt").unwrap();
        for line in results.lines() {
            if line == "" {
                continue;
            }

            let mut line_split = line.split(": ");
            let id = line_split.next().unwrap();
            let expected = line_split.next().unwrap();
            let mut id_split = id.split("-");
            let year = id_split.next().unwrap().parse::<u16>().unwrap();
            let day = id_split.next().unwrap().parse::<u8>().unwrap();
            let part = id_split.next().unwrap().parse::<u8>().unwrap();

            let input = read_input(year, day).unwrap();
            let fun = get_fun(day, part).unwrap();
            let result = fun(input.as_str()).unwrap();
            assert_eq!(result, expected, "2017-{}-{}", day, part);
        }
    }
}
