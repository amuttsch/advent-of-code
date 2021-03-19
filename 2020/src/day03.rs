use crate::AoC;

pub struct Day03 {}

static DATA: &'static str = include_str!("input/day03.txt");

impl AoC for Day03 {
    fn day(&self) -> u8 {
        3
    }

    fn part1(&self) -> String {
        let trees = trees(3, 1);

        trees.to_string()
    }

    fn part2(&self) -> String {
        let sum1 = trees(1,1);
        let sum2 = trees(3, 1);
        let sum3 = trees(5,1 );
        let sum4 = trees(7, 1);
        let sum5 = trees(1, 2);

        (sum1 * sum2 * sum3 * sum4 * sum5).to_string()
    }
}

fn trees(right: usize, down: usize) -> i64 {
    let mut position = 0;
    let mut trees = 0;
    let mut lines = DATA.lines().skip(1);

    loop {
        let mut line = "";

        for _ in 0..down {
            line = match lines.next() {
                None => {
                    return trees;
                },
                Some(l) => l
            }
        }

        position = (position + right) % line.len();
        if line.chars().nth(position).unwrap() == '#' {
            trees += 1;
        }
    }
}

#[cfg(test)]
mod tests {
    use crate::day03::Day03;
    use crate::AoC;

    #[test]
    fn part1() {
        assert_eq!("218", Day03 {}.part1());
    }

    #[test]
    fn part2() {
        assert_eq!("3847183340", Day03 {}.part2());
    }
}
