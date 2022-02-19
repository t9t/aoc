use std::env;
use std::fs;
use std::io;

mod day1;
/*mod newday*/

fn main() {
    let funs = [day1::part1, day1::part2 /*newday*/];
    let args: Vec<String> = env::args().collect();

    let day = args[1].parse::<u8>().unwrap();
    let part = args[2].parse::<u8>().unwrap();

    println!("Day: {}; part: {}", day, part);
    let input = read_input(2017, day).unwrap();
    let fun = funs[(((day - 1) * 2) + part - 1) as usize];
    println!("Result: {}", fun(input.trim()));
}

fn read_input(year: u16, day: u8) -> io::Result<String> {
    return fs::read_to_string(format!("../input/{}/{}.txt", year, day));
}
