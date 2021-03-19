use std::collections::HashMap;

use lazy_static::lazy_static;
use regex::Regex;

use crate::AoC;

pub struct Day04 {}

static DATA: &'static str = include_str!("input/day04.txt");

impl AoC for Day04 {
    fn day(&self) -> u8 {
        4
    }

    fn part1(&self) -> String {
        let valid_passports = parse_passports()
            .into_iter()
            .filter(|p| p.is_valid())
            .count();

        valid_passports.to_string()
    }

    fn part2(&self) -> String {
        let valid_passports = parse_passports()
            .into_iter()
            .filter(|p| p.is_extended_valid())
            .count();

        valid_passports.to_string()
    }
}

#[derive(Debug)]
struct Passport {
    fields: HashMap<String, String>
}

lazy_static! {
    static ref HEIGHT_REGEX: Regex = Regex::new(r"^(\d+)(cm|in)$").unwrap();
    static ref HAIR_COLOR_REGEX: Regex = Regex::new(r"^#[0-9a-f]{6}$").unwrap();
    static ref PASSPORT_ID_REGEX: Regex = Regex::new(r"^[0-9]{9}$").unwrap();
    static ref VALID_EYE_COLORS: [&'static str; 7] = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"];
}

impl Passport {

    fn is_valid(&self) -> bool {
        self.fields.contains_key("byr") &&
        self.fields.contains_key("iyr") &&
        self.fields.contains_key("eyr") &&
        self.fields.contains_key("hgt") &&
        self.fields.contains_key("hcl") &&
        self.fields.contains_key("ecl") &&
        self.fields.contains_key("pid")
    }

    fn is_extended_valid(&self) -> bool {
        match self.fields.get("byr") {
            None => return false,
            Some(byr) => {
                match byr.parse::<i32>() {
                    Ok(byr) => {
                        if byr < 1920 || byr > 2002 {
                            return false
                        }
                    }
                    Err(_) => return false
                }
            }
        };

        match self.fields.get("iyr") {
            None => return false,
            Some(iyr) => {
                match iyr.parse::<i32>() {
                    Ok(iyr) => {
                        if iyr < 2010 || iyr > 2020 {
                            return false
                        }
                    }
                    Err(_) => return false
                }
            }
        };

        match self.fields.get("eyr") {
            None => return false,
            Some(eyr) => {
                match eyr.parse::<i32>() {
                    Ok(eyr) => {
                        if eyr < 2020 || eyr > 2030 {
                            return false
                        }
                    }
                    Err(_) => return false
                }
            }
        };

        match self.fields.get("hgt") {
            None => return false,
            Some(hgt) => {
                match HEIGHT_REGEX.captures(hgt) {
                    None => return false,
                    Some(c) => {
                        let h = c[1].parse::<i32>().unwrap_or(0);
                        if &c[2] == "cm" && (h < 150 || h > 193) { return false; };
                        if &c[2] == "in" && (h < 59 || h > 76) { return false; }
                    }
                }

            }
        }

        match self.fields.get("hcl") {
            None => return false,
            Some(hcl) => {
                if !HAIR_COLOR_REGEX.is_match(hcl) {
                    return false;
                }
            }
        }

        match self.fields.get("ecl") {
            None => return false,
            Some(ecl) => {
                if !VALID_EYE_COLORS.contains(&ecl.as_str()) {
                    return false
                }
            }
        }

        match self.fields.get("pid") {
            None => return false,
            Some(pid) => {
                if !PASSPORT_ID_REGEX.is_match(pid) {
                    return false;
                }
            }
        }

        true
    }
}

fn parse_passports() -> Vec<Passport> {
    let mut previous_line_empty = false;
    let mut current_passport = Passport {
        fields: HashMap::new()
    };
    let mut passports: Vec<Passport> = Vec::new();

    for line in DATA.lines() {
        if previous_line_empty {
            passports.push(current_passport);
            current_passport = Passport {
                fields: HashMap::new()
            };
        }

        let fields: Vec<&str> = line
            .split(' ')
            .collect();

        fields
            .into_iter()
            .map(|field| field.split(':').collect::<Vec<&str>>())
            .filter(|field| field.len() == 2)
            .for_each(|field| {
                let k = String::from(*field.get(0).unwrap());
                let v = String::from(*field.get(1).unwrap());
                current_passport.fields.insert(k, v);
            });

        previous_line_empty = line.is_empty();
    }

    passports.push(current_passport);

    passports
}

#[cfg(test)]
mod tests {
    use crate::AoC;

    use super::*;

    #[test]
    fn part1() {
        assert_eq!("239", Day04 {}.part1());
    }

    #[test]
    fn part2() {
        assert_eq!("188", Day04 {}.part2());
    }
}
