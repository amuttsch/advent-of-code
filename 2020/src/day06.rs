use crate::AoC;
use std::collections::HashSet;

pub struct Day06 {}

static DATA: &'static str = include_str!("input/day06.txt");

impl AoC for Day06 {
    fn day(&self) -> u8 {
        6
    }

    fn part1(&self) -> String {
        let mut previous_line_empty = false;
        let mut current_group_answers = HashSet::new();
        let mut group_answers: Vec<HashSet<char>> = Vec::new();

        for line in DATA.lines() {
            if previous_line_empty {
                group_answers.push(current_group_answers);
                current_group_answers = HashSet::new();
            }

            for c in line.chars() {
                current_group_answers.insert(c);
            }

            previous_line_empty = line.is_empty();
        }

        group_answers.push(current_group_answers);

        let answer_count: usize = group_answers
            .iter()
            .map(|set| set.len())
            .sum();

        answer_count.to_string()
    }

    fn part2(&self) -> String {
        let mut current_group_answers = HashSet::new();
        let mut group_answers: Vec<HashSet<char>> = Vec::new();
        let mut is_totally_distinct = false;

        for line in DATA.lines() {
            if line.is_empty() {
                group_answers.push(current_group_answers);
                current_group_answers = HashSet::new();
                is_totally_distinct = false;
                continue;
            }

            if is_totally_distinct {
                continue;
            }

            let mut person_answers = HashSet::new();
            for c in line.chars() {
                person_answers.insert(c);
            }

            if current_group_answers.len() > 0 {
                let intersection: HashSet<_> = current_group_answers
                    .intersection(&person_answers)
                    .collect();

                let mut new_current_group_answers = HashSet::new();
                for a in intersection {
                    new_current_group_answers.insert(*a);
                }

                is_totally_distinct = new_current_group_answers.is_empty();

                current_group_answers = new_current_group_answers;
            } else {
                current_group_answers = person_answers;
            }
        }

        group_answers.push(current_group_answers);

        let answer_count: usize = group_answers
            .iter()
            .map(|set| set.len())
            .sum();

        answer_count.to_string()
    }
}

#[cfg(test)]
mod tests {
    use crate::AoC;

    use super::*;

    #[test]
    fn part1() {
        assert_eq!("6542", Day06 {}.part1());
    }

    #[test]
    fn part2() {
        assert_eq!("3299", Day06 {}.part2());
    }
}
