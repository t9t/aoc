use std::error::Error;
use std::io::{self, Write};
use std::time::{Duration, Instant};

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
mod day20;
mod day21;
mod day3;
mod day4;
mod day5;
mod day6;
mod day7;
mod day8;
mod day9;
/*mod newday*/

// https://unix.stackexchange.com/a/26592
const CLEAR_LINE: &str = "\u{001b}[2K\r";
const YEAR: u16 = 2017;

fn main() {
    let args: Vec<String> = std::env::args().collect();

    if args.len() == 2 && (args[1] == "all" || args[1] == "results" || args[1] == "benchmark") {
        run_all(&args[1]).unwrap();
        return;
    }

    if args.len() != 3 {
        println!("Usage:");
        println!("\t{} <day> <part>", args[0]);
        println!("\tor:");
        println!("\t{} <all | benchmark | results>", args[0]);
        std::process::exit(1);
    }

    let day = args[1].parse::<u8>().unwrap();
    let part = args[2].parse::<u8>().unwrap();

    println!("Running Year: {}; Day: {}; Part: {}", YEAR, day, part);
    let input = read_input(YEAR, day).unwrap();
    let fun = get_fun(day, part).unwrap();
    let start = Instant::now();
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

fn run_all(mode: &str) -> Result<(), Box<dyn Error>> {
    let benchmark_mode = mode == "benchmark";
    let results_mode = mode == "results";

    if !results_mode {
        println!("| Year | Day | Part | Output                           | Run time   |");
        println!("|------|-----|------|----------------------------------|------------|");
    }

    let total = all_funs().len();
    let begin = Instant::now();
    let mut total_time = Duration::ZERO;
    let mut results: Vec<(u16, u8, u8, String, Duration)> = Vec::new();
    for day in 1..26 {
        let input = read_input(YEAR, day)?;
        for part in 1..3 {
            let fun_opt = get_fun(day, part);
            if let Some(fun) = fun_opt {
                if !results_mode {
                    print!(
                        "{} {:}/{:2}; {:?}; day: {}; part: {}",
                        CLEAR_LINE,
                        (day - 1) * 2 + part,
                        total,
                        begin.elapsed(),
                        day,
                        part
                    );
                    flush_stdout();
                }

                let start = Instant::now();
                let result = fun(input.as_str())?;
                let took = start.elapsed();
                total_time += took;

                if benchmark_mode {
                    results.push((YEAR, day, part, result, took));
                } else if results_mode {
                    println!("{}-{}-{}: {}", YEAR, day, part, result);
                } else {
                    println!(
                        "{}| {:4} | {:3} | {:4} | {:>32} | {:>10} |",
                        CLEAR_LINE,
                        YEAR,
                        day,
                        part,
                        result,
                        format!("{:?}", took)
                    );
                }
                flush_stdout();
            }
        }
    }

    if benchmark_mode {
        print!("{}", CLEAR_LINE);
        results.sort_by(|a, b| a.4.cmp(&b.4));
        for (year, day, part, result, took) in results {
            println!(
                "| {:4} | {:3} | {:4} | {:>32} | {:>10} |",
                year,
                day,
                part,
                result,
                format!("{:?}", took)
            );
        }
    }

    if !results_mode {
        println!("\nTotal run time: {:?}", total_time);
    }

    return Ok(());
}

fn flush_stdout() {
    io::stdout().flush().unwrap();
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
        day20::part2,
        day21::part1,
        day21::part2, /*newday*/
    ];
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[ignore] // use "cargo test test_all" to run
    fn test_all() {
        let results = std::fs::read_to_string(format!("../input/{}/results.txt", YEAR)).unwrap();
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
            assert_eq!(result, expected, "{}-{}-{}", YEAR, day, part);
        }
    }
}
