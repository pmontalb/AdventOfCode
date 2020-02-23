use std::fs::File;
use std::io::Read;
use std::io::Write;

fn decode(encoded: &str) -> String {
    let mut ret = String::new();

    let mut i = 0;
    while i < encoded.len() {
        let c = encoded.chars().nth(i).unwrap();
        if c != '(' {
            if c == ')' {
                panic!("{} <- {}", ret, &encoded[..i]);
            }
            ret.push(c);
            i += 1;
        //println!("i({}) pushing[{}]", i, c);
        } else {
            let mut j = i + 1;
            loop {
                if encoded.chars().nth(j).unwrap() == 'x' {
                    break;
                }
                j += 1;
            }
            let n_chars: usize = encoded[i + 1..j].parse().unwrap();

            let mut k = j + 1;
            loop {
                if encoded.chars().nth(k).unwrap() == ')' {
                    break;
                }
                k += 1;
            }
            let n_times: usize = encoded[j + 1..k].parse().unwrap();
            for _ in 0..n_times {
                ret.push_str(&encoded[k + 1..k + 1 + n_chars]);
            }
            //println!("i({}) pushing[{}] [{}] times", i, &encoded[k + 1..k + 1 + n_chars], n_times);
            i = k + 1 + n_chars;
        }
    }

    ret
}

fn get_decompressed_len(encoded: &str) -> usize {
    let mut ret = 0;

    let mut i = 0;
    while i < encoded.len() {
        let c = encoded.chars().nth(i).unwrap();
        if c == '(' {
            let mut j = i + 1;
            loop {
                if encoded.chars().nth(j).unwrap() == 'x' {
                    break;
                }
                j += 1;
            }
            let n_chars: usize = encoded[i + 1..j].parse().unwrap();

            let mut k = j + 1;
            loop {
                if encoded.chars().nth(k).unwrap() == ')' {
                    break;
                }
                k += 1;
            }
            let n_times: usize = encoded[j + 1..k].parse().unwrap();

            ret += n_times * get_decompressed_len(&encoded[k + 1..k + 1 + n_chars]);
            i = k + 1 + n_chars;
        } else {
            ret += 1;
            i += 1;
        }
    }

    ret
}

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");
    //println!("{}", contents);
    //println!("{}", decode("ADVENT"));
    //println!("{}", decode("A(1x5)BC"));
    //println!("{}", decode("(3x3)XYZ"));
    //println!("{}", decode("A(2x2)BCD(2x2)EFG"));
    //println!("{}", decode("(6x1)(1x3)A"));
    //println!("{}", decode("X(8x2)(3x3)ABCY"));

    let decompressed_content = decode(&contents);
    let len_decompressed_output = decompressed_content
        .chars()
        .filter(|c| !c.is_whitespace())
        .collect::<String>()
        .len();
    println!("{}", len_decompressed_output);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(len_decompressed_output.to_string().as_bytes())
        .expect("Unable to write");

    let total_decompressed_length = get_decompressed_len(&contents);
    println!("{}", total_decompressed_length);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(total_decompressed_length.to_string().as_bytes())
        .expect("Unable to write");
}
