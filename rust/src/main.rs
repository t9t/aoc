use std::env;
use std::fs;
use std::io;

mod day1;

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);

    let day = args[1].parse::<u8>().unwrap();
    let part = args[2].parse::<u8>().unwrap();

    println!("Day: {}; part: {}", day, part);

    let input = read_input(2017, day).unwrap();
    let s = &input.trim();
    let output: u32;
    if part == 1 {
        output = day1::calc_sum(s);
    } else {
        output = day1::calc_sum2(s);
    }
    println!("Result: {}", output);
}

fn read_input(year: u16, day: u8) -> io::Result<String> {
    return fs::read_to_string(format!("../input/{}/{}.txt", year, day));
}
