use std::fs::File;
use std::io::Read;
use std::io::Write;

struct TriangleCandidate {
    sides: [u32; 3],
}
impl TriangleCandidate {
    fn parse(&mut self, line: &str) {
        let tokens = line.split_whitespace().collect::<Vec<&str>>();
        if tokens.len() != 3 {
            panic!("");
        }
        for i in 0..3 {
            self.sides[i] = tokens[i].parse().expect("Unable to parse");
        }
        self.sides.sort();
    }

    fn is_valid(&self) -> bool {
        self.sides[0] + self.sides[1] > self.sides[2]
    }
}

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");

    let mut valid_triangles = 0;
    let lines = &contents.split("\n").collect::<Vec<&str>>();
    for line in lines {
        let mut triangle_candidate = TriangleCandidate { sides: [0, 0, 0] };
        triangle_candidate.parse(line);

        valid_triangles += triangle_candidate.is_valid() as i32;
    }
    println!("{}", valid_triangles);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(valid_triangles.to_string().as_bytes())
        .expect("Unable to write");

    if lines.len() % 3 != 0 {
        panic!("Needs multiples of 3");
    }

    valid_triangles = 0;
    for i in 0..lines.len() / 3 {
        // parse the 3x3 matrix
        let mut tokenized_rows = Vec::new();
        for j in 0..3 {
            tokenized_rows.push(
                lines[3 * i + j]
                    .trim()
                    .split_whitespace()
                    .collect::<Vec<&str>>(),
            );
        }

        for j in 0..3 {
            let mut s = String::new();
            for k in 0..3 {
                s.push_str(" ");
                s.push_str(tokenized_rows[k][j]);
            }
            let mut triangle_candidate = TriangleCandidate { sides: [0, 0, 0] };
            triangle_candidate.parse(&s);

            valid_triangles += triangle_candidate.is_valid() as i32;
        }
    }
    println!("{}", valid_triangles);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(valid_triangles.to_string().as_bytes())
        .expect("Unable to write");
}
