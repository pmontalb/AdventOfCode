use std::fs::File;
use std::io::Read;
use std::io::Write;

pub mod aoc;

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");

    let mut sum_valid_sector = 0;
    let lines = &contents.split("\n").collect::<Vec<&str>>();

    for line in lines {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse(line);

        if r.is_real() {
            sum_valid_sector += r.sector_id;
        }
    }
    println!("{}", sum_valid_sector);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(sum_valid_sector.to_string().as_bytes())
        .expect("Unable to write");

    for line in lines {
        let mut r = main::aoc::Room {
            ..Default::default()
        };
        r.parse(line);

        if r.decode().contains("northpole") {
            //println!("{} -> {}", r.name, r.decode());
            println!("{}", r.sector_id);
            let mut output2 = File::create(&"output2").expect("Unable to create");
            output2
                .write_all(r.sector_id.to_string().as_bytes())
                .expect("Unable to write");
            break;
        }
    }
}
