use std::fs::File;
use std::io::Read;
use std::io::Write;

mod aoc;

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");
    let lines = &contents.split("\n").collect::<Vec<&str>>();
    if lines.len() != 1 {
        panic!("");
    }
    let input = lines[0];

    const N_ITERATIONS: u64 = 8;
    let solution_part1 = aoc::aoc::find_unordered_password(input, N_ITERATIONS);
    println!("{}", solution_part1);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(solution_part1.to_string().as_bytes())
        .expect("Unable to write");

    let solution_part2 = aoc::aoc::find_ordered_password(input, N_ITERATIONS);
    println!("{}", solution_part2);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(solution_part2.to_string().as_bytes())
        .expect("Unable to write");
}
