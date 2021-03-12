use crate::AoC;
use regex::Regex;

pub struct Day02 {
    data: String,
}

impl Day02 {
    pub fn new() -> Day02 {
        Day02 {
            data: String::from(include_str!("input/day02.txt"))
        }
    }
}

impl AoC for Day02 {
    fn day(&self) -> u8 {
        2
    }

    fn part1(&self) -> String {
        let matched_passwords = self.count_passwords(Box::new(|min, max, char, password| {
            let char_count = password.chars()
                .filter(|c| *c == char)
                .collect::<Vec<char>>()
                .len();

            char_count >= min && char_count <= max
        }));

        matched_passwords.to_string()
    }

    fn part2(&self) -> String {
        let matched_passwords = self.count_passwords(Box::new(|first, second, char, password| {
            let first_char = password.chars().nth(first - 1).unwrap();
            let second_char = password.chars().nth(second - 1).unwrap();

            (first_char == char && second_char != char)
                || (first_char != char && second_char == char)
        }));

        matched_passwords.to_string()
    }
}

impl Day02 {
    fn count_passwords(&self, f: Box<dyn Fn(usize, usize, char, String) -> bool>) -> i32 {
        let re = Regex::new(r"(\d+)-(\d+) (\w): (\w+)").unwrap();
        let mut matched_passwords = 0;

        for line in self.data.lines() {
            let cap = re.captures(line).unwrap();
            let min = cap[1].parse::<usize>().unwrap();
            let max = cap[2].parse::<usize>().unwrap();
            let char = cap[3].chars().next().unwrap();
            let password = cap[4].to_string();

            if f(min, max, char, password) {
                matched_passwords += 1;
            }
        }

        matched_passwords
    }
}


#[cfg(test)]
mod tests {
    use crate::day02::Day02;
    use crate::AoC;

    #[test]
    fn part1() {
        assert_eq!("469", Day02::new().part1());
    }

    #[test]
    fn part2() {
        assert_eq!("267", Day02::new().part2());
    }
}
