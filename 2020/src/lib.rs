use day01::Day01;
use day02::Day02;
use day03::Day03;
use day04::Day04;
use day05::Day05;

mod day01;
mod day02;
mod day03;
mod day04;
mod day05;

pub trait AoC {
    fn day(&self) -> u8;
    fn part1(&self) -> String;
    fn part2(&self) -> String;
}

pub fn run_and_print_result() {
    run_day_and_print_result(Day01::new());
    run_day_and_print_result(Day02::new());
    run_day_and_print_result(Day03 {});
    run_day_and_print_result(Day04 {});
    run_day_and_print_result(Day05 {});
}

fn run_day_and_print_result<T>(aoc: T)
where T: AoC
{
    let result1 = aoc.part1();
    let result2 = aoc.part2();

    println!();
    println!("*** Results for day {:#02} ***", aoc.day());
    println!("Part 1: {}", result1);
    println!("Part 2: {}", result2);
}
