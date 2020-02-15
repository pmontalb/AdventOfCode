use std::fs::File;
use std::io::Read;
use std::io::Write;

struct IpV7 {
    supernet_sequences: Vec<String>,
    hypernet_sequences: Vec<String>,
}
impl Default for IpV7 {
    fn default() -> Self {
        IpV7 {
            supernet_sequences: Vec::new(),
            hypernet_sequences: Vec::new(),
        }
    }
}

impl IpV7 {
    const LENGTH_ABBA: usize = 4;
    const LENGTH_ABA: usize = 3;

    fn parse(&mut self, line: &str) {
        let mut running_index = 0;
        let mut parsing_brackets = false;
        for i in 0..line.len() {
            if line.chars().nth(i).unwrap() == '[' {
                self.supernet_sequences
                    .push(String::from(&line[running_index..i]));
                //println!("Found supernet: {}", &line[running_index..i]);
                running_index = i;
                parsing_brackets = true;
            }
            if line.chars().nth(i).unwrap() == ']' {
                if !parsing_brackets {
                    panic!("");
                }
                parsing_brackets = false;

                self.hypernet_sequences
                    .push(String::from(&line[running_index + 1..i]));
                //println!("Found hypernet: {}", &line[running_index + 1..i]);
                running_index = i + 1;
            }
        }
        // parse the last token
        if parsing_brackets {
            panic!("");
        }
        self.supernet_sequences
            .push(String::from(&line[running_index..]));
        //println!("Found supernet: {}", &line[running_index..]);
    }

    fn contains_abba_sequence(token: &str) -> bool {
        if IpV7::LENGTH_ABBA % 2 != 0 {
            panic!("");
        }

        for i in 0..token.len() - IpV7::LENGTH_ABBA + 1 {
            let lhs = &token[i..i + IpV7::LENGTH_ABBA / 2];
            let rhs = &token[i + IpV7::LENGTH_ABBA / 2..i + IpV7::LENGTH_ABBA];
            if lhs != rhs && lhs == rhs.chars().rev().collect::<String>() {
                // println!(
                //     "support TLS: left({}) right({})",
                //     &token[i..i + IpV7::LENGTH_ABBA / 2],
                //     &token[i + IpV7::LENGTH_ABBA / 2..i + IpV7::LENGTH_ABBA]
                // );
                return true;
            }
        }
        false
    }

    fn contains_aba_sequence(token: &str) -> Vec<String> {
        if IpV7::LENGTH_ABA % 3 != 0 {
            panic!("");
        }
        let mut ret: Vec<String> = Vec::new();

        for i in 0..token.len() - IpV7::LENGTH_ABA + 1 {
            let lhs = &token[i..i + IpV7::LENGTH_ABA / 3];
            let mhs = &token[i + IpV7::LENGTH_ABA / 3..i + IpV7::LENGTH_ABA * 2 / 3];
            let rhs = &token[i + IpV7::LENGTH_ABA * 2 / 3..i + IpV7::LENGTH_ABA];
            if lhs == rhs && lhs != mhs {
                ret.push(String::from(&token[i..i + IpV7::LENGTH_ABA]));
            }
        }
        ret
    }
    fn get_bab_from_aba(token: &str) -> String {
        if token.len() != IpV7::LENGTH_ABA {
            panic!("{}", &token);
        }
        let mut ret = String::with_capacity(IpV7::LENGTH_ABA);

        ret.push(token.chars().nth(1).unwrap());
        if token.chars().nth(2).unwrap() != token.chars().nth(0).unwrap() {
            panic!("");
        }
        ret.push(token.chars().nth(0).unwrap());
        ret.push(token.chars().nth(1).unwrap());
        //println!("{} --> {}", token, &ret);

        ret
    }
    fn contains_bab_sequence_from_aba(&self, aba_sequence: &str) -> bool {
        let target_bab = IpV7::get_bab_from_aba(&aba_sequence);
        for token in self.hypernet_sequences.iter() {
            for i in 0..token.len() - IpV7::LENGTH_ABA + 1 {
                // println!(
                //     "hyp({}) == ({}) [{}] ",
                //     &token[i..i + IpV7::LENGTH_ABA],
                //     &target_bab,
                //     token[i..i + IpV7::LENGTH_ABA] == target_bab
                // );
                if token[i..i + IpV7::LENGTH_ABA] == target_bab {
                    return true;
                }
            }
        }
        false
    }

    fn support_tls(&self) -> bool {
        let mut supernet_sequences_contain_abba = false;
        for sequence in self.supernet_sequences.iter() {
            supernet_sequences_contain_abba |= IpV7::contains_abba_sequence(&sequence);
        }

        let mut hypernet_sequence_does_not_contain_abba = true;
        for sequence in self.hypernet_sequences.iter() {
            hypernet_sequence_does_not_contain_abba &= !IpV7::contains_abba_sequence(&sequence);
        }
        // println!(
        //     "seq({}) hyp({})",
        //     sequences_contain_abba, sequences_contain_abba
        // );

        return supernet_sequences_contain_abba && hypernet_sequence_does_not_contain_abba;
    }

    fn support_ssl(&self) -> bool {
        for supernet_sequence in self.supernet_sequences.iter() {
            let aba_tokens = &IpV7::contains_aba_sequence(&supernet_sequence);
            for aba in aba_tokens.iter() {
                //println!("seq({}) is aba", &aba);
                if self.contains_bab_sequence_from_aba(aba) {
                    return true;
                }
            }
        }
        false
    }
}

fn main() {
    // let mut x: IpV7 = Default::default();
    // x.parse("abba[mnop]qrst");
    // println!("{}", x.support_tls());

    // x.parse("abcd[bddb]xyyx");
    // println!("{}", x.support_tls());

    // x.parse("aaaa[qwer]tyui");
    // println!("{}", x.support_tls());

    // x.parse("ioxxoj[asdfgh]zxcvbn");
    // println!("{}", x.support_tls());

    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");
    let lines = &contents.split("\n").collect::<Vec<&str>>();

    let mut tls_compliant_addresses = 0;
    for line in lines.iter() {
        let mut address: IpV7 = Default::default();
        address.parse(line);

        if address.support_tls() {
            tls_compliant_addresses += 1;
        }
    }
    println!("{}", tls_compliant_addresses);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(tls_compliant_addresses.to_string().as_bytes())
        .expect("Unable to write");

    // let mut x: IpV7 = Default::default();
    // x.parse("aba[bab]xyz");
    // println!("{}", x.support_ssl());

    // let mut x: IpV7 = Default::default();
    // x.parse("xyx[xyx]xyx");
    // println!("{}", x.support_ssl());

    // let mut x: IpV7 = Default::default();
    // x.parse("aaa[kek]eke");
    // println!("{}", x.support_ssl());

    // let mut x: IpV7 = Default::default();
    // x.parse("zazbz[bzb]cdb");
    // println!("{}", x.support_ssl());

    let mut ssl_compliant_addresses = 0;
    for line in lines.iter() {
        let mut address: IpV7 = Default::default();
        address.parse(line);

        if address.support_ssl() {
            ssl_compliant_addresses += 1;
        }
    }
    println!("{}", ssl_compliant_addresses);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(ssl_compliant_addresses.to_string().as_bytes())
        .expect("Unable to write");
}
