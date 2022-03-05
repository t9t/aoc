use std::error::Error;

pub fn part1(s: &str) -> Result<String, Box<dyn Error>> {
    // The one with lowest total acceleration will eventuall stay closest to 0
    let mut lowest = i32::MAX;
    let mut lowest_i = 0;

    for (i, line) in s.lines().enumerate() {
        let mut parts = line.split("a=<");
        parts.next();
        let a = parts.next().unwrap().split(">").next().unwrap();
        let (ax, ay, az) = parse_coords(a)?;
        let tot = ax.abs() + ay.abs() + az.abs();
        if tot < lowest {
            lowest = tot;
            lowest_i = i;
        }
    }
    return Ok(format!("{}", lowest_i));
}

pub fn part2(s: &str) -> Result<String, Box<dyn Error>> {
    let mut particles = parse_particles(s)?;

    // TODO: proper stop condition rather than just looping 100 times
    for _ in 0..1_00 {
        let mut to_remove: Vec<usize> = Vec::new();
        for (i, left) in particles.iter().enumerate() {
            for j in i + 1..particles.len() {
                let right = &particles[j];
                if left.pos == right.pos {
                    if !to_remove.contains(&i) {
                        to_remove.push(i);
                    }
                    if !to_remove.contains(&j) {
                        to_remove.push(j);
                    }
                }
            }
        }
        for i in (0..to_remove.len()).rev() {
            particles.remove(to_remove[i]);
        }

        for mut part in &mut particles {
            let v = part.v;
            part.v = (v.0 + part.a.0, v.1 + part.a.1, v.2 + part.a.2);
            let pos = part.pos;
            part.pos = (pos.0 + part.v.0, pos.1 + part.v.1, pos.2 + part.v.2);
        }
    }
    return Ok(format!("{}", particles.len()));
}

fn parse_particles(s: &str) -> Result<Vec<Particle>, Box<dyn Error>> {
    let mut particles: Vec<Particle> = Vec::new();
    for line in s.lines() {
        let mut pva = line.split(">, ");
        let mut p_part = pva.next().unwrap().split("<");
        let mut v_part = pva.next().unwrap().split("<");
        let mut a_part = pva.next().unwrap().split(">").next().unwrap().split("<");
        p_part.next();
        v_part.next();
        a_part.next();

        let particle = Particle {
            pos: parse_coords(p_part.next().unwrap())?,
            v: parse_coords(v_part.next().unwrap())?,
            a: parse_coords(a_part.next().unwrap())?,
        };

        particles.push(particle);
    }
    return Ok(particles);
}

fn parse_coords(s: &str) -> Result<(i32, i32, i32), Box<dyn Error>> {
    let nums = s
        .split(",")
        .map(|x| x.trim().parse::<i32>())
        .collect::<Result<Vec<i32>, _>>()?;
    return Ok((nums[0], nums[1], nums[2]));
}

#[derive(Debug)]
struct Particle {
    pos: (i32, i32, i32),
    v: (i32, i32, i32),
    a: (i32, i32, i32),
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = "p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>";
        assert_eq!(part1(input).unwrap(), "0");
    }

    #[test]
    fn test_part2() {
        let input = "p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>";
        assert_eq!(part2(input).unwrap(), "1");
    }
}
