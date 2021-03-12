use crate::AoC;
use itertools::Itertools;

pub struct Day01 {
    data01: String,
}

impl Day01 {
    pub fn new() -> Day01 {
        Day01 {
            data01: String::from(include_str!("input/day01.txt"))
        }
    }
}

impl AoC for Day01 {
    fn day(&self) -> u8 {
        1
    }

    fn part1(&self) -> String {
        self.compute_expense(2)
    }

    fn part2(&self) -> String {
        self.compute_expense(3)
    }
}

impl Day01 {
    fn compute_expense(&self, n: usize) -> String {
        let mut numbers: Vec<u32> = self.data01
            .lines()
            .map(|line| line.parse::<u32>())
            .filter(|number| number.is_ok())
            .map(|n| n.unwrap())
            .collect();

        numbers.sort();
        for perm in numbers.into_iter().permutations(n) {
            if perm.iter().fold(0, |acc, x| acc + x) == 2020 {
                let mul = perm.into_iter().fold(1,  |acc, x| acc * x);
                return mul.to_string()
            }
        }

        String::new()
    }
}

#[cfg(test)]
mod tests {
    use crate::day01::Day01;
    use crate::AoC;

    #[test]
    fn part1() {
        assert_eq!("1009899", Day01::new().part1());
    }

    #[test]
    fn part2() {
        assert_eq!("44211152", Day01::new().part2());
    }
}
