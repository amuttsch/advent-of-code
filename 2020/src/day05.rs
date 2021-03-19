use crate::AoC;
use std::cmp;

pub struct Day05 {}

static DATA: &'static str = include_str!("input/day05.txt");

impl AoC for Day05 {
    fn day(&self) -> u8 {
        5
    }

    fn part1(&self) -> String {

        let seat_id = DATA.lines()
            .map(|line| decode(line))
            .fold(0, |acc, cur| {
                let seat_id = get_seat_id(cur);

                if seat_id > acc {
                    seat_id
                } else {
                    acc
                }
            });

        seat_id.to_string()
    }

    fn part2(&self) -> String {
        let mut seat_ids: Vec<i32> = DATA.lines()
            .map(|line| get_seat_id(decode(line)))
            .collect();

        seat_ids.sort();

        for (i, seat_id) in seat_ids.iter().enumerate() {
            if seat_ids.get(i+1).unwrap_or(&0) - seat_id != 1 {
                return (seat_id + 1).to_string()
            }
        }

        String::new()
    }
}

fn get_seat_id(seat: (i32, i32)) -> i32 {
    seat.0*8+seat.1
}

fn decode(seat: &str) -> (i32, i32) {
    let row_init: i32 = 64;
    let column_init: i32 = 4;
    let row = seat.chars().take(7);
    let column = seat.chars().skip(7).take(3);

    let row = row.enumerate().fold(row_init - 1, |acc, (idx, cur)| {
        let new_val = row_init / i32::pow(2, idx as u32 + 1);
        return if cur == 'F' {
            acc - new_val
        } else {
            acc + cmp::max(1,new_val)
        }
    });

    let column = column.enumerate().fold(column_init - 1, |acc, (idx, cur)| {
        let new_val = column_init / i32::pow(2, idx as u32 + 1);
        return if cur == 'L' {
            acc - new_val
        } else {
            acc + cmp::max(1,new_val)
        }
    });

    (row , column)
}


#[cfg(test)]
mod tests {
    use crate::AoC;

    use super::*;

    #[test]
    fn test_case_1() {
        assert_eq!(357, get_seat_id(decode("FBFBBFFRLR")));
    }

    #[test]
    fn test_case_2() {
        assert_eq!(567, get_seat_id(decode("BFFFBBFRRR")));
    }

    #[test]
    fn test_case_3() {
        assert_eq!(119, get_seat_id(decode("FFFBBBFRRR")));
    }

    #[test]
    fn test_case_4() {
        assert_eq!(820, get_seat_id(decode("BBFFBBFRLL")));
    }

    #[test]
    fn part1() {
        assert_eq!("970", Day05 {}.part1());
    }

    #[test]
    fn part2() {
        assert_eq!("587", Day05 {}.part2());
    }
}
