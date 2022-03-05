use std::collections::HashMap;
use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", enhance_multi_and_count_on(s, 5)?));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    return Ok(format!("{}", enhance_multi_and_count_on(s, 18)?));
}

type Picture = Vec<String>;

fn enhance_multi_and_count_on(s: &str, num: u32) -> Result<u32, Box<dyn Error>> {
    let enhancements = parse_and_expand_rules(s)?;
    let mut picture = split(".#.\n..#\n###");
    let enhanced = enhance_multi(&picture, &enhancements, num);
    picture = enhanced;

    return Ok(count_on(&picture));
}

fn count_on(picture: &Picture) -> u32 {
    let mut count = 0;
    for row in picture {
        for c in row.chars() {
            if c == '#' {
                count += 1;
            }
        }
    }
    return count;
}

fn enhance_multi(picture: &Picture, enhancements: &HashMap<Picture, Picture>, num: u32) -> Picture {
    let mut pic = picture.clone();
    for _ in 0..num {
        pic = enhance(&pic, enhancements);
    }
    return pic;
}

fn enhance(picture: &Picture, enhancements: &HashMap<Picture, Picture>) -> Picture {
    if picture.len() % 2 != 0 && picture.len() % 3 != 0 {
        panic!("Invalid picture");
    }

    let section_size = if picture.len() % 2 == 0 { 2 } else { 3 };
    let sections = divide_picture(&picture, section_size);

    let x: Vec<Vec<&Picture>> = sections
        .iter()
        .map(|row| {
            row.iter()
                .map(|p| enhancements.get(p).unwrap())
                .collect::<Vec<&Picture>>()
        })
        .collect();

    let stitched = stitch(&x);

    return stitched;
}

fn stitch(sections: &Vec<Vec<&Picture>>) -> Picture {
    let mut ret = Picture::new();
    let section_size = sections[0][0].len();
    for section_row in sections {
        for j in 0..section_size {
            let mut new_line = Vec::new();
            for section_col in section_row {
                new_line.push(section_col[j].clone());
            }
            let x = new_line.join("");
            ret.push(x);
        }
    }
    return ret;
}

fn divide_picture(picture: &Picture, size: usize) -> Vec<Vec<Picture>> {
    let mut divisions: Vec<Vec<Picture>> = Vec::new();
    let sections = picture.len() / size;
    for sub_row in 0..sections {
        let mut sections_row: Vec<Picture> = Vec::new();
        for sub_col in 0..sections {
            let min_row = sub_row * size;
            let max_row = sub_row * size + size;
            let min_col = sub_col * size;
            let max_col = sub_col * size + size;

            let section = picture[min_row..max_row]
                .iter()
                .map(|line| {
                    return String::from(&line[min_col..max_col]);
                })
                .collect::<Picture>();
            sections_row.push(section);
        }
        divisions.push(sections_row);
    }
    return divisions;
}

fn parse_and_expand_rules(s: &str) -> Result<HashMap<Picture, Picture>, Box<dyn Error>> {
    let mut enhancements: HashMap<Picture, Picture> = HashMap::new();
    for line in s.lines() {
        let mut parts = line.split(" => ");
        let left = parts.next().ok_or(format!("Invalid line: {}", line))?;
        let right = parts.next().ok_or(format!("Invalid line: {}", line))?;

        let input = left.replace("/", "\n");
        let output = right.replace("/", "\n");

        let all_inputs = all_orientations(input.as_str());
        for xformed_input in all_inputs {
            enhancements.insert(split(&xformed_input), split(&output));
        }
    }
    return Ok(enhancements);
}

fn split(s: &str) -> Picture {
    return s.lines().map(|x| String::from(x)).collect::<Picture>();
}

fn all_orientations(s: &str) -> Picture {
    let mut orientations = Vec::new();
    for flip in [flip_no, flip_h, flip_v] {
        let flipped = flip(s);
        for rotate in [rotate_0, rotate_90, rotate_180, rotate_270] {
            orientations.push(rotate(flipped.as_str()));
        }
    }

    return orientations;
}

fn flip_no(s: &str) -> String {
    return String::from(s);
}

fn flip_h(s: &str) -> String {
    // Reverse each line
    return s
        .lines()
        .map(|line| line.chars().rev().collect::<String>())
        .collect::<Vec<String>>()
        .join("\n");
}

fn flip_v(s: &str) -> String {
    // Just reverse the lines
    return s.lines().rev().collect::<Vec<&str>>().join("\n");
}

fn rotate_0(s: &str) -> String {
    return String::from(s);
}

fn rotate_90(s: &str) -> String {
    let mut new_lines = Vec::new();
    let lines = s.lines().collect::<Vec<&str>>();
    for col in 0..lines.len() {
        let mut new_line = Vec::new();
        for row in (0..lines.len()).rev() {
            new_line.push(lines[row].chars().nth(col).unwrap());
        }
        new_lines.push(new_line.iter().collect::<String>());
    }
    return new_lines.join("\n");
}

fn rotate_180(s: &str) -> String {
    return rotate_90(rotate_90(s).as_str());
}

fn rotate_270(s: &str) -> String {
    return rotate_90(rotate_180(s).as_str());
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_flip_h() {
        assert_eq!(flip_h(".#.\n..#\n###"), ".#.\n#..\n###");
    }

    #[test]
    fn test_flip_v() {
        assert_eq!(flip_v(".#.\n..#\n###"), "###\n..#\n.#.");
    }

    #[test]
    fn test_rotate_90() {
        assert_eq!(rotate_90(".#.\n..#\n###"), "#..\n#.#\n##.");
    }

    #[test]
    fn test_rotate_180() {
        assert_eq!(rotate_180(".#.\n..#\n###"), "###\n#..\n.#.");
    }

    #[test]
    fn test_rotate_270() {
        assert_eq!(rotate_270(".#.\n..#\n###"), ".##\n#.#\n..#");
    }

    #[test]
    fn test_enhance_multi_and_count_on() {
        let input = "../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#";
        assert_eq!(enhance_multi_and_count_on(input, 2).unwrap(), 12);
    }
}
