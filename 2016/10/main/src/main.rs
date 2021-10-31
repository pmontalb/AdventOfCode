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
    let instructionsStr = &contents.split("\n").collect::<Vec<&str>>();

    const needles: [i32; main::aoc::N_NEEDLES] = [61, 17];
    let mut instructions = main::aoc::Instructions::new(&needles);
    let _targetBotIdx = instructions.process(&instructionsStr);
   
    assert_eq!(instructions.bots.len(), 210);
    assert_eq!(instructions.outputs.len(), 210);
    assert_eq!(instructions.pendingInstructions.len(), 0);

    println!("checkIdx = {:?}", instructions.checkIdx);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(instructions.checkIdx.to_string().as_bytes())
        .expect("Unable to write");

    let multRes = instructions.outputs[0] * instructions.outputs[1] * instructions.outputs[2];
    println!("multRes = {:?}", multRes);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(multRes.to_string().as_bytes())
        .expect("Unable to write");
}
