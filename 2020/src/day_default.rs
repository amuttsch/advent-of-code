use crate::AoC;

pub struct DayDefault {}

static DATA: &'static str = include_str!("input/DayDefault.txt");

impl AoC for DayDefault {
    fn day(&self) -> u8 {
        5
    }

    fn part1(&self) -> String {
       String::new()
    }

    fn part2(&self) -> String {
        String::new()
    }
}


#[cfg(test)]
mod tests {
    use crate::AoC;

    use super::*;

    #[test]
    fn part1() {
        assert_eq!("foo", DayDefault {}.part1());
    }

    #[test]
    fn part2() {
        assert_eq!("bar", DayDefault {}.part2());
    }
}
