pub mod day01;
pub mod day02;

pub trait AoC {
    fn day(&self) -> u8;
    fn part1(&self) -> String;
    fn part2(&self) -> String;
}

pub fn run_and_print_result<T>(aoc: T)
where T: AoC
{
    let result1 = aoc.part1();
    let result2 = aoc.part2();

    println!();
    println!("*** Results for day {:#02} ***", aoc.day());
    println!("Part 1: {}", result1);
    println!("Part 2: {}", result2);
}
