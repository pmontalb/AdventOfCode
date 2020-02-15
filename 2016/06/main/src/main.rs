use std::cmp::*;
use std::error::Error;
use std::fs::File;
use std::io::Read;
use std::io::Write;

#[derive(Clone, Debug, Default, PartialOrd, PartialEq, Eq)]
struct Letter {
    grapheme: char,
    count: u32,
}
impl Ord for Letter {
    fn cmp(&self, other: &Self) -> Ordering {
        self.count.cmp(&other.count)
    }
}

fn decode_message_worker<'a>(
    entries: &'a Vec<&str>,
    sorterFunctor: &impl Fn(&Letter, &Letter) -> Ordering,
) -> Result<String, &'a str> {
    const N_LETTERS: usize = 26;
    let first_entry = entries[0];

    let mut word_counts = Vec::<Vec<Letter>>::new();
    word_counts.resize(first_entry.len(), Vec::<Letter>::new());
    for word_count in word_counts.iter_mut() {
        word_count.resize(N_LETTERS, Default::default());
        for i in 0..N_LETTERS {
            word_count[i].grapheme = ('a' as u8 + i as u8) as char;
            word_count[i].count = 0;
        }
    }

    for entry in entries.iter() {
        if entry.len() != first_entry.len() {
            return Err("len");
        }
        for i in 0..first_entry.len() {
            let idx = (entry.chars().nth(i).unwrap() as u8 - 'a' as u8) as usize;
            word_counts[i][idx].count += 1;
        }
    }

    for i in 0..first_entry.len() {
        // sort by word count
        word_counts[i].sort_by(|m, n| sorterFunctor(&m, &n));
    }

    let mut ret = String::new();
    for i in 0..first_entry.len() {
        ret.push(word_counts[i][0].grapheme);
    }

    Ok(ret)
}

fn get_most_common_entries<'a>(entries: &'a Vec<&str>) -> Result<String, &'a str> {
    decode_message_worker(entries, &|m: &Letter, n: &Letter| m.cmp(n).reverse())
}
fn get_least_common_entries<'a>(entries: &'a Vec<&str>) -> Result<String, &'a str> {
    decode_message_worker(entries, &|m: &Letter, n: &Letter| m.cmp(n))
}

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");
    let lines = &contents.split("\n").collect::<Vec<&str>>();

    let decoded_message_1 = get_most_common_entries(lines).unwrap();
    println!("{}", decoded_message_1);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(decoded_message_1.as_bytes())
        .expect("Unable to write");

    let decoded_message_2 = get_least_common_entries(lines).unwrap();
    println!("{}", decoded_message_2);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(decoded_message_2.as_bytes())
        .expect("Unable to write");
}
